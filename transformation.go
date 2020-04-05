package main

import (
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
	t Tuple
	m Matrix
}

// method chaining
// not very idiomatic, particularly here we encourage panic
func NewTransform(t Tuple) Transform {
	return Transform{t, NewIdentityMatrix(4)}
}

func (m Matrix) Translate(x, y, z float64) Matrix {
	return NewTranslation(x, y, z).MustMulM(m)
}

func (c Transform) Translate(x, y, z float64) Transform {
	return Transform{c.t, c.m.Translate(x, y, z)}
}

func (m Matrix) Scale(x, y, z float64) Matrix {
	return NewScaling(x, y, z).MustMulM(m)
}

func (c Transform) Scale(x, y, z float64) Transform {
	return Transform{c.t, c.m.Scale(x, y, z)}
}

func (m Matrix) RotateX(radians float64) Matrix {
	return NewRotationX(radians).MustMulM(m)
}

func (c Transform) RotateX(radians float64) Transform {
	return Transform{c.t, c.m.RotateX(radians)}
}

func (m Matrix) RotateY(radians float64) Matrix {
	return NewRotationY(radians).MustMulM(m)
}

func (c Transform) RotateY(radians float64) Transform {
	return Transform{c.t, c.m.RotateY(radians)}
}

func (m Matrix) RotateZ(radians float64) Matrix {
	return NewRotationZ(radians).MustMulM(m)
}

func (c Transform) RotateZ(radians float64) Transform {
	return Transform{c.t, c.m.RotateZ(radians)}
}

func (m Matrix) Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	return NewShearing(xy, xz, yx, yz, zx, zy).MustMulM(m)
}

func (c Transform) Shear(xy, xz, yx, yz, zx, zy float64) Transform {
	return Transform{c.t, c.m.Shear(xy, xz, yx, yz, zx, zy)}
}

func (c Transform) Value() Tuple {
	return c.m.MustMulT(c.t)
}
