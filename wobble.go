package main

import (
	"math/rand"
	"time"

	"github.com/divan/graphx/layout"

	"github.com/gopherjs/gopherjs/js"
)

const (
	WobblingPeriod = 2 * time.Second
	WobblingCoeff  = 0.2 // the larger, the bigger distance of wobbling
)

// Wobbling implements wobbling effect for graph nodes.
//
// It updates positions and calls scene methods to update objects accordingly.
type Wobbling struct {
	lastChange time.Time // timestamp of last direction change
	forward    bool

	directions map[string]*Direction
	positions  map[string]*layout.Object
}

// NewWobbling creates a new wobbling effect for graph.
func NewWobbling(positions map[string]*layout.Object) *Wobbling {
	return &Wobbling{
		positions:  positions,
		directions: make(map[string]*Direction),
	}
}

// Animate will run on every requestAnimationStep call.
func (w *Wobbling) Animate() {
	if w == nil {
		return
	}

	if time.Now().After(w.lastChange) {
		w.changeDirection()
		w.lastChange = time.Now().Add(WobblingPeriod)
	}

	w.wobbleNodes()
}

func (w *Wobbling) wobbleNodes() {
	for id, pos := range w.positions {
		d := w.directions[id]
		pos.X += d.X
		pos.Y += d.Y
		pos.Z += d.Z
	}
}

func (w *Wobbling) changeDirection() {
	rand.Seed(time.Now().UnixNano())
	w.forward = !w.forward
	for id := range w.positions {
		var direction *Direction
		if w.forward {
			direction = NewRandomDirection()
		} else {
			// return to original position
			orig := w.directions[id]
			direction = orig.Reverse()
		}
		w.directions[id] = direction
	}
}

// moveRandom randomly moves the three.js object to given distance.
// TODO(divan): use Mesh's Position?
func moveRandom(obj *js.Object, d *Direction) {
	pos := obj.Get("position")
	pos.Set("x", pos.Get("x").Float()+d.X)
	pos.Set("y", pos.Get("y").Float()+d.Y)
	pos.Set("z", pos.Get("z").Float()+d.Z)
}

type Direction struct {
	X, Y, Z float64
}

func NewRandomDirection() *Direction {
	return &Direction{
		X: (rand.Float64() - 0.5) * WobblingCoeff,
		Y: (rand.Float64() - 0.5) * WobblingCoeff,
		Z: (rand.Float64() - 0.5) * WobblingCoeff,
	}
}

func (d *Direction) Reverse() *Direction {
	return &Direction{
		X: -d.X,
		Y: -d.Y,
		Z: -d.Z,
	}
}
