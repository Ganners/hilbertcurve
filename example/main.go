package main

import (
	"fmt"

	"github.com/Ganners/hilbertcurve"
)

func main() {

	var n int64 = 8 // 8 x 8 grid
	var x int64 = 5
	var y int64 = 0

	// Compute the distance from the given x and y
	distance := hilbertcurve.CoordinatesToDistance(n, x, y)

	// Compute the x and y from our previously computed distance
	cx, cy := hilbertcurve.DistanceToCoordinates(n, distance)

	if x != cx || y != cy {
		fmt.Println("This aint good...")
	}

	fmt.Printf("Computed distance: %d\nComputed X: %d\nComputed Y: %d", distance, cx, cy)
}
