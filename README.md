Hilbert Curve Package
=====================

Implementation of the hilbert curve, both in conversion from coordinates (x and
y) into distance, and the reverse from distance into x and y.

All functions take int64 arguments (and return likewise), this is to
make the bitwise operations somewhat safer and allow for larger
precision.

Install:

    go get -u github.com/Ganners/hilbertcurve

Example usage:

    var n int64 = 8 // 8 x 8 grid
    var x int64 = 5
    var y int64 = 0

    // Compute the distance from the given x and y
    distance := hilbertcurve.CoordinatesToDistance(n, x, y)

    // Compute the x and y from our previously computed distance
    cx, cy := hilbertcurve.DistanceToCoordinates(n, distance)

    if x != cx || y != cy {
        log.Println("This aint good...")
    }
