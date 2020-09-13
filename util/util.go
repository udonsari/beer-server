package util

import "math"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Floor(val float64, dim int) float64 {
	tenExp := 1.0
	for i := 0; i < dim; i++ {
		tenExp *= 10
	}
	return math.Floor(val*tenExp) / tenExp
}
