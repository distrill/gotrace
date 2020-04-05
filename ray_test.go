package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Scenario: Creating and querying a ray
	Given origin ← point(1, 2, 3)
	And direction ← vector(4, 5, 6)
	When r ← ray(origin, direction)
	Then r.origin = origin
	And r.direction = direction
*/
func TestCreateQueryRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	r := Ray{origin, direction}
	assert.Equal(t, r.Origin, origin)
	assert.Equal(t, r.Direction, direction)
}

/*
	Scenario: Computing a point from a distance
	Given r ← ray(point(2, 3, 4), vector(1, 0, 0))
	Then position(r, 0) = point(2, 3, 4)
	And position(r, 1) = point(3, 3, 4)
	And position(r, -1) = point(1, 3, 4)
	And position(r, 2.5) = point(4.5, 3, 4)
*/
func TestCompoutePointFromDistance(t *testing.T) {
	r := Ray{NewPoint(2, 3, 4), NewVector(1, 0, 0)}
	assert.Equal(t, r.Position(0), NewPoint(2, 3, 4))
	assert.Equal(t, r.Position(1), NewPoint(3, 3, 4))
	assert.Equal(t, r.Position(-1), NewPoint(1, 3, 4))
	assert.Equal(t, r.Position(2.5), NewPoint(4.5, 3, 4))
}

/*
	Scenario: Translating a ray
	Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	And m ← translation(3, 4, 5)
	When r2 ← transform(r, m)
	Then r2.origin = point(4, 6, 8)
	And r2.direction = vector(0, 1, 0)
*/
func TranslatingARay(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	m := NewTranslation(3, 4, 5)
	r2 := r.Transform(m)
	assert.Equal(t, NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, NewVector(0, 1, 0), r2.Direction)
}

/*
	Scenario: Scaling a ray
	Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	And m ← scaling(2, 3, 4)
	When r2 ← transform(r, m)
	Then r2.origin = point(2, 6, 12)
	And r2.direction = vector(0, 3, 0)
*/
func ScalingARay(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	m := NewScaling(2, 3, 4)
	r2 := r.Transform(m)
	assert.Equal(t, NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, NewVector(0, 3, 0), r2.Direction)
}

/*
	Scenario: Translating a ray
	Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	And m ← translation(3, 4, 5)
	When r2 ← transform(r, m)
	Then r2.origin = point(4, 6, 8)
	And r2.direction = vector(0, 1, 0)
*/
func TranslatingARayFluid(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	r2 := r.Translate(3, 4, 5)
	assert.Equal(t, NewPoint(4, 6, 8), r2.Origin)
	assert.Equal(t, NewVector(0, 1, 0), r2.Direction)
}

/*
	Scenario: Scaling a ray
	Given r ← ray(point(1, 2, 3), vector(0, 1, 0))
	And m ← scaling(2, 3, 4)
	When r2 ← transform(r, m)
	Then r2.origin = point(2, 6, 12)
	And r2.direction = vector(0, 3, 0)
*/
func ScalingARayFluid(t *testing.T) {
	r := Ray{NewPoint(1, 2, 3), NewVector(0, 1, 0)}
	r2 := r.Scale(2, 3, 4)
	assert.Equal(t, NewPoint(2, 6, 12), r2.Origin)
	assert.Equal(t, NewVector(0, 3, 0), r2.Direction)
}
