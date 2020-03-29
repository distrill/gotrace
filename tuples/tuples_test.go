package tuples

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

/*
	Scenario: A tuple with w=1 is a point
	Given a ← tuple(4.3, -4.2, 3.1, 1)
	Then a.x = 4.3
	And a.y = -4.2
	And a.z = 3.1
	And a.w = 1
	And a is a point
	And a is not a vector
*/
func TestIsPoint(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1}
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 1.0)
	assert.Equal(t, a.isPoint(), true)
	assert.Equal(t, a.isVector(), false)
}

/*
	Scenario: A tuple with w=0 is a vector
	Given a ← tuple(4.3, -4.2, 3.1, 0)
	Then a.x = 4.3
	And a.y = -4.2
	And a.z = 3.1
	And a.w = 0
	And a is not a point
	And a is a vector
*/
func TestIsVector(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 0}
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 0.0)
	assert.Equal(t, a.isPoint(), false)
	assert.Equal(t, a.isVector(), true)
}

/*
	Scenario: point() creates tuples with w=1
	Given p ← point(4, -4, 3)
	Then p = tuple(4, -4, 3, 1)
*/
func TestNewPoint(t *testing.T) {
	p := NewPoint(4, -4, 3)
	assert.True(t, p.Equal(Tuple{4, -4, 3, 1}))
}

/*
	Scenario: vector() creates tuples with w=0
	Given v ← vector(4, -4, 3)
	Then v = tuple(4, -4, 3, 0)
*/
func TestNewVector(t *testing.T) {
	p := NewVector(4, -4, 3)
	assert.True(t, p.Equal(Tuple{4, -4, 3, 0}))
}

/*
	Scenario: Adding two tuples
	Given a1 ← tuple(3, -2, 5, 1)
	And a2 ← tuple(-2, 3, 1, 0)
	Then a1 + a2 = tuple(1, 1, 6, 1)
*/
func TestAddTuples(t *testing.T) {
	a1 := Tuple{3, -2, 5, 1}
	a2 := Tuple{-2, 3, 1, 0}
	assert.True(t, a1.Add(a2).Equal(Tuple{1, 1, 6, 1}))
}

/*
	Scenario: Subtracting two points
	Given p1 ← point(3, 2, 1)
	And p2 ← point(5, 6, 7)
	Then p1 - p2 = vector(-2, -4, -6)
*/
func TestSubTwoPoints(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	assert.True(t, p1.Sub(p2).Equal(NewVector(-2, -4, -6)))
}

/*
	Scenario: Subtracting a vector from a point
	Given p ← point(3, 2, 1)
	And v ← vector(5, 6, 7)
	Then p - v = point(-2, -4, -6)
*/
func TestSubVectorFromPoint(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)
	assert.True(t, p.Sub(v).Equal(NewPoint(-2, -4, -6)))
}

/*
	Scenario: Subtracting two vectors
	Given v1 ← vector(3, 2, 1)
	And v2 ← vector(5, 6, 7)
	Then v1 - v2 = vector(-2, -4, -6)
*/
func TestSubTwoVectors(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	assert.True(t, v1.Sub(v2).Equal(NewVector(-2, -4, -6)))
}

/*
	Scenario: Subtracting a vector from the zero vector
	Given zero ← vector(0, 0, 0)
	And v ← vector(1, -2, 3)
	Then zero - v = vector(-1, 2, -3)
*/
func TestSubVectorFromZero(t *testing.T) {
	zero := NewVector(0, 0, 0)
	v := NewVector(1, -2, 3)
	assert.True(t, zero.Sub(v).Equal(NewVector(-1, 2, -3)))
}

/*
	Scenario: Negating a tuple
	Given a ← tuple(1, -2, 3, -4)
	Then -a = tuple(-1, 2, -3, 4)
*/
func TestNegTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.True(t, a.Neg().Equal(Tuple{-1, 2, -3, 4}))
}

/*
	Scenario: Multiplying a tuple by a scalar
	Given a ← tuple(1, -2, 3, -4)
	Then a * 3.5 = tuple(3.5, -7, 10.5, -14)
*/
func TestMulTupleScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.True(t, a.Mul(3.5).Equal(Tuple{3.5, -7, 10.5, -14}))
}

