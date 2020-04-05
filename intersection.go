package main

type Intersection struct {
	T      float64
	Object Sphere
}

type Intersections []Intersection

func (is Intersections) Hit() *Intersection {
	var hit *Intersection
	for i := range is {
		if is[i].T > 0 && (hit == nil || is[i].T < hit.T) {
			hit = &is[i]
		}
	}

	return hit
}
