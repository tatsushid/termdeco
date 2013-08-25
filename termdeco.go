// termdeco is a small library for cross-platform console text decoration.
//
// It provides functions for text decoration, Formatter, that can be used with
// fmt.Printf like (Fprintf, Sprintf etc.) functions like
//
//	fmt.Println(termdeco.Red(v).BgGreen().Bold())
//
// In this case, v is printed with red, bold text (on Windows, bold is
// translated into text brighter) on green background.
//
// It also provides wrappers for functions in fmt package. For cross-platform
// compatibility, it is recommended using these wrappers for printing decorated
// text like
//
//	termdeco.Println(termdeco.Red(v).BgGreen().Bold())
//
// It implements following ANSI compatible decoration.
//
// Black, Red, Green, Yellow, Blue, Magenta, Cyan, White text and background
// colors.
//
// Bright Black, Bright Red, Bright Green, Bright Yellow, Bright Blue, Bright
// Magenta, Bright Cyan, Bright White text and background colors
//
// Bold and Underline (Underscore) text decoration
//
// Those decoration is implemented as a function of its name whith returns
// Decoration type Formatter so it can be called in a chain like
//
//	termdeco.Red(v).BgGreen().Bold()
//
// and latter called method overwrites previous one, for example,
//
//	termdeco.Red(v).Green()
//
// applies red and after that it overwrites so text printed as green
package termdeco

import (
	"fmt"
)

const (
	c_NONE = iota
	c_BLACK
	c_RED
	c_GREEN
	c_YELLOW
	c_BLUE
	c_MAGENTA
	c_CYAN
	c_WHITE
	c_BRIGHT_BLACK
	c_BRIGHT_RED
	c_BRIGHT_GREEN
	c_BRIGHT_YELLOW
	c_BRIGHT_BLUE
	c_BRIGHT_MAGENTA
	c_BRIGHT_CYAN
	c_BRIGHT_WHITE
)

// Decorator represents a value and its decoration for printing. This type
// implements fmt.Formatter interface and should be used with fmt.Printf
// like (Fprintf, Sprintf etc.) functions like
//
//	fmt.Println(termdeco.Red(v).BgGreen())
//
// For cross-platform compatibility, it is recommended using with wrappers of
// fmt package functions for printing this type like
//
//	termdeco.Println(termdeco.Red(v).BgGreen())
type Decorator struct {
	Value               interface{}
	fgClr, bgClr        int
	isBold, isUnderline bool
}

// It returns an empty Decorator.
func NewDecorator() *Decorator { return &Decorator{} }

func Black(v interface{}) *Decorator         { return &Decorator{Value: v, fgClr: c_BLACK} }
func Red(v interface{}) *Decorator           { return &Decorator{Value: v, fgClr: c_RED} }
func Green(v interface{}) *Decorator         { return &Decorator{Value: v, fgClr: c_GREEN} }
func Yellow(v interface{}) *Decorator        { return &Decorator{Value: v, fgClr: c_YELLOW} }
func Blue(v interface{}) *Decorator          { return &Decorator{Value: v, fgClr: c_BLUE} }
func Magenta(v interface{}) *Decorator       { return &Decorator{Value: v, fgClr: c_MAGENTA} }
func Cyan(v interface{}) *Decorator          { return &Decorator{Value: v, fgClr: c_CYAN} }
func White(v interface{}) *Decorator         { return &Decorator{Value: v, fgClr: c_WHITE} }
func BrightBlack(v interface{}) *Decorator   { return &Decorator{Value: v, fgClr: c_BRIGHT_BLACK} }
func BrightRed(v interface{}) *Decorator     { return &Decorator{Value: v, fgClr: c_BRIGHT_RED} }
func BrightGreen(v interface{}) *Decorator   { return &Decorator{Value: v, fgClr: c_BRIGHT_GREEN} }
func BrightYellow(v interface{}) *Decorator  { return &Decorator{Value: v, fgClr: c_BRIGHT_YELLOW} }
func BrightBlue(v interface{}) *Decorator    { return &Decorator{Value: v, fgClr: c_BRIGHT_BLUE} }
func BrightMagenta(v interface{}) *Decorator { return &Decorator{Value: v, fgClr: c_BRIGHT_MAGENTA} }
func BrightCyan(v interface{}) *Decorator    { return &Decorator{Value: v, fgClr: c_BRIGHT_CYAN} }
func BrightWhite(v interface{}) *Decorator   { return &Decorator{Value: v, fgClr: c_BRIGHT_WHITE} }

