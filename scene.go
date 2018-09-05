package main

import "github.com/lngramos/three"

// InitScene inits a new scene, sets up camera, lights and all that.
func (p *Page) InitScene(width, height float64) {
	p.camera = three.NewPerspectiveCamera(70, width/height, 1, 1000)
	p.camera.Position.Set(0, 0, 400)

	p.scene = three.NewScene()

	p.InitLights()
	p.InitControls()
}

// InitLights init lights for the scene.
func (p *Page) InitLights() {
	ambLight := three.NewAmbientLight(three.NewColor(187, 187, 187), 0.5)
	p.scene.Add(ambLight)

	light := three.NewDirectionalLight(three.NewColor(255, 255, 255), 0.3)
	//light.Position.Set(256, 256, 256).Normalize()
	p.scene.Add(light)
}

// InitControls init controls for the scene.
func (p *Page) InitControls() {
	p.controls = NewTrackBallControl(p.camera, p.renderer)
}
