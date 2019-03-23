package Count_Words

import (
	"regexp"
	"strings"
)

func CountWords(s string) int {
	s = regexp.MustCompile("[^a-zA-Z\\s]").ReplaceAllString(s, "")
	return len(strings.Fields(s))
}
