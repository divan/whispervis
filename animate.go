package main

import "github.com/gopherjs/gopherjs/js"

func (p *Page) animate() {
	if p.renderer == nil {
		return
	}

	p.controls.Update()

	js.Global.Call("requestAnimationFrame", p.animate)

	if p.autoRotate {
		pos := p.group.Object.Get("rotation")
		pos.Set("y", pos.Get("y").Float()+float64(0.01))
	}

	p.renderer.Render(p.scene, p.camera)
}

func (p *Page) ToggleAutoRotation() {
	p.autoRotate = !p.autoRotate
}
