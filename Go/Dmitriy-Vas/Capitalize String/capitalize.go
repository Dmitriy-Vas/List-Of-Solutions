package Capitalize_String

import (
	"regexp"
	"strings"
)

func Capitalize(s string) string {
	s = strings.ToLower(regexp.MustCompile("[^A-Za-z\\s]").ReplaceAllString(s, ""))
	o := strings.Fields(s)
	for i, w := range o {
		o[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(o, " ")
}
