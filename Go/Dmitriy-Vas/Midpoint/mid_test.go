package Midpoint

import "testing"

var points = []struct{
	p1 Point
	p2 Point
	mid Point
}{
	{Point{2,2}, Point{4, 4}, Point{3, 3}},
	{Point{4, 4}, Point{6, 6}, Point{5, 5}},
}

func TestMid(t *testing.T) {
	for _, d := range points {
		if Mid(d.p1, d.p2) != d.mid {
			t.Fatal(d)
		}
	}
}
