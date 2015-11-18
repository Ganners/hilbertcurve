// Implementation of the hilbert curve, converting from coordinates to
// distance and from distance to coordinates.
package hilbertcurve

// CoordinatesToDistance converts a given X and Y into a distance D. The square
// division is set by N
func CoordinatesToDistance(n, x, y int64) (d int64) {
	var rx, ry int64
	for s := n / 2; s > 0; s /= 2 {
		rx = boolToInt64((x & s) > 0)
		ry = boolToInt64((y & s) > 0)
		d += s * s * ((3 * rx) ^ ry)
		x, y = flipQuadrant(n, x, y, rx, ry)
	}
	return d
}

// DistanceToCoordinates converts a given distance D into coordinates X and Y
func DistanceToCoordinates(n, d int64) (x, y int64) {
	var rx, ry, s int64
	t := d
	for s = 1; s < n; s *= 2 {
		rx = 1 & (t / 2)
		ry = 1 & (t ^ rx)
		x, y = flipQuadrant(s, x, y, rx, ry)
		x += s * rx
		y += s * ry
		t /= 4
	}
	return x, y
}

// Flips a quadrant
func flipQuadrant(n, x, y, rx, ry int64) (int64, int64) {
	if ry == 0 {
		if rx == 1 {
			x = n - 1 - x
			y = n - 1 - y
		}
		x, y = y, x
	}
	return x, y
}

// Converts a boolean to an integer
func boolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}
