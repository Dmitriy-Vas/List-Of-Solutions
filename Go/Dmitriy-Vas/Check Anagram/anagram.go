package Check_Anagram

import (
	"regexp"
	"sort"
	"strings"
)

func Anagram(s string, w string) bool {
	r := regexp.MustCompile("[^a-z]")
	s = r.ReplaceAllString(strings.ToLower(s), "")
	w = r.ReplaceAllString(strings.ToLower(w), "")
	return Normalize(w) == Normalize(s)
}

func Normalize(s string) string {
	runes := strings.Split(s, "")
	sort.Strings(runes)
	return strings.Join(runes, "")
}
