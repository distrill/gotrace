package main

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

func (r Ray) Position(t float64) Tuple {
	return r.Origin.Add(r.Direction.Mul(t))
}

func (r Ray) Transform(m Matrix) Ray {
	o := m.MustMulT(r.Origin)
	d := m.MustMulT(r.Direction)
	return Ray{o, d}
}

func (r Ray) Translate(x, y, z float64) Ray {
	o := NewTransform(r.Origin).Translate(x, y, z).Value()
	d := NewTransform(r.Direction).Translate(x, y, z).Value()
	return Ray{o, d}
}

func (r Ray) Scale(x, y, z float64) Ray {
	o := NewTransform(r.Origin).Scale(x, y, z).Value()
	d := NewTransform(r.Direction).Scale(x, y, z).Value()
	return Ray{o, d}
}
