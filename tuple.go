package main

import (
	"math"
)

func square(x float64) float64 {
	return math.Pow(x, 2)
}

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

const epsilon = 0.00001

func (t Tuple) Equal(o Tuple) bool {
	return (math.Abs(t.X-o.X) < epsilon &&
		math.Abs(t.Y-o.Y) < epsilon &&
		math.Abs(t.Z-o.Z) < epsilon &&
		// W should never be manipulated
		t.W == o.W)
}

func (t Tuple) isPoint() bool {
	return t.W == 1.0
}

func (t Tuple) isVector() bool {
	return t.W == 0.0
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) Add(o Tuple) Tuple {
	return Tuple{t.X + o.X, t.Y + o.Y, t.Z + o.Z, t.W + o.W}
}

func (t Tuple) Sub(o Tuple) Tuple {
	return Tuple{t.X - o.X, t.Y - o.Y, t.Z - o.Z, t.W - o.W}
}

func (t Tuple) Neg() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}

}

func (t Tuple) Mul(s float64) Tuple {
	return Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func (t Tuple) Div(v float64) Tuple {
	return t.Mul(1 / v)
}

func (t Tuple) Mag() float64 {
	return math.Sqrt(square(t.X) + square(t.Y) + square(t.Z) + square(t.W))
}

func (t Tuple) Norm() Tuple {
	return Tuple{
		X: t.X / t.Mag(),
		Y: t.Y / t.Mag(),
		Z: t.Z / t.Mag(),
		W: t.W / t.Mag(),
	}
}

func (t Tuple) Dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z + t.W*o.W
}

func (t Tuple) Cross(o Tuple) Tuple {
	if !t.isVector() {
		// err handle pls
		panic("Cross product only supported for vectors")
	}
	if !o.isVector() {
		// err handle pls
		panic("Cross product only supported for vectors")
	}
	return NewVector(
		(t.Y*o.Z)-(t.Z*o.Y),
		(t.Z*o.X)-(t.X*o.Z),
		(t.X*o.Y)-(t.Y*o.X),
	)
}

func (t Tuple) Reflect(n Tuple) Tuple {
	return t.Sub(n.Mul(2 * t.Dot(n)))
}
