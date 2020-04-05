package main

import (
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
	s.SetTransform(transform)
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
