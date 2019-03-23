package Count_Words

import "testing"

var (
	words = []string {
		"Two words",
		"Two words &",
		"Two words 1",
		"     Two    Words    ",
	}
)

func TestCounter(t *testing.T) {
	for _, w := range words {
		if CountWords(w) != 2 {
			t.Fatal(w)
		}
	}
}
