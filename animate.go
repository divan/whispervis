package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
)

const BlinkDecay = 100 * time.Millisecond

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
