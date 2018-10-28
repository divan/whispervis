package main

import (
	"fmt"

	"github.com/divan/graphx/layout"
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
	"github.com/status-im/whispervis/storage"
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

	rt    *RenderThrottler // used as a helper to reduce rendering calls when animation is not needed (experimental)
	fps   int              // frames per second
	blink int              // time in ms for blinking nodes/edges

	initFn func() // function to run on initialization
}

// NewWebGLScene inits and returns new WebGL scene and canvas.
func NewWebGLScene(initFn func()) *WebGLScene {
	w := &WebGLScene{
		rt:     NewRenderThrottler(),
		initFn: initFn,
		fps:    storage.FPS(),
		blink:  storage.BlinkTime(),
	}
	w.WebGLRenderer = vthree.NewWebGLRenderer(vthree.WebGLOptions{
		Init:      w.init,
		Shutdown:  w.shutdown,
		Antialias: true,
	})
	return w
}

func (w *WebGLScene) init(renderer *three.WebGLRenderer) {
	fmt.Println("WebGL init")
	windowWidth := js.Global.Get("innerWidth").Float() - 300 // TODO(divan): sync this with page layout
	windowHeight := js.Global.Get("innerHeight").Float() - 20

	w.renderer = renderer
	w.renderer.SetSize(windowWidth, windowHeight, true)

	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()
	w.renderer.SetPixelRatio(devicePixelRatio)

	w.InitScene(windowWidth, windowHeight)

	if w.initFn != nil {
		w.initFn()
	}
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
}

// InitScene inits a new scene, sets up camera, lights and all that.
func (w *WebGLScene) InitScene(width, height float64) {
	w.camera = three.NewPerspectiveCamera(70, width/height, 1, 50000)
	w.camera.Position.Set(0, 0, 100)

	w.scene = three.NewScene()
	w.scene.Background = three.NewColorRGB(0, 0, 17)

	w.InitLights()
	w.InitControls()
	w.Reset()
}

// InitLights init lights for the scene.
func (w *WebGLScene) InitLights() {
	ambLight := three.NewAmbientLight(three.NewColorHex(0xbbbbbb), 1)
	ambLight.MatrixAutoUpdate = false
	w.scene.Add(ambLight)

	light := three.NewDirectionalLight(three.NewColor("white"), 1)
	light.MatrixAutoUpdate = false
	w.scene.Add(light)
}

// InitControls init controls for the scene.
func (w *WebGLScene) InitControls() {
	w.controls = NewTrackBallControl(w.camera, w.renderer)
}

// WindowSize returns current window's width and height.
func WindowSize() (int, int) {
	w := js.Global.Get("innerWidth").Float()
	h := js.Global.Get("innerHeight").Float()
	return int(w), int(h)
}

// PageViewSize returns current page aread width and height (adjusted for sidebar and margins).
func PageViewSize() (int, int) {
	w, h := WindowSize()
	w = w - 300 // TODO: remove magic (300 is a width of sidebar)
	h = h - 20  // some top margin
	return w, h
}

// ToggleRenderThrottler switches render throttling option. Implements SceneConfigurator.
func (w *WebGLScene) ToggleRenderThrottler() {
	w.rt.Toggle()
}

// ChangeFPS changes rendering FPS. Implements SceneConfigurator.
func (w *WebGLScene) ChangeFPS(value int) {
	fmt.Println("Changing FPS to", value)
	w.fps = value
	storage.SetFPS(value)
}

// ChangeBlinkTime changes propagation animation blink time. Implements SceneConfigurator.
func (w *WebGLScene) ChangeBlinkTime(value int) {
	w.blink = value
	storage.SetBlinkTime(value)
}
