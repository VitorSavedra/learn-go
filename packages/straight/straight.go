package main

import "math"

// Point represents a coordinate in the Cartesian plane
type Point struct {
	x float64
	y float64
}

func cathetus(a, b Point) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distance is responsible for calculate the distance between two points
func Distance(a, b Point) float64 {
	cx, cy := cathetus(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
