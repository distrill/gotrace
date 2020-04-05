package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Scenario: An intersection encapsulates t and object
	Given s ← sphere()
	When i ← intersection(3.5, s)
	Then i.t = 3.5
	And i.object = s
*/
func TestIntersecionCreateQuery(t *testing.T) {
	s := Sphere{}
	i := Intersection{3.5, s}
	assert.Equal(t, 3.5, i.T)
	assert.Equal(t, s, i.Object)
}

/*
	Scenario: Aggregating intersections
	Given s ← sphere()
	And i1 ← intersection(1, s)
	And i2 ← intersection(2, s)
	When xs ← intersections(i1, i2)
	Then xs.count = 2
	And xs[0].t = 1
	And xs[1].t = 2
*/
func TestAggregateIntersections(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections{i1, i2}
	assert.Equal(t, 2, len(xs))
	assert.Equal(t, 1.0, xs[0].T)
	assert.Equal(t, 2.0, xs[1].T)
}

/*
	Scenario: The hit, when all intersections have positive t
	Given s ← sphere()
	And i1 ← intersection(1, s)
	And i2 ← intersection(2, s)
	And xs ← intersections(i2, i1)
	When i ← hit(xs)
	Then i = i1
*/
func TestHitAllPositive(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections{i2, i1}
	i := xs.Hit()
	assert.Equal(t, i1, *i)
}

/*
	Scenario: The hit, when some intersections have negative t
	Given s ← sphere()
	And i1 ← intersection(-1, s)
	And i2 ← intersection(1, s)
	And xs ← intersections(i2, i1)
	When i ← hit(xs)
	Then i = i2
*/
func TestHitSomeNegative(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := Intersections{i2, i1}
	i := xs.Hit()
	assert.Equal(t, i2, *i)
}

/*
	Scenario: The hit, when all intersections have negative t
	Given s ← sphere()
	And i1 ← intersection(-2, s)
	And i2 ← intersection(-1, s)
	And xs ← intersections(i2, i1)
	When i ← hit(xs)
	Then i is nothing
*/
func TestHitAllNegative(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := Intersections{i2, i1}
	i := xs.Hit()
	assert.Nil(t, i)
}

/*
	Scenario: The hit is always the lowest nonnegative intersection
	Given s ← sphere()
	And i1 ← intersection(5, s)
	And i2 ← intersection(7, s)
	And i3 ← intersection(-3, s)
	And i4 ← intersection(2, s)
	And xs ← intersections(i1, i2, i3, i4)
	When i ← hit(xs)
	Then i = i4
*/
func TestHitIsLowestNonNegative(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{5, s}
	i2 := Intersection{7, s}
	i3 := Intersection{-3, s}
	i4 := Intersection{2, s}
	xs := Intersections{i1, i2, i3, i4}
	i := xs.Hit()
	assert.Equal(t, i4, *i)
}
