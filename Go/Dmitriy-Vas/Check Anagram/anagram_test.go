package Check_Anagram

import (
	"testing"
)

var (
	anagrames = map[string]string {
		"iceman": "cinema",
		"conversation": "voices rant on",
		"dynamite": "may it end",
	}
)

func TestAnagram(t *testing.T) {
	for k, v := range anagrames {
		if !Anagram(k, v) {
			t.Fatal(k, v)
		}
	}
}
