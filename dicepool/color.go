package starwarsdice

// Color represents the shade of a die
type Color uint

// Black is a color
const Black Color = 0

// Blue is a color
const Blue Color = 1

// Green is a color
const Green Color = 2

// Yellow is a color
const Yellow Color = 3

// Purple is a color
const Purple Color = 4

// Red is a color
const Red Color = 5

const RESET Color = 66

func (c Color) String() string {
	switch c {
	case Black:
		return "\x1B[30;1m"
	case Blue:
		return "\x1B[36m"
	case Green:
		return "\x1B[32m"
	case Yellow:
		return "\x1B[33m"
	case Purple:
		return "\x1B[35;1m"
	case Red:
		return "\x1B[31m"
	default:
		return "\x1B[0m"
	}
}
