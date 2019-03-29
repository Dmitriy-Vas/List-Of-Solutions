package Pad_String

import (
	"strings"
)

func Pad(s string, l int, char string) string {
	if len(s) > l {
		return s
	}
	return strings.Repeat(char, l/2) + s + strings.Repeat(char, l/2)
}
