package main

import (
	"math"
)

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() Material {
	return Material{
		Color:     Color{1, 1, 1},
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}
}

func reflect(l, n Tuple) Tuple {
	return l.Reflect(n)
}

func (m Material) Lighting(light PointLight, position, eyev, normv Tuple) Color {
	// combine surface color with light's color/intensity
	effectiveColor := m.Color.MulC(light.Intensity)

	// find direction to the light source
	lightv := light.Position.Sub(position).Norm()

	// compute ambient contribution
	ambient := effectiveColor.MulS(m.Ambient)
	var diffuse, specular Color

	// lightDotNorm is cosine of angle btn lightv and normv
	// negative means light is on other side of surface
	lightDotNorm := lightv.Dot(normv)

	if lightDotNorm < 0 {
		diffuse = Black
		specular = Black
	} else {
		// compute diffuse contribution
		diffuse = effectiveColor.MulS(m.Diffuse).MulS(lightDotNorm)

		// reflectDotEye is cosine of angle btn reflectv and eyev
		// negative means light reflects away from eye
		// NOTE the pseudo code had lightv negated here. that seems to just
		// break everything. it works not-negated. is this a real bug somewhere else?
		reflectv := lightv.Reflect(normv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye <= 0 {
			specular = Black
		} else {
			// compute specular contribution
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.MulS(m.Specular).MulS(factor)
		}

	}
	return ambient.Add(diffuse).Add(specular)
}
