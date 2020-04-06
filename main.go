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
	wallSize := 10.0
	canvasPixels := 400

	w := canvasPixels
	h := canvasPixels

	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	canvas := NewCanvas(w, h)
	// color := Red

	shape1 := NewSphere().WithTransform(NewScaling(0.6, 0.6, 1).Translate(0.7, 0.7, 0))
	shape1.Material.Color = Color{0.3, 0.4, 0.8}

	shape2 := NewSphere().WithTransform(NewScaling(0.6, 0.6, 1).Translate(-0.5, -0.5, 0))
	shape2.Material.Color = Color{0.8, 0.2, 0.3}

	shapes := []Sphere{shape1, shape2}

	// shape2

	// light source
	lightPosition := NewPoint(-10, 10, -10)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

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

			for _, shape := range shapes {
				xs, err := shape.Intersect(r)
				if err != nil {
					panic(err)
				}

				if hit := xs.Hit(); hit != nil {
					/*
						point ← position(ray, hit.t)
						normal ← normal_at(hit.object, point)
						eye ← -ray.direction
						color ← lighting(hit.object.material, light, point, eye, normal)
					*/
					point := r.Position(hit.T)
					norm := hit.Object.NormalAt(point)
					eye := r.Direction
					color := hit.Object.Material.Lighting(light, point, eye, norm)
					canvas.WritePixel(
						x,
						y,
						color,
					)
				}
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
