package main

import (
	"fmt"
	// "math"
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
	// ray_origin ← point(0, 0, -5)
	rayOrigin := NewPoint(0, 0, -5)
	// put the wall at z = 10
	// wall_z ← 10
	wallZ := 10.0
	wallSize := 7
	canvasPixels := 200

	w := canvasPixels
	h := canvasPixels

	pixelSize := float64(wallSize) / float64(canvasPixels)
	fmt.Println(pixelSize)
	half := float64(wallSize / 2)

	canvas := NewCanvas(w, h)
	color := Red
	shape := NewSphere().WithTransform(NewTranslation(0.15, -0.15, 0))

	i := 0
	t := canvasPixels * canvasPixels
	whatever := 1
	for y := 0; y < canvasPixels-1; y++ {
		// compute world y
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels-1; x++ {
			i += 1
			// p := i / t
			if i%(t/100) == 0 {
				fmt.Printf("\r%02d/ 100", whatever)
				whatever += 1
			}
			// compute world x
			worldX := -half + pixelSize*float64(x)

			// describe point on the wall that the ray will target
			position := NewPoint(worldX, worldY, wallZ)
			// fmt.Println(position)

			r := Ray{rayOrigin, position.Sub(rayOrigin).Norm()}
			xs, err := shape.Intersect(r)
			if err != nil {
				panic(err)
			}

			// fmt.Println(x, y)
			if xs.Hit() != nil {
				canvas.WritePixel(
					x,
					y,
					color,
				)
			} else {
				// fmt.Println(xs)
			}
		}
	}
	fmt.Print("\r100/100\r       \r")

	err := canvas.ToPPM("bh.ppm")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("done!")

	/*
		canvas ← canvas(canvas_pixels, canvas_pixels)
		color ← color(1, 0, 0) # red
		shape ← sphere()
		# for each row of pixels in the canvas
		for y ← 0 to canvas_pixels - 1
			# compute the world y coordinate (top = +half, bottom = -half)
			world_y ← half - pixel_size * y
			# for each pixel in the row
			for x ← 0 to canvas_pixels - 1
				# compute the world x coordinate (left = -half, right = half)
				world_x ← -half + pixel_size * x

				# describe the point on the wall that the ray will target
				position ← point(world_x, world_y, wall_z)

				r ← ray(ray_origin, normalize(position - ray_origin))
				xs ← intersect(shape, r)

				if hit(xs) is defined
					write_pixel(canvas, x, y, color)
					end if
			end for
		end for
	*/

	/*
		w := 100
		h := 100
		center := 50.0
		radius := center * .8
		c := NewCanvas(w, h)

		p := NewPoint(0, -radius, 0)

		for i := 0; i <= 12; i++ {
			p = NewTransform(p).
				RotateZ(math.Pi / 6).
				Value()

			c.WritePixel(
				int(center-math.Round(p.X)),
				int(center-math.Round(p.Y)),
				Red,
			)
		}

		err := c.ToPPM("bh.ppm")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("done!")
	*/

	/*
		start := NewPoint(1, 1.5, 0)
		velocity := NewVector(1, 2, 0).Norm().Mul(11.25)
		p := projectile{start, velocity}

		gravity := NewVector(0, -0.11, 0)
		wind := NewVector(-0.01, 0, 0)
		e := environment{gravity, wind}

		w := 900
		h := 550
		c := NewCanvas(w, h)

		ts := 0
		for {
			if p.position.Y <= 0 {
				break
			}
			c.WritePixel(
				int(math.Round(p.position.X)),
				int(math.Round(p.position.Y)),
				Red,
			)
			p = tick(e, p)
			ts++
		}
		fmt.Printf("%v ticks to hit ground\n", ts)
		err := c.ToPPM("bh.ppm")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("done!")
	*/
}
