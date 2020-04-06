package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Scenario: The default material
	Given m ← material()
	Then m.color = color(1, 1, 1)
	And m.ambient = 0.1
	And m.diffuse = 0.9
	And m.specular = 0.9
	And m.shininess = 200.0
*/
func TestDefaultMaterial(t *testing.T) {
	m := NewMaterial()
	assert.Equal(t, Color{1, 1, 1}, m.Color)
	assert.Equal(t, 0.1, m.Ambient)
	assert.Equal(t, 0.9, m.Diffuse)
	assert.Equal(t, 0.9, m.Specular)
	assert.Equal(t, 200.0, m.Shininess)
}

/*
	Background:

	Given m ← material()
	And position ← point(0, 0, 0)
*/

/*
	Scenario: Lighting with the eye between the light and the surface
	Given eyev ← vector(0, 0, -1)
	And normalv ← vector(0, 0, -1)
	And light ← point_light(point(0, 0, -10), color(1, 1, 1))
	When result ← lighting(m, light, position, eyev, normalv)
	Then result = color(1.9, 1.9, 1.9)
*/
func LightingEyeBetweenLightSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -1), Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normv)
	assert.True(t, result.Equal(Color{1.9, 1.9, 1.9}))
}

/*
	Scenario: Lighting with the eye between light and surface, eye offset 45°
	Given eyev ← vector(0, √2/2, -√2/2)
	And normalv ← vector(0, 0, -1)
	And light ← point_light(point(0, 0, -10), color(1, 1, 1))
	When result ← lighting(m, light, position, eyev, normalv)
	Then result = color(1.0, 1.0, 1.0)
*/
func LightingEyeBetweenLightSurfaceEyeOffset45(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, math.Sqrt2/2, -math.Sqrt2/2)
	normv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, -10), Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normv)
	assert.True(t, Color{1, 1, 1}.Equal(result))
}

/*
	Scenario: Lighting with eye opposite surface, light offset 45°
	Given eyev ← vector(0, 0, -1)
	And normalv ← vector(0, 0, -1)
	And light ← point_light(point(0, 10, -10), color(1, 1, 1))
	When result ← lighting(m, light, position, eyev, normalv)
	Then result = color(0.7364, 0.7364, 0.7364)
*/
func LightingEyeOppositeSurphaseLightOffset45(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normv)
	assert.True(t, Color{0.7364, 0.7364, 0.7364}.Equal(result))
}

/*
	Scenario: Lighting with eye in the path of the reflection vector
	Given eyev ← vector(0, -√2/2, -√2/2)
	And normalv ← vector(0, 0, -1)
	And light ← point_light(point(0, 10, -10), color(1, 1, 1))
	When result ← lighting(m, light, position, eyev, normalv)
	Then result = color(1.6364, 1.6364, 1.6364)
*/
func TestLightingEyeInPathReflectionVector(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 10, -10), Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normv)
	assert.True(t, Color{1.6364, 1.6364, 1.6364}.Equal(result))
}

/*
	Scenario: Lighting with the light behind the surface
	Given eyev ← vector(0, 0, -1)
	And normalv ← vector(0, 0, -1)
	And light ← point_light(point(0, 0, 10), color(1, 1, 1))
	When result ← lighting(m, light, position, eyev, normalv)
	Then result = color(0.1, 0.1, 0.1)
*/
func TestLightingLightBehindSurface(t *testing.T) {
	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	eyev := NewVector(0, 0, -1)
	normv := NewVector(0, 0, -1)
	light := PointLight{NewPoint(0, 0, 10), Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normv)
	fmt.Println(result)
	assert.True(t, Color{0.1, 0.1, 0.1}.Equal(result))
}
