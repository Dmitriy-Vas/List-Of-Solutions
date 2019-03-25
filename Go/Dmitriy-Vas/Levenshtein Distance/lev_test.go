package Levenshtein_Distance

import "testing"

var tests = []struct{
	first string
	second string
	distance int
}{
	{"test", "test", 0},
	{"kitten", "sitting", 3},
}

func TestLev(t *testing.T) {
	for _, d := range tests {
		if res := Lev(d.first, d.second); res != d.distance {
			t.Fatal(res)
		}
	}
}
