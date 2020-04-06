package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	Scenario: A point light has a position and intensity
	Given intensity ← color(1, 1, 1)
	And position ← point(0, 0, 0)
	When light ← point_light(position, intensity)
	Then light.position = position
	And light.intensity = intensity
*/
func TestPointLightPositionIntensity(t *testing.T) {
	intensity := Color{1, 1, 1}
	position := NewPoint(0, 0, 0)
	light := PointLight{position, intensity}
	assert.True(t, light.Position.Equal(position))
	assert.Equal(t, intensity, light.Intensity)
}