func BgBlack(v interface{}) *Decorator         { return &Decorator{Value: v, bgClr: c_BLACK} }
func BgRed(v interface{}) *Decorator           { return &Decorator{Value: v, bgClr: c_RED} }
func BgGreen(v interface{}) *Decorator         { return &Decorator{Value: v, bgClr: c_GREEN} }
func BgYellow(v interface{}) *Decorator        { return &Decorator{Value: v, bgClr: c_YELLOW} }
func BgBlue(v interface{}) *Decorator          { return &Decorator{Value: v, bgClr: c_BLUE} }
func BgMagenta(v interface{}) *Decorator       { return &Decorator{Value: v, bgClr: c_MAGENTA} }
func BgCyan(v interface{}) *Decorator          { return &Decorator{Value: v, bgClr: c_CYAN} }
func BgWhite(v interface{}) *Decorator         { return &Decorator{Value: v, bgClr: c_WHITE} }
func BgBrightBlack(v interface{}) *Decorator   { return &Decorator{Value: v, bgClr: c_BRIGHT_BLACK} }
func BgBrightRed(v interface{}) *Decorator     { return &Decorator{Value: v, bgClr: c_BRIGHT_RED} }
func BgBrightGreen(v interface{}) *Decorator   { return &Decorator{Value: v, bgClr: c_BRIGHT_GREEN} }
func BgBrightYellow(v interface{}) *Decorator  { return &Decorator{Value: v, bgClr: c_BRIGHT_YELLOW} }
func BgBrightBlue(v interface{}) *Decorator    { return &Decorator{Value: v, bgClr: c_BRIGHT_BLUE} }
func BgBrightMagenta(v interface{}) *Decorator { return &Decorator{Value: v, bgClr: c_BRIGHT_MAGENTA} }
func BgBrightCyan(v interface{}) *Decorator    { return &Decorator{Value: v, bgClr: c_BRIGHT_CYAN} }
func BgBrightWhite(v interface{}) *Decorator   { return &Decorator{Value: v, bgClr: c_BRIGHT_WHITE} }

func Bold(v interface{}) *Decorator       { return &Decorator{Value: v, isBold: true} }
func Underline(v interface{}) *Decorator  { return &Decorator{Value: v, isUnderline: true} }
func Underscore(v interface{}) *Decorator { return &Decorator{Value: v, isUnderline: true} }

func (d *Decorator) Black() *Decorator         { d.fgClr = c_BLACK; return d }
func (d *Decorator) Red() *Decorator           { d.fgClr = c_RED; return d }
func (d *Decorator) Green() *Decorator         { d.fgClr = c_GREEN; return d }
func (d *Decorator) Yellow() *Decorator        { d.fgClr = c_YELLOW; return d }
func (d *Decorator) Blue() *Decorator          { d.fgClr = c_BLUE; return d }
func (d *Decorator) Magenta() *Decorator       { d.fgClr = c_MAGENTA; return d }
func (d *Decorator) Cyan() *Decorator          { d.fgClr = c_CYAN; return d }
func (d *Decorator) White() *Decorator         { d.fgClr = c_WHITE; return d }
func (d *Decorator) BrightBlack() *Decorator   { d.fgClr = c_BRIGHT_BLACK; return d }
func (d *Decorator) BrightRed() *Decorator     { d.fgClr = c_BRIGHT_RED; return d }
func (d *Decorator) BrightGreen() *Decorator   { d.fgClr = c_BRIGHT_GREEN; return d }
func (d *Decorator) BrightYellow() *Decorator  { d.fgClr = c_BRIGHT_YELLOW; return d }
func (d *Decorator) BrightBlue() *Decorator    { d.fgClr = c_BRIGHT_BLUE; return d }
func (d *Decorator) BrightMagenta() *Decorator { d.fgClr = c_BRIGHT_MAGENTA; return d }
func (d *Decorator) BrightCyan() *Decorator    { d.fgClr = c_BRIGHT_CYAN; return d }
func (d *Decorator) BrightWhite() *Decorator   { d.fgClr = c_BRIGHT_WHITE; return d }

func (d *Decorator) BgBlack() *Decorator         { d.bgClr = c_BLACK; return d }
func (d *Decorator) BgRed() *Decorator           { d.bgClr = c_RED; return d }
func (d *Decorator) BgGreen() *Decorator         { d.bgClr = c_GREEN; return d }
func (d *Decorator) BgYellow() *Decorator        { d.bgClr = c_YELLOW; return d }
func (d *Decorator) BgBlue() *Decorator          { d.bgClr = c_BLUE; return d }
func (d *Decorator) BgMagenta() *Decorator       { d.bgClr = c_MAGENTA; return d }
func (d *Decorator) BgCyan() *Decorator          { d.bgClr = c_CYAN; return d }
func (d *Decorator) BgWhite() *Decorator         { d.bgClr = c_WHITE; return d }
func (d *Decorator) BgBrightBlack() *Decorator   { d.bgClr = c_BRIGHT_BLACK; return d }
func (d *Decorator) BgBrightRed() *Decorator     { d.bgClr = c_BRIGHT_RED; return d }
func (d *Decorator) BgBrightGreen() *Decorator   { d.bgClr = c_BRIGHT_GREEN; return d }
func (d *Decorator) BgBrightYellow() *Decorator  { d.bgClr = c_BRIGHT_YELLOW; return d }
func (d *Decorator) BgBrightBlue() *Decorator    { d.bgClr = c_BRIGHT_BLUE; return d }
func (d *Decorator) BgBrightMagenta() *Decorator { d.bgClr = c_BRIGHT_MAGENTA; return d }
func (d *Decorator) BgBrightCyan() *Decorator    { d.bgClr = c_BRIGHT_CYAN; return d }
func (d *Decorator) BgBrightWhite() *Decorator   { d.bgClr = c_BRIGHT_WHITE; return d }

