package canvas

import (
	"bufio"
	"os"
	"testing"

	"github.com/distrill/gotrace/colors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
	Scenario: Creating a canvas
	Given c ← canvas(10, 20)
	Then c.width = 10
	And c.height = 20
	And every pixel of c is color(0, 0, 0)
*/
func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	assert.Equal(t, c.Width, 10)
	assert.Equal(t, c.Height, 20)
	for _, row := range c.Pixels {
		for _, p := range row {
			assert.True(t, p.Equal(colors.Black))
		}
	}
}

/*
	Scenario: Writing pixels to a canvas
	Given c ← canvas(10, 20)
	And red ← color(1, 0, 0)
	When write_pixel(c, 2, 3, red)
	Then pixel_at(c, 2, 3) = red
*/
func TestWritePixelsToCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	c.WritePixel(2, 3, colors.Red)
	assert.True(t, c.PixelAt(2, 3).Equal(colors.Red))
}

/*
	Scenario: Constructing the PPM header
	Given c ← canvas(5, 3)
	When ppm ← canvas_to_ppm(c)
	Then lines 1-3 of ppm are
		"""
		P3
		5 3
		255
		"""
*/
func TestPPMHeader(t *testing.T) {
	fn := "test.ppm"
	c := NewCanvas(5, 3)
	err := c.ToPPM(fn)
	require.Nil(t, err)

	file, err := os.Open(fn)
	require.Nil(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	assert.Equal(t, scanner.Text(), "P3")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "5 3")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "255")

	require.Nil(t, scanner.Err())
	require.Nil(t, os.Remove(fn))
}

/*
	Scenario: Constructing the PPM pixel data
	Given c ← canvas(5, 3)
	And c1 ← color(1.5, 0, 0)
	And c2 ← color(0, 0.5, 0)
	And c3 ← color(-0.5, 0, 1)
	When write_pixel(c, 0, 0, c1)
	And write_pixel(c, 2, 1, c2)
	And write_pixel(c, 4, 2, c3)
	And ppm ← canvas_to_ppm(c)
	Then lines 4-6 of ppm are
		"""
		255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
		"""
*/
func TestPPMPixels(t *testing.T) {
	fn := "test.ppm"
	c := NewCanvas(5, 3)
	c1 := colors.Color{1.5, 0, 0}
	c2 := colors.Color{0, 0.5, 0}
	c3 := colors.Color{-0.5, 0, 1}
	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	err := c.ToPPM(fn)
	require.Nil(t, err)

	file, err := os.Open(fn)
	require.Nil(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scanner.Scan()
	scanner.Scan()

	scanner.Scan()
	assert.Equal(t, scanner.Text(), "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255")

	require.Nil(t, scanner.Err())
	require.Nil(t, os.Remove(fn))
}

/*
	Scenario: Splitting long lines in PPM files
	Given c ← canvas(10, 2)
	When every pixel of c is set to color(1, 0.8, 0.6)
	And ppm ← canvas_to_ppm(c)
	Then lines 4-7 of ppm are
		"""
		255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
		153 255 204 153 255 204 153 255 204 153 255 204 153
		255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
		153 255 204 153 255 204 153 255 204 153 255 204 153
		"""
*/
func TestPPMLongLines(t *testing.T) {
	fn := "test.ppm"
	c := NewCanvas(10, 2)
	c1 := colors.Color{1, 0.8, 0.6}
	for x := 0; x < 10; x++ {
		for y := 0; y < 2; y++ {
			c.WritePixel(x, y, c1)
		}
	}

	err := c.ToPPM(fn)
	require.Nil(t, err)

	file, err := os.Open(fn)
	require.Nil(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scanner.Scan()
	scanner.Scan()

	scanner.Scan()
	assert.Equal(t, scanner.Text(), "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "153 255 204 153 255 204 153 255 204 153 255 204 153")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204")
	scanner.Scan()
	assert.Equal(t, scanner.Text(), "153 255 204 153 255 204 153 255 204 153 255 204 153")

	require.Nil(t, scanner.Err())
	require.Nil(t, os.Remove(fn))
}
