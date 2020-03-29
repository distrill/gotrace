package main

import (
	"fmt"
	"github.com/distrill/gotrace/tuples"
)

/*
	end of chapter messing around:
*/
// type point Tuple
// type vector Tuple

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
	p := projectile{
		tuples.NewPoint(0, 1, 0),
		tuples.NewVector(1, 1, 0).Norm(),
	}
	e := environment{
		tuples.NewVector(0, -0.1, 0),
		tuples.NewVector(-0.01, 0, 0),
	}
	ts := 0
	for {
		if p.position.Y <= 0 {
			break
		}
		fmt.Printf("%+v\n", p)
		p = tick(e, p)
		ts++
	}
	fmt.Printf("%v ticks to hit ground\n", ts)
}
