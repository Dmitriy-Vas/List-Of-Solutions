package Prime_Numbers

import (
	"math"
)

func Prime(N float64) []int {
	var x, y, n int
	Sqrt := math.Sqrt(N)

	isPrime := make([]bool, int(N))

	for x = 1; float64(x) <= Sqrt; x++ {
		for y = 1; float64(y) <= Sqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= int(N) && (n%12 == 1 || n%12 == 5) {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) + y*y
			if n <= int(N) && n%12 == 7 {
				isPrime[n] = !isPrime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= int(N) && n%12 == 11 {
				isPrime[n] = !isPrime[n]
			}
		}
	}

	for n = 5; float64(n) <= Sqrt; n++ {
		if isPrime[n] {
			for y = n * n; y < int(N); y += n * n {
				isPrime[y] = false
			}
		}
	}

	isPrime[2] = true
	isPrime[3] = true

	var primes []int
	for x = 0; x < len(isPrime)-1; x++ {
		if isPrime[x] {
			primes = append(primes, x)
		}
	}

	return primes
}
