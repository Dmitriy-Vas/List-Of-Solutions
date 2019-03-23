package Count_Vowels

import "testing"

func TestVowels(t *testing.T) {
	if a := Vowels("testo"); a != 2 {
		t.Fatal(a)
	}
	if a := Vowels("Real cool thing, isn't it?"); a != 7 {
		t.Fatal(a)
	}
}
