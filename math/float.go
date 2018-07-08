package math

import (
	"math"
)

func RoundFloat64(f float64, place int) float64 {
	shift := math.Pow10(place)

	return math.Floor(f*shift+.5) / shift
}

func FloorFloat64(f float64, place int) float64 {
	shift := math.Pow10(place)

	return math.Floor(f*shift+.1) / shift
}
