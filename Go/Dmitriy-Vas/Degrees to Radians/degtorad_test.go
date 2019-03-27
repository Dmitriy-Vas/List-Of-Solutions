package Degrees_To_Radian

import (
	"testing"
)

var arr = []struct{
	deg float64
	rad float64
}{
	{57.29577951308232, 1},
	{286.4788975654116, 5},
}

func TestDegreesToRad(t *testing.T) {
	for _, d := range arr {
		if DegreesToRad(d.deg) != d.rad {
			t.Fatal(d)
		}
	}
}

func TestRadiansToDeg(t *testing.T) {
	for _, d := range arr {
		if RadiansToDeg(d.rad) != d.deg {
			t.Fatal(d)
		}
	}
}
