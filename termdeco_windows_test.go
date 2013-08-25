// +build windows

package termdeco

import (
	"fmt"
	"os"
	"testing"
)

func TestSetConsoleTextAttribute(t *testing.T) {
	if err := setConsoleTextAttribute(os.Stdout, c_FOREGROUND_CYAN | c_BACKGROUND_WHITE); err != nil {
		t.Fatalf("setConsoleTextAttribute failed before output: %v", err)
	}
	fmt.Println("Hello, World")
	if err := setConsoleTextAttribute(os.Stdout, c_FOREGROUND_WHITE | c_BACKGROUND_BLACK); err != nil {
		t.Fatalf("setConsoleTextAttribute failed after output: %v", err)
	}
}

func TestGetConsoleScreenBufferInfo(t *testing.T) {
	var info consoleScreenBufferInfo
	if err := getConsoleScreenBufferInfo(os.Stdout, &info); err != nil {
		t.Fatalf("getConsoleScreenBufferInfo failed: %v", err)
	}
	fmt.Printf("Default Attribute: %#04x\n", info.attributes)
}

func TestDecoration(t *testing.T) {
	Printf("Test %s = Bright Red Fg, Green Bg\n", BrightRed("test").BgGreen())
	Println("Test", BrightWhite("bold").Bold().Underline().BgBlack(), "= Bright White Fg, Black Bg, Bold, Underline")
	s := Sprintf("Test %d = Blue Fg, Yellow Bg", Blue(1234).BgYellow())
	Println(s)
}
