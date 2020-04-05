package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	Scenario: Multiplying by a translation matrix
	Given transform ← translation(5, -3, 2)
	And p ← point(-3, 4, 5)
	Then transform * p = point(2, 1, 7)
*/
func TestMulTranslationMatrix(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	p := NewPoint(-3, 4, 5)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(2, 1, 7)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Multiplying by the inverse of a translation matrix
	Given transform ← translation(5, -3, 2)
	And inv ← inverse(transform)
	And p ← point(-3, 4, 5)
	Then inv * p = point(-8, 7, 3)
*/
func TestMulTranslationMatrixInverse(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	inv, err := transform.Inverse()
	require.Nil(t, err)
	p := NewPoint(-3, 4, 5)
	actual, err := inv.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(-8, 7, 3)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Translation does not affect vectors
	Given transform ← translation(5, -3, 2)
	And v ← vector(-3, 4, 5)
	Then transform * v = v
*/
func TestMulTranslationNotAffectVectors(t *testing.T) {
	transform := NewTranslation(5, -3, 2)
	v := NewVector(-3, 4, 5)
	actual, err := transform.MulT(v)
	require.Nil(t, err)
	assert.Equal(t, v, actual)
}

/*
	Scenario: A scaling matrix applied to a point
	Given transform ← scaling(2, 3, 4)
	And p ← point(-4, 6, 8)
	Then transform * p = point(-8, 18, 32)
*/
func TestMulScalingMatrixPoint(t *testing.T) {
	transform := NewScaling(2, 3, 4)
	p := NewPoint(-4, 6, 8)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(-8, 18, 32)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: A scaling matrix applied to a vector
	Given transform ← scaling(2, 3, 4)
	And v ← vector(-4, 6, 8)
	Then transform * v = vector(-8, 18, 32)
*/
func TestMulScalingMatrixVector(t *testing.T) {
	transform := NewScaling(2, 3, 4)
	v := NewVector(-4, 6, 8)
	actual, err := transform.MulT(v)
	require.Nil(t, err)
	expected := NewVector(-8, 18, 32)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Multiplying by the inverse of a scaling matrix
	Given transform ← scaling(2, 3, 4)
	And inv ← inverse(transform)
	And v ← vector(-4, 6, 8)
	Then inv * v = vector(-2, 2, 2)
*/
func TestMulScalingMatrixInverse(t *testing.T) {
	transform := NewScaling(2, 3, 4)
	inv, err := transform.Inverse()
	require.Nil(t, err)
	v := NewVector(-4, 6, 8)
	actual, err := inv.MulT(v)
	require.Nil(t, err)
	expected := NewVector(-2, 2, 2)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Reflection is scaling by a negative value
	Given transform ← scaling(-1, 1, 1)
	And p ← point(2, 3, 4)
	Then transform * p = point(-2, 3, 4)
*/
func TestReflectNegativeScaling(t *testing.T) {
	transform := NewScaling(-1, 1, 1)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(-2, 3, 4)
	assert.Equal(t, expected, actual)
}

/*
	Scenario: Rotating a point around the x axis
	Given p ← point(0, 1, 0)
	And half_quarter ← rotation_x(π / 4)
	And full_quarter ← rotation_x(π / 2)
	Then half_quarter * p = point(0, √2/2, √2/2)
	And full_quarter * p = point(0, 0, 1)
*/
func TestRotateAroundX(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationX(math.Pi / 4)
	fullQuarter := NewRotationX(math.Pi / 2)
	hqa, err := halfQuarter.MulT(p)
	require.Nil(t, err)
	fqa, err := fullQuarter.MulT(p)
	require.Nil(t, err)
	hqe := NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)
	fqe := NewPoint(0, 0, 1)
	assert.True(t, hqe.Equal(hqa))
	assert.True(t, fqe.Equal(fqa))
}

/*
	Scenario: The inverse of an x-rotation rotates in the opposite direction
	Given p ← point(0, 1, 0)
	And half_quarter ← rotation_x(π / 4)
	And inv ← inverse(half_quarter)
	Then inv * p = point(0, √2/2, -√2/2)
*/
func TestInverseRotateX(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationX(math.Pi / 4)
	inv, err := halfQuarter.Inverse()
	require.Nil(t, err)
	actual, err := inv.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: Rotating a point around the y axis
	Given p ← point(0, 0, 1)
	And half_quarter ← rotation_y(π / 4)
	And full_quarter ← rotation_y(π / 2)
	Then half_quarter * p = point(√2/2, 0, √2/2)
	And full_quarter * p = point(1, 0, 0)
*/
func TestRotateAroundY(t *testing.T) {
	p := NewPoint(0, 0, 1)
	halfQuarter := NewRotationY(math.Pi / 4)
	fullQuarter := NewRotationY(math.Pi / 2)
	hqa, err := halfQuarter.MulT(p)
	require.Nil(t, err)
	fqa, err := fullQuarter.MulT(p)
	require.Nil(t, err)
	hqe := NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)
	fqe := NewPoint(1, 0, 0)
	assert.True(t, hqe.Equal(hqa))
	assert.True(t, fqe.Equal(fqa))
}

/*
	Scenario: Rotating a point around the z axis
	Given p ← point(0, 1, 0)
	And half_quarter ← rotation_z(π / 4)
	And full_quarter ← rotation_z(π / 2)
	Then half_quarter * p = point(-√2/2, √2/2, 0)
	And full_quarter * p = point(-1, 0, 0)
*/
func TestRotateAroundZ(t *testing.T) {
	p := NewPoint(0, 1, 0)
	halfQuarter := NewRotationZ(math.Pi / 4)
	fullQuarter := NewRotationZ(math.Pi / 2)
	hqa, err := halfQuarter.MulT(p)
	require.Nil(t, err)
	fqa, err := fullQuarter.MulT(p)
	require.Nil(t, err)
	hqe := NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0)
	fqe := NewPoint(-1, 0, 0)
	assert.True(t, hqe.Equal(hqa))
	assert.True(t, fqe.Equal(fqa))
}

/*
	Scenario: A shearing transformation moves x in proportion to y
	Given transform ← shearing(1, 0, 0, 0, 0, 0)
	And p ← point(2, 3, 4)
	Then transform * p = point(5, 3, 4)
*/
func TestShearXPropY(t *testing.T) {
	transform := NewShearing(1, 0, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(5, 3, 4)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: A shearing transformation moves x in proportion to z
	Given transform ← shearing(0, 1, 0, 0, 0, 0)
	And p ← point(2, 3, 4)
	Then transform * p = point(6, 3, 4)
*/
func TestShearXPropZ(t *testing.T) {
	transform := NewShearing(0, 1, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(6, 3, 4)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: A shearing transformation moves y in proportion to x
	Given transform ← shearing(0, 0, 1, 0, 0, 0)
	And p ← point(2, 3, 4)
	Then transform * p = point(2, 5, 4)
*/
func TestShearYPropX(t *testing.T) {
	transform := NewShearing(0, 0, 1, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(2, 5, 4)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: A shearing transformation moves y in proportion to z
	Given transform ← shearing(0, 0, 0, 1, 0, 0)
	And p ← point(2, 3, 4)
	Then transform * p = point(2, 7, 4)
*/
func TestShearYPropz(t *testing.T) {
	transform := NewShearing(0, 0, 0, 1, 0, 0)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(2, 7, 4)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: A shearing transformation moves z in proportion to x
	Given transform ← shearing(0, 0, 0, 0, 1, 0)
	And p ← point(2, 3, 4)
	Then transform * p = point(2, 3, 6)
*/
func TestShearZPropX(t *testing.T) {
	transform := NewShearing(0, 0, 0, 0, 1, 0)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(2, 3, 6)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: A shearing transformation moves z in proportion to y
	Given transform ← shearing(0, 0, 0, 0, 0, 1)
	And p ← point(2, 3, 4)
	Then transform * p = point(2, 3, 7)
*/
func TestShearZPropY(t *testing.T) {
	transform := NewShearing(0, 0, 0, 0, 0, 1)
	p := NewPoint(2, 3, 4)
	actual, err := transform.MulT(p)
	require.Nil(t, err)
	expected := NewPoint(2, 3, 7)
	assert.True(t, expected.Equal(actual))
}

/*
	Scenario: Individual transformations are applied in sequence
	Given p ← point(1, 0, 1)
	And A ← rotation_x(π / 2)
	And B ← scaling(5, 5, 5)
	// And C ← translation(10, 5, 7)
	# apply rotation first
	When p2 ← A * p
	Then p2 = point(1, -1, 0)
	# then apply scaling
	When p3 ← B * p2
	Then p3 = point(5, -5, 0)
	# then apply translation
	When p4 ← C * p3
	Then p4 = point(15, 0, 7)
*/
func TestIndividualTransformsAppliedInSequence(t *testing.T) {
	p := NewPoint(1, 0, 1)
	A := NewRotationX(math.Pi / 2)
	B := NewScaling(5, 5, 5)
	C := NewTranslation(10, 5, 7)

	// apply first rotation
	p2, err := A.MulT(p)
	require.Nil(t, err)
	assert.True(t, p2.Equal(NewPoint(1, -1, 0)))

	// then apply scaling
	p3, err := B.MulT(p2)
	require.Nil(t, err)
	assert.True(t, p3.Equal(NewPoint(5, -5, 0)))

	// then apply translation
	p4, err := C.MulT(p3)
	require.Nil(t, err)
	assert.True(t, p4.Equal(NewPoint(15, 0, 7)))
}

/*
	Scenario: Chained transformations must be applied in reverse order
	Given p ← point(1, 0, 1)
	And A ← rotation_x(π / 2)
	And B ← scaling(5, 5, 5)
	And C ← translation(10, 5, 7)
	When T ← C * B * A
	Then T * p = point(15, 0, 7)
*/
func TestChainedTransformationReverseOrder(t *testing.T) {
	p := NewPoint(1, 0, 1)
	A := NewRotationX(math.Pi / 2)
	B := NewScaling(5, 5, 5)
	C := NewTranslation(10, 5, 7)

	T := C.MustMulM(B).MustMulM(A)
	Tp, err := T.MulT(p)
	require.Nil(t, err)

	assert.Equal(t, Tp, NewPoint(15, 0, 7))
}

/*
	Scenario: Chained transformations must be applied in reverse order
	Given p ← point(1, 0, 1)
	And A ← rotation_x(π / 2)
	And B ← scaling(5, 5, 5)
	And C ← translation(10, 5, 7)
	When T ← C * B * A
	Then T * p = point(15, 0, 7)
*/
func TestFluidApiMethodChainOrder(t *testing.T) {
	tp := NewTransform(NewPoint(1, 0, 1)).
		RotateX(math.Pi/2).
		Scale(5, 5, 5).
		Translate(10, 5, 7).
		Value()

	assert.Equal(t, tp, NewPoint(15, 0, 7))
}
