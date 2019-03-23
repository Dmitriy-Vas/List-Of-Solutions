package Reverse_String

import "testing"

func TestReverse(t *testing.T) {
	if s := Reverse("Racecar"); s != "racecaR" {
		t.Fatal(s)
	}
	if s := Reverse("test"); s != "tset" {
		t.Fatal(s)
	}
	if s := Reverse("something101"); s != "101gnihtemos" {
		t.Fatal(s)
	}
}
