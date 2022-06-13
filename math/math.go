package math

import "math"

func Max[T int](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxFloat(a, b float64) float64 {
	return math.Max(a, b)
}
