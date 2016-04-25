// +build windows

package termdeco

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"syscall"
	"unsafe"
)

const (
	c_FOREGROUND_BLACK     = 0x0000
	c_FOREGROUND_BLUE      = 0x0001
	c_FOREGROUND_GREEN     = 0x0002
	c_FOREGROUND_RED       = 0x0004
	c_FOREGROUND_YELLOW    = c_FOREGROUND_RED   | c_FOREGROUND_GREEN
	c_FOREGROUND_MAGENTA   = c_FOREGROUND_RED   | c_FOREGROUND_BLUE
	c_FOREGROUND_CYAN      = c_FOREGROUND_GREEN | c_FOREGROUND_BLUE
	c_FOREGROUND_WHITE     = c_FOREGROUND_RED   | c_FOREGROUND_GREEN | c_FOREGROUND_BLUE
	c_FOREGROUND_INTENSITY = 0x0008

	c_BACKGROUND_BLACK     = 0x0000
	c_BACKGROUND_BLUE      = 0x0010
	c_BACKGROUND_GREEN     = 0x0020
	c_BACKGROUND_RED       = 0x0040
	c_BACKGROUND_YELLOW    = c_BACKGROUND_RED   | c_BACKGROUND_GREEN
	c_BACKGROUND_MAGENTA   = c_BACKGROUND_RED   | c_BACKGROUND_BLUE
	c_BACKGROUND_CYAN      = c_BACKGROUND_GREEN | c_BACKGROUND_BLUE
	c_BACKGROUND_WHITE     = c_BACKGROUND_RED   | c_BACKGROUND_GREEN | c_BACKGROUND_BLUE
	c_BACKGROUND_INTENSITY = 0x0080

	c_COMMON_LVB_UNDERSCORE = 0x8000
)

type (
	wchar uint16
	short int16
	dword uint32
	word  uint16
	coord struct {
		x, y short
	}
	smallRect struct {
		left, top, right, bottom short
	}
	consoleScreenBufferInfo struct {
		size              coord
		cursorPosition    coord
		attributes        word
		window            smallRect
		maximumWindowSize coord
	}
)

var modkernel32 = syscall.NewLazyDLL("kernel32.dll")
var (
	procSetConsoleTextAttribute    = modkernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo = modkernel32.NewProc("GetConsoleScreenBufferInfo")
)

var stdoutDefaultAttr, stderrDefaultAttr word

func init() {
	var stdoutInfo, stderrInfo consoleScreenBufferInfo
	if err := getConsoleScreenBufferInfo(os.Stdout, &stdoutInfo); err != nil {
		stdoutDefaultAttr = c_FOREGROUND_WHITE | c_BACKGROUND_BLACK
	} else {
		stdoutDefaultAttr = stdoutInfo.attributes
	}

	if err := getConsoleScreenBufferInfo(os.Stdout, &stderrInfo); err != nil {
		stderrDefaultAttr = c_FOREGROUND_WHITE | c_BACKGROUND_BLACK
	} else {
		stderrDefaultAttr = stderrInfo.attributes
	}
}

func prepareConsole(f *os.File) (syscall.Handle, error) {
	h := syscall.Handle(f.Fd())
	if h == syscall.InvalidHandle {
		return h, errors.New("Invalid Handler")
	}
	var m uint32
	if syscall.GetConsoleMode(h, &m) != nil {
		return h, errors.New("Not console file")
	}
	return h, nil
}

func setConsoleTextAttribute(f *os.File, attr word) error {
	h, err := prepareConsole(f)
	if err != nil {
		return err
	}
	return sysSetConsoleTextAttribute(h, attr)
}

