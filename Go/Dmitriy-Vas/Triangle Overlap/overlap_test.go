package Triangle_Overlap

import (
	"testing"
)

func TestOverlap(t *testing.T) {
	t1 := &triangle{point{0.0, 0.0}, point{5.0, 0.0}, point{0.0, 5.0}}
	t2 := &triangle{point{0.0, 0.0}, point{5.0, 0.0}, point{0.0, 6.0}}
	overlapping := triTri2D(t1, t2, 0.0, false, true)
	if !overlapping {
		t.Fatal(overlapping)
	}

	t1 = &triangle{point{0.0, 0.0}, point{0.0, 5.0}, point{5.0, 0.0}}
	t2 = t1
	overlapping = triTri2D(t1, t2, 0.0, true, true)
	if !overlapping {
		t.Fatal(overlapping)
	}

	t1 = &triangle{point{0.0, 0.0}, point{5.0, 0.0}, point{0.0, 5.0}}
	t2 = &triangle{point{-10.0, 0.0}, point{-5.0, 0.0}, point{-1.0, 6.0}}
	overlapping = triTri2D(t1, t2, 0.0, false, true)
	if overlapping {
		t.Fatal(overlapping)
	}

	t1.p3 = point{2.5, 5.0}
	t2 = &triangle{point{0.0, 4.0}, point{2.5, -1.0}, point{5.0, 4.0}}
	overlapping = triTri2D(t1, t2, 0.0, false, true)
	if !overlapping {
		t.Fatal(overlapping)
	}

	t1 = &triangle{point{0.0, 0.0}, point{1.0, 1.0}, point{0.0, 2.0}}
	t2 = &triangle{point{2.0, 1.0}, point{3.0, 0.0}, point{3.0, 2.0}}
	overlapping = triTri2D(t1, t2, 0.0, false, true)
	if overlapping {
		t.Fatal(overlapping)
	}

	t2 = &triangle{point{2.0, 1.0}, point{3.0, -2.0}, point{3.0, 4.0}}
	overlapping = triTri2D(t1, t2, 0.0, false, true)
	if overlapping {
		t.Fatal(overlapping)
	}

	t1 = &triangle{point{0.0, 0.0}, point{1.0, 0.0}, point{0.0, 1.0}}
	t2 = &triangle{point{1.0, 0.0}, point{2.0, 0.0}, point{1.0, 1.1}}
	overlapping = triTri2D(t1, t2, 0.0, false, true)
	if !overlapping {
		t.Fatal(overlapping)
	}
}
