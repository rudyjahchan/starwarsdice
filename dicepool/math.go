package starwarsdice

import "math"

// Abs exists because go only seems to have a float64 version, geez
func Abs(c int) uint {
	return uint(math.Abs(float64(c)))
}
