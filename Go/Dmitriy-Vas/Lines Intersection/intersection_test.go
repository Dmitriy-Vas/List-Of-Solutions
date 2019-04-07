package Lines_Intersection

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	l1 := CreateLine(Point{4, 0}, Point{6, 10})
	l2 := CreateLine(Point{0, 3}, Point{10, 7})
	if result, err := Intersection(l1, l2); err != nil {
		t.Fatal(result, err)
	} else {
		fmt.Println(result)
	}
}
