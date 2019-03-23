package Check_Palindrome

import (
	"testing"
)

var (
	palindromes = []string{
		"racecar",
		"rotor!",
		"@sixa$xis!",
		"Kayak",
		"Was it a car or a cat I saw?",
		"A Santa lived as a devil at NASA",
	}
	nonPalindromes = []string{
		"palindrome",
		"weird",
		"51256",
		"real",
	}
)

func TestPalindrome(t *testing.T) {
	for _, p := range palindromes {
		if !isPalindrome(p) {
			t.Fatal(p)
		}
	}
	for _, np := range nonPalindromes {
		if isPalindrome(np) {
			t.Fatal(np)
		}
	}
}
