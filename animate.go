package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func (w *WebGLScene) animate() {
	if w.renderer == nil {
		return
	}

	w.controls.Update()

	js.Global.Call("requestAnimationFrame", w.animate)

	if w.autoRotate {
		pos := w.graph.Object.Get("rotation")
		pos.Set("y", pos.Get("y").Float()+float64(0.001))
	}

	w.renderer.Render(w.scene, w.camera)
}

// ToggleAutoRotation switches auto rotation option.
func (w *WebGLScene) ToggleAutoRotation() {
	w.autoRotate = !w.autoRotate
}
