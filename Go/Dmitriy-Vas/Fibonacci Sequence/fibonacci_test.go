package Fibonacci_Sequence

import (
	"testing"
)

var arr = []struct{
	in int
	out int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{10, 55},
	{20, 6765},
}

func TestFibonacci(t *testing.T) {
	for _, d := range arr {
		if Fibonacci(d.in) != d.out {
			t.Fatal(d)
		}
	}
}
