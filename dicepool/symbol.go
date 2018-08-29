package starwarsdice

type Symbol string

const Cube = "■"
const Diamond = "⬥"
const Hexagon = "⬣"

func (s Symbol) String() string {
	return string(s)
}