/*
	Scenario: Multiplying a tuple by a fraction
	Given a ← tuple(1, -2, 3, -4)
	Then a * 0.5 = tuple(0.5, -1, 1.5, -2)
*/
func TestMulTupleFraction(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.True(t, a.Mul(0.5).Equal(Tuple{0.5, -1, 1.5, -2}))
}

/*
	Scenario: Dividing a tuple by a scalar
	Given a ← tuple(1, -2, 3, -4)
	Then a / 2 = tuple(0.5, -1, 1.5, -2)
*/
func TestDivTupleScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	assert.True(t, a.Div(2).Equal(Tuple{0.5, -1, 1.5, -2}))
}

/*
	Scenario: Computing the magnitude of vector(1, 0, 0)
	Given v ← vector(1, 0, 0)
	Then magnitude(v) = 1
*/
func TestMagnitudeVector100(t *testing.T) {
	v := NewVector(1, 0, 0)
	assert.Equal(t, v.Mag(), 1.0)
}

/*
	Scenario: Computing the magnitude of vector(0, 1, 0)
	Given v ← vector(0, 1, 0)
	Then magnitude(v) = 1
*/
func TestMagnitudeVector010(t *testing.T) {
	v := NewVector(0, 1, 0)
	assert.Equal(t, v.Mag(), 1.0)
}

/*
	Scenario: Computing the magnitude of vector(0, 0, 1)
	Given v ← vector(0, 0, 1)
	Then magnitude(v) = 1
*/
func TestMagnitudeVector001(t *testing.T) {
	v := NewVector(0, 0, 1)
	assert.Equal(t, v.Mag(), 1.0)
}

/*
	Scenario: Computing the magnitude of vector(1, 2, 3)
	Given v ← vector(1, 2, 3)
	Then magnitude(v) = √14
*/
func TestMagnitudeVectorPos(t *testing.T) {
	v := NewVector(1, 2, 3)
	assert.Equal(t, v.Mag(), math.Sqrt(14))
}

/*
	Scenario: Computing the magnitude of vector(-1, -2, -3)
	Given v ← vector(-1, -2, -3)
	Then magnitude(v) = √14
*/
func TestMagnitudeVectorNeg(t *testing.T) {
	v := NewVector(-1, -2, -3)
	assert.Equal(t, v.Mag(), math.Sqrt(14))
}

/*
	Scenario: Normalizing vector(4, 0, 0) gives (1, 0, 0)
	Given v ← vector(4, 0, 0)
	Then normalize(v) = vector(1, 0, 0)
*/
func TestNormVector400(t *testing.T) {
	v := NewVector(4, 0, 0)
	assert.True(t, v.Norm().Equal(NewVector(1, 0, 0)))
}

/*
	Scenario: Normalizing vector(1, 2, 3)
	Given v ← vector(1, 2, 3)
					# vector(1/√14,   2/√14,   3/√14)
	Then normalize(v) = approximately vector(0.26726, 0.53452, 0.80178)
*/
func TestNormVector123(t *testing.T) {
	v := NewVector(1, 2, 3)
	assert.True(t, v.Norm().Equal(NewVector(0.26726, 0.53452, 0.80178)))
}

/*
	Scenario: The magnitude of a normalized vector
	Given v ← vector(1, 2, 3)
	When norm ← normalize(v)
	Then magnitude(norm) = 1
*/
func TestMagNormalizedVector(t *testing.T) {
	v := NewVector(1, 2, 3)
	n := v.Norm()
	assert.Equal(t, n.Mag(), 1.0)
}

/*
	Scenario: The dot product of two tuples
	Given a ← vector(1, 2, 3)
	And b ← vector(2, 3, 4)
	Then dot(a, b) = 20
*/
func TestDotTwoTupples(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	assert.Equal(t, a.Dot(b), 20.0)
}

/*
	Scenario: The cross product of two vectors
	Given a ← vector(1, 2, 3)
	And b ← vector(2, 3, 4)
	Then cross(a, b) = vector(-1, 2, -1)
	And cross(b, a) = vector(1, -2, 1)
*/
func TestCrossTwoVectors(t *testing.T) {
	a := NewVector(1, 2, 3)
	b := NewVector(2, 3, 4)
	assert.True(t, a.Cross(b).Equal(NewVector(-1, 2, -1)))
	assert.True(t, b.Cross(a).Equal(NewVector(1, -2, 1)))
}
