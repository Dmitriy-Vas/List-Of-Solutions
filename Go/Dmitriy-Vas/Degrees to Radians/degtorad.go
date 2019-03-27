package Degrees_To_Radian

import (
	"math"
)

func DegreesToRad(n float64) float64 {
	return n * (math.Pi / 180)
}

func RadiansToDeg(n float64) float64 {
	return n * (180 / math.Pi)
}
