package intutils

import "math"

// Min returns the lower of two input values
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the higher of two inputs
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value of the input
func Abs(i int) int {
	if i <= math.MinInt64 {
		panic("Abs overflow")
	}
	return Max(i, -i)
}
