package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/status-im/simulation/propagation"
)

const (
	// TODO(divan): move this as variables to the frontend
	BlinkDecay        = 100 * time.Millisecond // time for highlighted node/link to be active
	AnimationSlowdown = 1                      // slowdown factor for propagation animation
)

func (w *WebGLScene) animate() {
	if w.renderer == nil {
		return
	}

	w.controls.Update()

	js.Global.Call("requestAnimationFrame", w.animate)

	if w.autoRotate {
		pos := w.graphGroup.Object.Get("rotation")
		pos.Set("y", pos.Get("y").Float()+float64(0.001))
	}

	if w.wobble {
		w.wobbling.Animate()
		w.updatePositions()
	}

	w.renderer.Render(w.scene, w.camera)
}

// ToggleAutoRotation switches auto rotation option.
func (w *WebGLScene) ToggleAutoRotation() {
	w.autoRotate = !w.autoRotate
}

// ToggleWobbling switches wobbling option.
func (w *WebGLScene) ToggleWobbling() {
	w.wobble = !w.wobble
}

// BlinkNode animates a single node blinking. Node specified by its idx.
func (w *WebGLScene) BlinkNode(id int) {
	node := w.nodes[id]
	node.Set("material", BlinkedNodeMaterial)
	restore := func() { node.Object.Set("material", DefaultNodeMaterial) }
	time.AfterFunc(BlinkDecay, restore)

}

// BlinkEdge animates a single edge blinking. Edge specified by its idx.
func (w *WebGLScene) BlinkEdge(id int) {
	edge := w.lines[id]
	edge.Set("material", BlinkedEdgeMaterial)
	restore := func() { edge.Object.Set("material", DefaultEdgeMaterial) }
	time.AfterFunc(BlinkDecay, restore)
}

// AnimatePropagation visualizes propagation of message based on plog.
func (w *WebGLScene) AnimatePropagation(plog *propagation.Log) {
	fmt.Println("Animating plog")
	for i, ts := range plog.Timestamps {
		duration := time.Duration(time.Duration(ts) * time.Millisecond)
		duration = duration * AnimationSlowdown

		nodes := plog.Nodes[i]
		edges := plog.Indices[i]
		fn := func() {
			// blink nodes for this timestamp
			for _, idx := range nodes {
				w.BlinkNode(idx)
			}
			// blink links for this timestamp
			for _, idx := range edges {
				w.BlinkEdge(idx)
			}
		}
		time.AfterFunc(duration, fn)
	}
}
