package colors

import (
	"math"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

// duplicated from tuple, should be pulled out when we have a better idea what the abstraction looks like
const epsilon = 0.00001

// duplicated from tuple, should be pulled out when we have a better idea what the abstraction looks like
func (c Color) Equal(o Color) bool {
	return (math.Abs(c.Red-o.Red) < epsilon &&
		math.Abs(c.Green-o.Green) < epsilon &&
		math.Abs(c.Blue-o.Blue) < epsilon)
}

// duplicated from tuple, should be pulled out when we have a better idea what the abstraction looks like
func (c Color) Add(o Color) Color {
	return Color{c.Red + o.Red, c.Green + o.Green, c.Blue + o.Blue}
}

// duplicated from tuple, should be pulled out when we have a better idea what the abstraction looks like
func (c Color) Sub(o Color) Color {
	return Color{c.Red - o.Red, c.Green - o.Green, c.Blue - o.Blue}
}

// duplicated from tuple, should be pulled out when we have a better idea what the abstraction looks like
func (c Color) MulS(s float64) Color {
	return Color{c.Red * s, c.Green * s, c.Blue * s}
}

func (c Color) MulC(o Color) Color {
	return Color{c.Red * o.Red, c.Green * o.Green, c.Blue * o.Blue}
}
