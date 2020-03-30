package canvas

import (
	"fmt"
	"math"
	"os"

	"github.com/distrill/gotrace/colors"
)

// Canvas - grid of pixels representing image
type Canvas struct {
	Width  int
	Height int
	Pixels [][]colors.Color
}

// NewCanvas - initialize pixels grid to given width and heigh
func NewCanvas(width, height int) Canvas {
	pixels := make([][]colors.Color, width)
	for i := 0; i < width; i++ {
		pixels[i] = make([]colors.Color, height)
	}
	return Canvas{width, height, pixels}
}

// WritePixel - paint a given pixel a given color
func (c *Canvas) WritePixel(x, y int, col colors.Color) {
	if x < 0 {
		x = 0
	}
	if x >= c.Width {
		x = c.Width - 1
	}
	if y < 0 {
		y = 0
	}
	if y >= c.Height {
		y = c.Height - 1
	}

	c.Pixels[x][y] = col
}

// PixelAt - return pixel color at given coordinates
func (c Canvas) PixelAt(x, y int) colors.Color {
	return c.Pixels[x][y]
}

func getPixelValue(p float64) float64 {
	if p <= 0 {
		return 0
	}
	if p >= 1 {
		return 255
	}
	return math.Round(p * 255)
}

// ToPPM - write canvas to given file name in PPM format
func (c Canvas) ToPPM(fn string) error {
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	// write header
	f.WriteString("P3\n")
	f.WriteString(fmt.Sprintf("%v %v\n", c.Width, c.Height))
	f.WriteString("255\n")
	f.Sync()

	// write pixels
	p := 1
	for y := 0; y < c.Height; y++ {
		w := 0
		if y%int(c.Height/10) == 0 {
			fmt.Printf("\rwriting %v%%", p*10)
			p += 1
		}
		for x := 0; x < c.Width; x++ {
			color := c.PixelAt(x, y)

			// write red, adding newline if length would have exceeded 70 char
			r := fmt.Sprintf("%v", getPixelValue(color.Red))
			if w+len(r) >= 70 {
				f.WriteString("\n")
				w = 0
			} else if w > 0 {
				f.WriteString(" ")
				w++
			}
			f.WriteString(r)
			w += len(r)

			// write green, adding newline if length would have exceeded 70 char
			g := fmt.Sprintf("%v", getPixelValue(color.Green))
			if w+len(r) >= 70 {
				f.WriteString("\n")
				w = 0
			} else if w > 0 {
				f.WriteString(" ")
				w++
			}
			f.WriteString(g)
			w += len(r)

			// write blue , adding newline if length would have exceeded 70 char
			b := fmt.Sprintf("%v", getPixelValue(color.Blue))
			if w+len(r) >= 70 {
				f.WriteString("\n")
				w = 0
			} else if w > 0 {
				f.WriteString(" ")
				w++
			}
			f.WriteString(b)
			w += len(r)
		}
		f.WriteString("\n")
	}
	fmt.Print("\n")
	f.WriteString("\n")
	f.Sync()

	return nil
}
