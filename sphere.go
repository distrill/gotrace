package main

import (
	"math"
)

type Sphere struct {
	Transform Matrix
}

func NewSphere() Sphere {
	return Sphere{NewIdentityMatrix(4)}
}

func (s Sphere) WithTransform(t Matrix) Sphere {
	return Sphere{t}
}

func (s *Sphere) SetTransform(t Matrix) {
	s.Transform = t
}

func (s Sphere) Intersect(r Ray) (Intersections, error) {
	tm, err := s.Transform.Inverse()
	if err != nil {
		return nil, err
	}
	r = r.Transform(tm)
	sphereToRay := r.Origin.Sub(NewPoint(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := r.Direction.Dot(sphereToRay) * 2.0
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := (b * b) - (4 * a * c)

	if discriminant < 0 {
		return Intersections{}, nil
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	return Intersections{Intersection{t1, s}, Intersection{t2, s}}, nil
}
