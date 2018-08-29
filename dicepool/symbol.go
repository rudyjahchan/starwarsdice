package starwarsdice

// Symbol is the display shape of a die
type Symbol string

// Cube is for d6
const Cube = "■"

// Diamond is for d8
const Diamond = "⬥"

// Hexagon is for d12
const Hexagon = "⬣"

func (s Symbol) String() string {
	return string(s)
}
