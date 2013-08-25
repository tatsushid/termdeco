// +build darwin freebsd linux netbsd openbsd

package termdeco

import (
	"fmt"
	"testing"
)

func TestDecoration(t *testing.T) {
	Printf("Test %s = Bright Red Fg, Green Bg\n", BrightRed("test").BgGreen())
	Println("Test", BrightWhite("bold").Bold().Underline().BgBlack(), "= Bright White Fg, Black Bg, Bold, Underline")
	s := Sprintf("Test %d = Blue Fg, Yellow Bg", Blue(1234).BgYellow())
	fmt.Println(s)
}
