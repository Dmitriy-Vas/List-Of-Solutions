package Fizz_Buzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for i := 1; i <= 100; i++ {
		out := FizzBuzz(i)

		if i == 5 && out != "Buzz" {
			t.Fatal(i, out)
		}
		if i == 3 && out != "Fizz" {
			t.Fatal(i, out)
		}
		if i == 15 && out != "FizzBuzz" {
			t.Fatal(i, out)
		}
		if i == 2 && out != "2" {
			t.Fatal(i, out)
		}
	}
}
