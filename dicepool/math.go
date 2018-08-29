package starwarsdice

import "math"

func Abs(c int) uint {
	return uint(math.Abs(float64(c)))
}
