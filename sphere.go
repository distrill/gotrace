package main

import (
	"math"
)

type Sphere struct {
	Transform Matrix
	Material  Material
}

func NewSphere() Sphere {
	return Sphere{NewIdentityMatrix(4), NewMaterial()}
}

func (s Sphere) WithTransform(t Matrix) Sphere {
	return Sphere{t, NewMaterial()}
}

func (s Sphere) WithMaterial(m Material) Sphere {
	return Sphere{s.Transform, m}
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

func (s Sphere) NormalAt(p Tuple) Tuple {
	objectPoint := s.Transform.MustInverse().MustMulT(p)
	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := s.Transform.MustInverse().MustTranspose().MustMulT(objectNormal)
	worldNormal.W = 0
	return worldNormal.Norm()
}