func (d *Decorator) Bold() *Decorator       { d.isBold = true; return d }
func (d *Decorator) Underline() *Decorator  { d.isUnderline = true; return d }
func (d *Decorator) Underscore() *Decorator { d.isUnderline = true; return d }

const (
	keyEscape = 27
)

var (
	escSeq = []byte{keyEscape, '['}

	escBlack   = []byte{'3', '0'}
	escRed     = []byte{'3', '1'}
	escGreen   = []byte{'3', '2'}
	escYellow  = []byte{'3', '3'}
	escBlue    = []byte{'3', '4'}
	escMagenta = []byte{'3', '5'}
	escCyan    = []byte{'3', '6'}
	escWhite   = []byte{'3', '7'}

	escBrightBlack   = []byte{'9', '0'}
	escBrightRed     = []byte{'9', '1'}
	escBrightGreen   = []byte{'9', '2'}
	escBrightYellow  = []byte{'9', '3'}
	escBrightBlue    = []byte{'9', '4'}
	escBrightMagenta = []byte{'9', '5'}
	escBrightCyan    = []byte{'9', '6'}
	escBrightWhite   = []byte{'9', '7'}

	escBgBlack   = []byte{'4', '0'}
	escBgRed     = []byte{'4', '1'}
	escBgGreen   = []byte{'4', '2'}
	escBgYellow  = []byte{'4', '3'}
	escBgBlue    = []byte{'4', '4'}
	escBgMagenta = []byte{'4', '5'}
	escBgCyan    = []byte{'4', '6'}
	escBgWhite   = []byte{'4', '7'}

	escBgBrightBlack   = []byte{'1', '0', '0'}
	escBgBrightRed     = []byte{'1', '0', '1'}
	escBgBrightGreen   = []byte{'1', '0', '2'}
	escBgBrightYellow  = []byte{'1', '0', '3'}
	escBgBrightBlue    = []byte{'1', '0', '4'}
	escBgBrightMagenta = []byte{'1', '0', '5'}
	escBgBrightCyan    = []byte{'1', '0', '6'}
	escBgBrightWhite   = []byte{'1', '0', '7'}

	escReset     = []byte{'0'}
	escBold      = []byte{'1'}
	escUnderline = []byte{'4'}
)

var fgEscSeq = [][]byte{
	escBlack, escRed, escGreen, escYellow, escBlue, escMagenta, escCyan, escWhite,
	escBrightBlack, escBrightRed, escBrightGreen, escBrightYellow, escBrightBlue, escBrightMagenta, escBrightCyan, escBrightWhite,
}

var bgEscSeq = [][]byte{
	escBgBlack, escBgRed, escBgGreen, escBgYellow, escBgBlue, escBgMagenta, escBgCyan, escBgWhite,
	escBgBrightBlack, escBgBrightRed, escBgBrightGreen, escBgBrightYellow, escBgBrightBlue, escBgBrightMagenta, escBgBrightCyan, escBgBrightWhite,
}

// This is implementation of fmt.Formatter interface
func (d *Decorator) Format(f fmt.State, c rune) {
	format := ""
	args := make([]interface{}, 0)
	deco := d.buildEscSeq()
	if len(deco) > 0 {
		format += "%s"
		args = append(args, deco)
	}
	format += d.origFormat(f, c) + "%s"
	reset := append(escSeq, escReset...)
	reset = append(reset, 'm')
	args = append(args, d.Value, reset)
	fmt.Fprintf(f, format, args...)
}

func (d *Decorator) buildEscSeq() []byte {
	seq := make([]byte, 0)
	if d.fgClr > 0 {
		seq = append(seq, fgEscSeq[d.fgClr-1]...)
	}
	if d.bgClr > 0 {
		if len(seq) > 0 {
			seq = append(seq, ';')
		}
		seq = append(seq, bgEscSeq[d.bgClr-1]...)
	}
	if d.isBold {
		if len(seq) > 0 {
			seq = append(seq, ';')
		}
		seq = append(seq, escBold...)
	}
	if d.isUnderline {
		if len(seq) > 0 {
			seq = append(seq, ';')
		}
		seq = append(seq, escUnderline...)
	}
	if len(seq) > 0 {
		seq = append(escSeq, seq...)
		seq = append(seq, 'm')
	}
	return seq
}

func (d *Decorator) origFormat(f fmt.State, c rune) string {
	format := "%"
	for i := 0; i < 128; i++ {
		if f.Flag(i) {
			format += string(i)
		}
	}
	if w, ok := f.Width(); ok {
		format += fmt.Sprintf("%d", w)
	}
	if p, ok := f.Precision(); ok {
		format += fmt.Sprintf(".%d", p)
	}
	format += string(c)
	return format
}
