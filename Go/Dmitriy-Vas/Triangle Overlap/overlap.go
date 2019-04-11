package Triangle_Overlap

type point struct {
	x, y float64
}

type triangle struct {
	p1, p2, p3 point
}

func (t *triangle) det2D() float64 {
	return t.p1.x * (t.p2.y - t.p3.y) +
		t.p2.x * (t.p3.y - t.p1.y) +
		t.p3.x * (t.p1.y - t.p2.y)
}

func (t *triangle) checkTriWinding(allowReversed bool) {
	detTri := t.det2D()
	if detTri < 0.0 {
		if allowReversed {
			a := t.p3
			t.p3 = t.p2
			t.p2 = a
		} else {
			panic("Triangle has wrong winding direction.")
		}
	}
}

func boundaryCollideChk(t *triangle, eps float64) bool {
	return t.det2D() < eps
}

func boundaryDoesntCollideChk(t *triangle, eps float64) bool {
	return t.det2D() <= eps
}

func triTri2D(t1, t2 *triangle, eps float64, allowReversed, onBoundary bool) bool {
	// Triangles must be expressed anti-clockwise.
	t1.checkTriWinding(allowReversed)
	t2.checkTriWinding(allowReversed)

	// 'onBoundary' determines whether points on boundary are considered as colliding or not.
	var chkEdge func (*triangle, float64) bool
	if onBoundary {
		chkEdge = boundaryCollideChk
	} else {
		chkEdge = boundaryDoesntCollideChk
	}
	lp1 := [3]point{t1.p1, t1.p2, t1.p3}
	lp2 := [3]point{t2.p1, t2.p2, t2.p3}

	// for each edge E of t1
	for i := 0; i < 3; i++ {
		j := (i + 1) % 3
		// Check all points of t2 lay on the external side of edge E.
		// If they do, the triangles do not overlap.
		tri1 := &triangle{lp1[i], lp1[j], lp2[0]}
		tri2 := &triangle{lp1[i], lp1[j], lp2[1]}
		tri3 := &triangle{lp1[i], lp1[j], lp2[2]}
		if chkEdge(tri1, eps) && chkEdge(tri2, eps) && chkEdge(tri3, eps) {
			return false
		}
	}

	// for each edge E of t2
	for i := 0; i < 3; i++ {
		j := (i + 1) % 3
		// Check all points of t1 lay on the external side of edge E.
		// If they do, the triangles do not overlap.
		tri1 := &triangle{lp2[i], lp2[j], lp1[0]}
		tri2 := &triangle{lp2[i], lp2[j], lp1[1]}
		tri3 := &triangle{lp2[i], lp2[j], lp1[2]}
		if chkEdge(tri1, eps) && chkEdge(tri2, eps) && chkEdge(tri3, eps) {
			return false
		}
	}

	// The triangles overlap.
	return true
}