func sysSetConsoleTextAttribute(console syscall.Handle, attr word) (err error) {
	r1, _, e1 := syscall.Syscall(procSetConsoleTextAttribute.Addr(), 2, uintptr(console), uintptr(attr), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getConsoleScreenBufferInfo(f *os.File, info *consoleScreenBufferInfo) error {
	h, err := prepareConsole(f)
	if err != nil {
		return err
	}
	return sysGetConsoleScreenBufferInfo(h, info)
}

func sysGetConsoleScreenBufferInfo(console syscall.Handle, info *consoleScreenBufferInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procGetConsoleScreenBufferInfo.Addr(), 2, uintptr(console), uintptr(unsafe.Pointer(info)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

type seqAttr struct {
	Seq []byte
	Attr word
}

var seqAttrMap = []*seqAttr{
	&seqAttr{Seq: escBlack,   Attr: c_FOREGROUND_BLACK},
	&seqAttr{Seq: escRed,     Attr: c_FOREGROUND_RED},
	&seqAttr{Seq: escGreen,   Attr: c_FOREGROUND_GREEN},
	&seqAttr{Seq: escYellow,  Attr: c_FOREGROUND_YELLOW},
	&seqAttr{Seq: escBlue,    Attr: c_FOREGROUND_BLUE},
	&seqAttr{Seq: escMagenta, Attr: c_FOREGROUND_MAGENTA},
	&seqAttr{Seq: escCyan,    Attr: c_FOREGROUND_CYAN},
	&seqAttr{Seq: escWhite,   Attr: c_FOREGROUND_WHITE},

	&seqAttr{Seq: escBrightBlack,   Attr: c_FOREGROUND_BLACK   | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightRed,     Attr: c_FOREGROUND_RED     | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightGreen,   Attr: c_FOREGROUND_GREEN   | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightYellow,  Attr: c_FOREGROUND_YELLOW  | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightBlue,    Attr: c_FOREGROUND_BLUE    | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightMagenta, Attr: c_FOREGROUND_MAGENTA | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightCyan,    Attr: c_FOREGROUND_CYAN    | c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escBrightWhite,   Attr: c_FOREGROUND_WHITE   | c_FOREGROUND_INTENSITY},

	&seqAttr{Seq: escBgBlack,   Attr: c_BACKGROUND_BLACK},
	&seqAttr{Seq: escBgRed,     Attr: c_BACKGROUND_RED},
	&seqAttr{Seq: escBgGreen,   Attr: c_BACKGROUND_GREEN},
	&seqAttr{Seq: escBgYellow,  Attr: c_BACKGROUND_YELLOW},
	&seqAttr{Seq: escBgBlue,    Attr: c_BACKGROUND_BLUE},
	&seqAttr{Seq: escBgMagenta, Attr: c_BACKGROUND_MAGENTA},
	&seqAttr{Seq: escBgCyan,    Attr: c_BACKGROUND_CYAN},
	&seqAttr{Seq: escBgWhite,   Attr: c_BACKGROUND_WHITE},

	&seqAttr{Seq: escBgBrightBlack,   Attr: c_BACKGROUND_BLACK   | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightRed,     Attr: c_BACKGROUND_RED     | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightGreen,   Attr: c_BACKGROUND_GREEN   | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightYellow,  Attr: c_BACKGROUND_YELLOW  | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightBlue,    Attr: c_BACKGROUND_BLUE    | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightMagenta, Attr: c_BACKGROUND_MAGENTA | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightCyan,    Attr: c_BACKGROUND_CYAN    | c_BACKGROUND_INTENSITY},
	&seqAttr{Seq: escBgBrightWhite,   Attr: c_BACKGROUND_WHITE   | c_BACKGROUND_INTENSITY},

	&seqAttr{Seq: escBold,      Attr: c_FOREGROUND_INTENSITY},
	&seqAttr{Seq: escUnderline, Attr: c_COMMON_LVB_UNDERSCORE},
}

func addAttrOfSeq(attr word, defaultAttr word, seq []byte) word {
	if bytes.Equal(seq, escReset) {
		return defaultAttr
	}
	for _, sa := range seqAttrMap {
		if bytes.Equal(seq, sa.Seq) {
			attr |= sa.Attr
			return attr
		}
	}
	return attr
}

func printEscString(f *os.File, str string) (n int, err error) {
	var defaultAttr word
	if f == os.Stdout {
		defaultAttr = stdoutDefaultAttr
	} else if f == os.Stderr {
		defaultAttr = stderrDefaultAttr
	}

	runes := []rune(str)
	end := len(runes)
	printStr := ""
	for i := 0; i < end; {
		for i < end && runes[i] != keyEscape {
			printStr += string(runes[i])
			i++
		}
		if i >= end {
			wn, err := fmt.Fprint(f, printStr)
			n += wn
			if err != nil {
				return n, err
			}
			break
		}
		if runes[i+1] != '[' {
			printStr += string(runes[i])
			i++
			continue
		}
		i += 2
		wn, err := fmt.Fprint(f, printStr)
		n += wn
		if err != nil {
			return n, err
		}
		printStr = ""

		seq := make([]byte, 0)
		attr := word(0)
		for i < end && runes[i] != 'm' {
			if runes[i] == ';' {
				attr = addAttrOfSeq(attr, defaultAttr, seq)
				seq = make([]byte, 0)
			} else {
				seq = append(seq, byte(runes[i]))
			}
			i++
		}
		if i >= end {
			break
		}

		attr = addAttrOfSeq(attr, defaultAttr, seq)
		err = setConsoleTextAttribute(f, attr)
		if err != nil {
			return n, err
		}
		i++
	}
	err = setConsoleTextAttribute(f, defaultAttr)
	if err != nil {
		return n, err
	}
	return
}

// This is fmt.Fprintf wrapper. It works the same as fmt.Fprintf does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
//
// It prints decorated text when using with os.Stdout or os.Stderr as a Writer.
// In other case, It just prints text with ANSI escape sequence
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	if f, ok := w.(*os.File); ok && (f == os.Stdout || f == os.Stderr) {
		str := fmt.Sprintf(format, a...)
		return printEscString(f, str)
	} else {
		return fmt.Fprintf(w, format, a...)
	}
}

// This is fmt.Printf wrapper. It works the same as fmt.Printf does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Printf(format string, a ...interface{}) (n int, err error) {
	return Fprintf(os.Stdout, format, a...)
}

// This is fmt.Sprintf wrapper. It works the same as fmt.Sprintf does. It
// returns a string with ANSI escape sequence which can be used for printing
// later with Println or etc.
func Sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// This is fmt.Fprint wrapper. It works the same as fmt.Fprint does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
//
// It prints decorated text when using with os.Stdout or os.Stderr as a Writer.
// In other case, It just prints text with ANSI escape sequence
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	if f, ok := w.(*os.File); ok && (f == os.Stdout || f == os.Stderr) {
		str := fmt.Sprint(a...)
		return printEscString(f, str)
	} else {
		return fmt.Fprint(w, a...)
	}
}

// This is fmt.Print wrapper. It works the same as fmt.Print does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Print(a ...interface{}) (n int, err error) {
	return Fprint(os.Stdout, a...)
}

// This is fmt.Sprint wrapper. It works the same as fmt.Sprint does. It
// returns a string with ANSI escape sequence which can be used for printing
// later with Println or etc.
func Sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

// This is fmt.Fprintln wrapper. It works the same as fmt.Fprintln does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
//
// It prints decorated text when using with os.Stdout or os.Stderr as a Writer.
// In other case, It just prints text with ANSI escape sequence
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	if f, ok := w.(*os.File); ok && (f == os.Stdout || f == os.Stderr) {
		str := fmt.Sprintln(a...)
		return printEscString(f, str)
	} else {
		return fmt.Fprintln(w, a...)
	}
}

// This is fmt.Println wrapper. It works the same as fmt.Println does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

// This is fmt.Sprintln wrapper. It works the same as fmt.Sprintln does. It
// returns a string with ANSI escape sequence which can be used for printing
// later with Println or etc.
func Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
