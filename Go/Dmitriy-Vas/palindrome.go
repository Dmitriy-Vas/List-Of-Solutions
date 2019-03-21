package palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(regexp.MustCompile("[^\\w]").ReplaceAllString(s, ""))
	o := []rune(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		w := l - i -1
		o[w], o[i] = o[i], o[w]
	}
	return s == string(o)
}
