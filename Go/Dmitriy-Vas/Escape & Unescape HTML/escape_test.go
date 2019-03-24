package Escape_Unescape

import "testing"

func TestEscape(t *testing.T) {
	if a := Escape("Escape/Unescape HTML"); a != "Escape/Unescape HTML" {
		t.Fatal(a)
	}
	if a := Escape(`"bread"`); a != "&#34;bread&#34;" {
		t.Fatal(a)
	}
}

func TestUnescape(t *testing.T) {
	if a := Unescape("&#34;bread&#34;"); a != `"bread"` {
		t.Fatal(a)
	}
}
