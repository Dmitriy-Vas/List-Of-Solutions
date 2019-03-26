package Fizz_Buzz

import "strconv"

func FizzBuzz(i int) string {
	out := ""
	if i % 3 == 0 {
		out += "Fizz"
	}
	if i % 5 == 0 {
		out += "Buzz"
	}
	if out == "" {
		out = strconv.Itoa(i)
	}
	return out
}
