package Check_Anagram

import (
	"testing"
)

var anagrames = []struct {
	s string
	w string
	}{
		{"iceman", "cinema"},
		{"conversation", "voices rant on"},
		{"dynamite", "may it end"},
	}

func TestAnagram(t *testing.T) {
	for _, v := range anagrames {
		if !Anagram(v.s, v.w) {
			t.Fatal(v)
		}
	}
}
