package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	Scenario: A ray intersects a sphere at two points
	Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	And s ← sphere()
	When xs ← intersect(s, r)
	Then xs.count = 2
	And xs[0] = 4.0
	And xs[1] = 6.0
*/
func TestRayIntersectSphere2Points(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 4.0, xs[0].T)
	assert.Equal(t, 6.0, xs[1].T)
}

/*
	Scenario: A ray intersects a sphere at a tangent
	Given r ← ray(point(0, 1, -5), vector(0, 0, 1))
	And s ← sphere()
	When xs ← intersect(s, r)
	Then xs.count = 2
	And xs[0] = 5.0
	And xs[1] = 5.0
*/
func TestRayIntersectSphereTangent(t *testing.T) {
	r := Ray{NewPoint(0, 1, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 5.0, xs[0].T)
	assert.Equal(t, 5.0, xs[1].T)
}

/*
	Scenario: A ray misses a sphere
	Given r ← ray(point(0, 2, -5), vector(0, 0, 1))
	And s ← sphere()
	When xs ← intersect(s, r)
	Then xs.count = 0
*/
func TestRayMissesShpere(t *testing.T) {
	r := Ray{NewPoint(0, 2, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 0, len(xs))
}

/*
	Scenario: A ray originates inside a sphere
	Given r ← ray(point(0, 0, 0), vector(0, 0, 1))
	And s ← sphere()
	When xs ← intersect(s, r)
	Then xs.count = 2
	And xs[0] = -1.0
	And xs[1] = 1.0
*/
func TestRayOriginatesInsideSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, 0), NewVector(0, 0, 1)}
	s := NewSphere()
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -1.0, xs[0].T)
	assert.Equal(t, 1.0, xs[1].T)
}

/*
	Scenario: A sphere is behind a ray
	Given r ← ray(point(0, 0, 5), vector(0, 0, 1))
	And s ← sphere()
	When xs ← intersect(s, r)
	Then xs.count = 2
	And xs[0] = -6.0
	And xs[1] = -4.0
*/
func TestSphereBehindRay(t *testing.T) {
	r := Ray{NewPoint(0, 0, 5), NewVector(0, 0, 1)}
	s := NewSphere()
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, -6.0, xs[0].T)
	assert.Equal(t, -4.0, xs[1].T)
}

/*
	Scenario: Intersect sets the object on the intersection
	Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	And s ← sphere()
	When xs ← intersect(s, r)
	Then xs.count = 2
	And xs[0].object = s
	And xs[1].object = s
*/
func TestIntersectSetsIntersectionObject(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere()
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, xs[0].Object, s)
	assert.Equal(t, xs[1].Object, s)
}

/*
	Scenario: A sphere's default transformation
	Given s ← sphere()
	Then s.transform = identity_matrix
*/
func TestSphereDefaultTransformation(t *testing.T) {
	s := NewSphere()
	assert.Equal(t, NewIdentityMatrix(4), s.Transform)
}

/*
	Scenario: Changing a sphere's transformation
	Given s ← sphere()
	And t ← translation(2, 3, 4)
	When set_transform(s, t)
	Then s.transform = t
*/
func TestSetSphereTransformation(t *testing.T) {
	s := NewSphere()
	transform := NewTranslation(2, 3, 4)
	s.Transform = transform
	assert.Equal(t, transform, s.Transform)
}

/*
	Scenario: Changing a sphere's transformation fluid api
	Given t ← translation(2, 3, 4)
	and s ← sphere().set_transform(t)
	Then s.transform = t
*/
func TestSphereWithTransformation(t *testing.T) {
	transform := NewTranslation(2, 3, 4)
	s := NewSphere().WithTransform(transform)
	assert.Equal(t, transform, s.Transform)
}

/*
	Scenario: Intersecting a scaled sphere with a ray
	Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	And s ← sphere()
	When set_transform(s, scaling(2, 2, 2))
	And xs ← intersect(s, r)
	Then xs.count = 2
	And xs[0].t = 3
	And xs[1].t = 7
*/
func TestIntersectScaledSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere().WithTransform(NewScaling(2, 2, 2))
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 3.0, xs[0].T)
	assert.Equal(t, 7.0, xs[1].T)
}

