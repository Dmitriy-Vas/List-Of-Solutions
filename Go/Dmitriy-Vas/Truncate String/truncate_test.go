package Truncate_String

import "testing"

func TestTruncate(t *testing.T) {
	if a := Truncate("testing string", 7); a != "test..." {
		t.Fatal(a)
	}
	if a := Truncate("totoro", 7); a != "totoro" {
		t.Fatal(a)
	}
	if a := Truncate("testing", 2); a != "testing" {
		t.Fatal(a)
	}
}
