// +build darwin freebsd linux netbsd openbsd

package termdeco

import (
	"fmt"
	"io"
)

// This is fmt.Fprintf wrapper. It works the same as fmt.Fprintf does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(w, format, a...)
}

// This is fmt.Printf wrapper. It works the same as fmt.Printf does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
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
func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprint(w, a...)
}

// This is fmt.Print wrapper. It works the same as fmt.Print does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
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
func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	return fmt.Fprintln(w, a...)
}

// This is fmt.Println wrapper. It works the same as fmt.Println does. For
// cross-platform compatibility, it is recommended using this wrapper for
// printing decorated text.
func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}

// This is fmt.Sprintln wrapper. It works the same as fmt.Sprintln does. It
// returns a string with ANSI escape sequence which can be used for printing
// later with Println or etc.
func Sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}
