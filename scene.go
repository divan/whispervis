package main

import (
	"fmt"

	"github.com/divan/graphx/layout"
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/status-im/whispervis/vthree"
)

// WebGLScene represents WebGL part of app.
type WebGLScene struct {
	*vthree.WebGLRenderer

	scene      *three.Scene
	camera     three.PerspectiveCamera
	renderer   *three.WebGLRenderer
	graphGroup *three.Group
	nodesGroup *three.Group
	edgesGroup *three.Group
	controls   TrackBallControl

	autoRotate bool
	wobble     bool
	wobbling   *Wobbling

	positions map[string]*layout.Object

	// these slices exist here because we have no good way to access three.Group children for now
	// TODO(divan): as soon as three.js wrappers allow us to access children, get rid of it here
	nodes []*Mesh
	lines []*Line
}

// NewWebGLScene inits and returns new WebGL scene and canvas.
func NewWebGLScene() *WebGLScene {
	w := &WebGLScene{}
	w.WebGLRenderer = vthree.NewWebGLRenderer(vthree.WebGLOptions{
		Init:     w.init,
		Shutdown: w.shutdown,
	})
	return w
}

func (w *WebGLScene) init(renderer *three.WebGLRenderer) {
	fmt.Println("WebGL init")
	windowWidth := js.Global.Get("innerWidth").Float()*80/100 - 20
	windowHeight := js.Global.Get("innerHeight").Float() - 20

	w.renderer = renderer
	w.renderer.SetSize(windowWidth, windowHeight, true)

	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()
	w.renderer.SetPixelRatio(devicePixelRatio)

	w.InitScene(windowWidth, windowHeight)

	w.animate()
}

func (w *WebGLScene) shutdown(renderer *three.WebGLRenderer) {
	fmt.Println("WebGL shutdown")
	w.scene = nil
	w.camera = three.PerspectiveCamera{}
	w.renderer = nil
	w.RemoveObjects()
}

// Reset resets state of WebGLScene.
func (w *WebGLScene) Reset() {
	fmt.Println("Resetting WebGL")
	w.RemoveObjects()
	zeroCamera := three.PerspectiveCamera{}
	if w.camera != zeroCamera {
		w.camera.Position.Set(0, 0, 400)
	}
}

// InitScene inits a new scene, sets up camera, lights and all that.
func (w *WebGLScene) InitScene(width, height float64) {
	w.camera = three.NewPerspectiveCamera(70, width/height, 1, 1000)

	w.scene = three.NewScene()

	w.InitLights()
	w.InitControls()
	w.Reset()
}

// InitLights init lights for the scene.
func (w *WebGLScene) InitLights() {
	ambLight := three.NewAmbientLight(three.NewColor(187, 187, 187), 0.5)
	w.scene.Add(ambLight)

	light := three.NewDirectionalLight(three.NewColor(255, 255, 255), 0.3)
	//light.Position.Set(256, 256, 256).Normalize()
	w.scene.Add(light)
}

// InitControls init controls for the scene.
func (w *WebGLScene) InitControls() {
	w.controls = NewTrackBallControl(w.camera, w.renderer)
}
