package matrix

import (
	"github.com/distrill/gotrace/tuples"
	"math"
)

func NewTranslation(x, y, z float64) Matrix {
	t := NewIdentityMatrix(4)
	t[0][3] = x
	t[1][3] = y
	t[2][3] = z
	return t
}

func NewScaling(x, y, z float64) Matrix {
	s := NewIdentityMatrix(4)
	s[0][0] = x
	s[1][1] = y
	s[2][2] = z
	return s
}

func NewRotationX(radians float64) Matrix {
	r := NewIdentityMatrix(4)
	r[1][1] = math.Cos(radians)
	r[2][2] = math.Cos(radians)
	r[1][2] = -math.Sin(radians)
	r[2][1] = math.Sin(radians)
	return r
}

func NewRotationY(radians float64) Matrix {
	r := NewIdentityMatrix(4)
	r[0][0] = math.Cos(radians)
	r[2][2] = math.Cos(radians)
	r[0][2] = math.Sin(radians)
	r[2][0] = -math.Sin(radians)
	return r
}

func NewRotationZ(radians float64) Matrix {
	r := NewIdentityMatrix(4)
	r[0][0] = math.Cos(radians)
	r[1][1] = math.Cos(radians)
	r[0][1] = -math.Sin(radians)
	r[1][0] = math.Sin(radians)
	return r
}

func NewShearing(xy, xz, yx, yz, zx, zy float64) Matrix {
	s := NewIdentityMatrix(4)
	s[0][1] = xy
	s[0][2] = xz
	s[1][0] = yx
	s[1][2] = yz
	s[2][0] = zx
	s[2][1] = zy
	return s
}

type Transform struct {
	t tuples.Tuple
	m Matrix
}

// method chaining is not very idiomatic, particularly here we encourage panic
func NewTransform(t tuples.Tuple) Transform {
	return Transform{t, NewIdentityMatrix(4)}
}

func (c Transform) Translate(x, y, z float64) Transform {
	c.m = NewTranslation(x, y, z).MustMulM(c.m)
	return c
}

func (c Transform) Scale(x, y, z float64) Transform {
	c.m = NewScaling(x, y, z).MustMulM(c.m)
	return c
}

func (c Transform) RotateX(radians float64) Transform {
	c.m = NewRotationX(radians).MustMulM(c.m)
	return c
}

func (c Transform) RotateY(radians float64) Transform {
	c.m = NewRotationY(radians).MustMulM(c.m)
	return c
}

func (c Transform) RotateZ(radians float64) Transform {
	c.m = NewRotationZ(radians).MustMulM(c.m)
	return c
}

func (c Transform) Shear(xy, xz, yx, yz, zx, zy float64) Transform {
	c.m = NewShearing(xy, xz, yx, yz, zx, zy).MustMulM(c.m)
	return c
}

func (c Transform) Value() tuples.Tuple {
	return c.m.MustMulT(c.t)
}
