package Factorial

import "testing"

var arr = []struct{
	in int
	out int
}{
	{4, 24},
	{5, 120},
	{6, 720},
	{7, 5040},
}

func TestFactorial(t *testing.T) {
	for _, d := range arr {
		if Factorial(d.in) != d.out {
			t.Fatal(d)
		}
	}
}
