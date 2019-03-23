package Capitalize_String

import "testing"

func TestCapitalize(t *testing.T) {
	if a := Capitalize("test"); a != "Test" {
		t.Fatal(a)
	}
	if a := Capitalize("large test"); a != "Large Test" {
		t.Fatal(a)
	}
	if a := Capitalize("5t ? 1241"); a != "T" {
		t.Fatal(a)
	}
}
