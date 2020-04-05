package main

import (
	"fmt"
	"math"

	"github.com/distrill/gotrace/canvas"
	"github.com/distrill/gotrace/colors"
	"github.com/distrill/gotrace/matrix"
	"github.com/distrill/gotrace/tuples"
)

type projectile struct {
	position tuples.Tuple
	velocity tuples.Tuple
}

type environment struct {
	gravity tuples.Tuple
	wind    tuples.Tuple
}

func tick(e environment, p projectile) projectile {
	pos := p.position.Add(tuples.Tuple(p.velocity))
	vel := p.velocity.Add(tuples.Tuple(e.gravity))
	return projectile{pos, vel}
}

func main() {
	w := 100
	h := 100
	center := 50.0
	radius := center * .8
	c := canvas.NewCanvas(w, h)

	p := tuples.NewPoint(0, -radius, 0)

	for i := 0; i <= 12; i++ {
		p = matrix.NewTransform(p).
			RotateZ(math.Pi / 6).
			Value()

		c.WritePixel(
			int(center-math.Round(p.X)),
			int(center-math.Round(p.Y)),
			colors.Red,
		)
	}

	err := c.ToPPM("bh.ppm")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("done!")

	/*
		start := tuples.NewPoint(1, 1.5, 0)
		velocity := tuples.NewVector(1, 2, 0).Norm().Mul(11.25)
		p := projectile{start, velocity}

		gravity := tuples.NewVector(0, -0.11, 0)
		wind := tuples.NewVector(-0.01, 0, 0)
		e := environment{gravity, wind}

		w := 900
		h := 550
		c := canvas.NewCanvas(w, h)

		ts := 0
		for {
			if p.position.Y <= 0 {
				break
			}
			c.WritePixel(
				int(math.Round(p.position.X)),
				int(math.Round(p.position.Y)),
				colors.Red,
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
