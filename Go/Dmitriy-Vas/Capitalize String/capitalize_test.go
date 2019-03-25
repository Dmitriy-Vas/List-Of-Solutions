package Capitalize_String

import "testing"

var tests = []struct{
	s string
	w string
}{
	{"test", "Test"},
	{"large test", "Large Test"},
	{"5t ? 1241", "T"},
}

func TestCapitalize(t *testing.T) {
	for _, f := range tests {
		if a := Capitalize(f.s); a != f.w {
			t.Fatal(a)
		}
	}
}
