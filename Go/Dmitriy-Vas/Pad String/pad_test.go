package Pad_String

import (
	"strconv"
	"strings"
	"testing"
)

func TestPad(t *testing.T) {
	if a := Pad("test", 4, " "); a != "  test  " {
		t.Fatal(a)
	}
	if a := Pad(strconv.Itoa(500), 50, " "); a != strings.Repeat(" ", 25) + "500" + strings.Repeat(" ", 25) {
		t.Fatal(a)
	}
}
