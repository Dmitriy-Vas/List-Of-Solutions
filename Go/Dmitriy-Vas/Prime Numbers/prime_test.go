package Prime_Numbers

import (
	"math"
	"testing"
)

func TestPrime(t *testing.T) {
	primes := Prime(math.MaxInt32 / 100)
	if len(primes) != 1358124 {
		t.Fatal(len(primes))
	}
}
