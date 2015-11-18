package hilbertcurve

import "testing"

type fixture struct {
	n int64
	d int64
	x int64
	y int64
}

// Fixtures, these run in both d => x, y and x, y => d
var fixtures []fixture = []fixture{
	{n: 4, d: 15, x: 3, y: 0},
	{n: 4, d: 5, x: 0, y: 3},
	{n: 4, d: 2, x: 1, y: 1},
	{n: 4, d: 0, x: 0, y: 0},
}

// Test if rx and ry are both 1 or 0 then the return is x and y
func TestFlipQuadrantSame(t *testing.T) {

	var n int64 = 20

	var x int64 = 10
	var y int64 = 10

	var rx int64 = 1
	var ry int64 = 1

	// Expect x and y to remain the same
	resX, resY := flipQuadrant(n, x, y, rx, ry)

	if resX != x || resY != y {
		t.Errorf("Error flipping quadrant. Got %d %d, expected %d %d",
			resX, resY, x, y)
	}

	// Test the same is true with 0s
	rx, ry = 0, 0

	// Expect x and y to remain the same
	resX, resY = flipQuadrant(n, x, y, rx, ry)

	if resX != x || resY != y {
		t.Errorf("Error flipping quadrant. Got %d %d, expected %d %d",
			resX, resY, x, y)
	}
}

// Test conversion from bool to int64 works as expected. So basic but
// testing for the sake of it
func TestBoolToInt64(t *testing.T) {

	if res := boolToInt64(true); res != int64(1) {
		t.Errorf("Expected true to convert to int64(1), got %v", res)
	}

	if res := boolToInt64(false); res != int64(0) {
		t.Errorf("Expected true to convert to int64(1), got %v", res)
	}
}

// Run the fixtures to convert distance into coordinates
func TestDistanceToCoordinates(t *testing.T) {

	for _, f := range fixtures {
		resX, resY := DistanceToCoordinates(f.n, f.d)
		if resX != f.x || resY != f.y {
			t.Errorf("Error converting distance to coordinates. Got %d %d, expected %d %d",
				resX, resY, f.x, f.y)
		}
	}
}

// Runs fixtures converting coordinates to distances
func TestCoordinatesToDistance(t *testing.T) {

	for _, f := range fixtures {
		if dist := CoordinatesToDistance(f.n, f.x, f.y); dist != f.d {
			t.Errorf("Distance %d does not match expected %d", dist, f.d)
		}
	}
}
