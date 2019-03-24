package Pad_String

import (
	"strings"
)

func Pad(s string, l int) string {
	return strings.Repeat(" ", l/2) + s + strings.Repeat(" ", l/2)
}