/*
	Scenario: Intersecting a translated sphere with a ray
	Given r ← ray(point(0, 0, -5), vector(0, 0, 1))
	And s ← sphere()
	When set_transform(s, translation(5, 0, 0))
	And xs ← intersect(s, r)
	Then xs.count = 0
*/
func TestIntersectTranslatedSphere(t *testing.T) {
	r := Ray{NewPoint(0, 0, -5), NewVector(0, 0, 1)}
	s := NewSphere().WithTransform(NewTranslation(5, 0, 0))
	xs, err := s.Intersect(r)
	require.Nil(t, err)
	assert.Equal(t, 0, len(xs))
}

/*
	Scenario: The normal on a sphere at a point on the x axis
	Given s ← sphere()
	When n ← normal_at(s, point(1, 0, 0))
	Then n = vector(1, 0, 0)
*/
func TestNormalPointXAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(1, 0, 0))
	assert.True(t, n.Equal(NewVector(1, 0, 0)))
}

/*
	Scenario: The normal on a sphere at a point on the y axis
	Given s ← sphere()
	When n ← normal_at(s, point(0, 1, 0))
	Then n = vector(0, 1, 0)
*/
func TestNormalPointYAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(0, 1, 0))
	assert.True(t, n.Equal(NewVector(0, 1, 0)))
}

/*
	Scenario: The normal on a sphere at a point on the z axis
	Given s ← sphere()
	When n ← normal_at(s, point(0, 0, 1))
	Then n = vector(0, 0, 1)
*/
func TestNormalPointZAxis(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(0, 0, 1))
	assert.True(t, n.Equal(NewVector(0, 0, 1)))
}

/*
	Scenario: The normal on a sphere at a nonaxial point
	Given s ← sphere()
	When n ← normal_at(s, point(√3/3, √3/3, √3/3))
	Then n = vector(√3/3, √3/3, √3/3)
*/
func TestNormalPointNonaxial(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	assert.True(t, n.Equal(NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)))
}

/*
	Scenario: The normal is a normalized vector
	Given s ← sphere()
	When n ← normal_at(s, point(√3/3, √3/3, √3/3))
	Then n = normalize(n)
*/
func TestNormalPointIsNormalized(t *testing.T) {
	s := NewSphere()
	n := s.NormalAt(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	normd := n.Norm()
	assert.True(t, n.Equal(normd))
}

/*
	Scenario: Computing the normal on a translated sphere
	Given s ← sphere()
	And set_transform(s, translation(0, 1, 0))
	When n ← normal_at(s, point(0, 1.70711, -0.70711))
	Then n = vector(0, 0.70711, -0.70711)
*/
func TestNormalPointTranslatedSphere(t *testing.T) {
	s := NewSphere().WithTransform(NewTranslation(0, 1, 0))
	n := s.NormalAt(NewPoint(0, 1.70711, -0.70711))
	assert.True(t, n.Equal(NewVector(0, 0.70711, -0.70711)))
}

/*
	Scenario: Computing the normal on a transformed sphere
	Given s ← sphere()
	And m ← scaling(1, 0.5, 1) * rotation_z(π/5)
	And set_transform(s, m)
	When n ← normal_at(s, point(0, √2/2, -√2/2))
	Then n = vector(0, 0.97014, -0.24254)
*/
func TestNormalPointTransformedSphere(t *testing.T) {
	s := NewSphere().WithTransform(NewRotationZ(math.Pi/5).Scale(1, 0.5, 1))
	n := s.NormalAt(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2))
	assert.True(t, n.Equal(NewVector(0, 0.97014, -0.24254)))
}

/*
	Scenario: A sphere has a default material
	Given s ← sphere()
	When m ← s.material
	Then m = material()
*/
func TestSphereDefaultMaterial(t *testing.T) {
	s := NewSphere()
	m := s.Material
	assert.Equal(t, NewMaterial(), m)
}

/*
	Scenario: A sphere may be assigned a material
	Given s ← sphere()
	And m ← material()
	And m.ambient ← 1
	When s.material ← m
	Then s.material = m
*/
func TestSphereAssignMaterial(t *testing.T) {
	s := NewSphere()
	m := NewMaterial()
	m.Ambient = 1
	s.Material = m
	assert.Equal(t, m, s.Material)
}

/*
	Scenario: A sphere may be assigned a material Fluid API
	Given s ← sphere()
	And m ← material()
	And m.ambient ← 1
	When s.material ← m
	Then s.material = m
*/
func TestSphereAssignMaterialFluid(t *testing.T) {
	m := NewMaterial()
	m.Ambient = 1
	s := NewSphere().WithMaterial(m)
	assert.Equal(t, m, s.Material)
}
