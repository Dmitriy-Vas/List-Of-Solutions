package Count_Vowels

import "strings"

func Vowels(s string) int {
	s = strings.ToLower(s)

	vowels := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'u': 0, 'o': 0, 'y': 0}

	for _, c := range s {
		if _, exists := vowels[c]; exists {
			vowels[c]++
		}
	}

	counter := 0
	for _, a := range vowels {
		counter += a
	}
	return counter
}
