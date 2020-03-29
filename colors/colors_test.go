package colors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
	Scenario: Colors are (red, green, blue) tuples
	Given c ← color(-0.5, 0.4, 1.7)
	Then c.red = -0.5
	And c.green = 0.4
	And c.blue = 1.7
*/
func TestColor(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}
	assert.Equal(t, c.Red, -0.5)
	assert.Equal(t, c.Green, 0.4)
	assert.Equal(t, c.Blue, 1.7)
}

/*
	Scenario: Adding colors
	Given c1 ← color(0.9, 0.6, 0.75)
	And c2 ← color(0.7, 0.1, 0.25)
	Then c1 + c2 = color(1.6, 0.7, 1.0)
*/
func TestAddColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	assert.True(t, c1.Add(c2).Equal(Color{1.6, 0.7, 1.0}))
}

/*
	Scenario: Subtracting colors
	Given c1 ← color(0.9, 0.6, 0.75)
	And c2 ← color(0.7, 0.1, 0.25)
	Then c1 - c2 = color(0.2, 0.5, 0.5)
*/
func TestSubColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	assert.True(t, c1.Sub(c2).Equal(Color{0.2, 0.5, 0.5}))
}

/*
	Scenario: Multiplying a color by a scalar
	Given c ← color(0.2, 0.3, 0.4)
	Then c * 2 = color(0.4, 0.6, 0.8)
*/
func TestMulS(t *testing.T) {
	c := Color{0.2, 0.3, 0.4}
	assert.True(t, c.MulS(2).Equal(Color{0.4, 0.6, 0.8}))
}

/*
	Scenario: Multiplying colors
	Given c1 ← color(1, 0.2, 0.4)
	And c2 ← color(0.9, 1, 0.1)
	Then c1 * c2 = color(0.9, 0.2, 0.04)
*/
func TestMulC(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}
	assert.True(t, c1.MulC(c2).Equal(Color{0.9, 0.2, 0.04}))
}
