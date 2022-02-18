package qt

import (
	"math"
)

func Radians(deg float32) float32 {
	return (deg * math.Pi / 180.0)
}

func Degrees(rad float32) float32 {
	return (rad * 180.0 / math.Pi)
}