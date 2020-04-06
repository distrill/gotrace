package main

import (
	"fmt"
)

type projectile struct {
	position Tuple
	velocity Tuple
}

type environment struct {
	gravity Tuple
	wind    Tuple
}

func tick(e environment, p projectile) projectile {
	pos := p.position.Add(Tuple(p.velocity))
	vel := p.velocity.Add(Tuple(e.gravity))
	return projectile{pos, vel}
}

func main() {
	// start the ray at z = -5
	rayOrigin := NewPoint(0, 0, -5)
	// put the wall at z = 10
	wallZ := 10.0
	wallSize := 7.0
	canvasPixels := 200

	w := canvasPixels
	h := canvasPixels

	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := NewCanvas(w, h)
	color := Red
	shape := NewSphere()

	i := 0
	t := canvasPixels * canvasPixels
	percentCount := 1
	for y := 0; y < canvasPixels-1; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels-1; x++ {
			i += 1
			if i%(t/100) == 0 {
				fmt.Printf("\r%3d/100", percentCount)
				percentCount += 1
			}
			worldX := -half + pixelSize*float64(x)

			// describe point on the wall that the ray will target
			position := NewPoint(worldX, worldY, wallZ)

			r := Ray{rayOrigin, position.Sub(rayOrigin).Norm()}
			xs, err := shape.Intersect(r)
			if err != nil {
				panic(err)
			}

			if xs.Hit() != nil {
				canvas.WritePixel(
					x,
					y,
					color,
				)
			}
		}
	}
	fmt.Print("\r100/100\r       \r")

	err := canvas.ToPPM("bh.ppm")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("done!")
}
