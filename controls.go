package main

import (
	"github.com/divan/three"
	"github.com/gopherjs/gopherjs/js"
)

type TrackBallControl struct {
	*js.Object
}

func NewTrackBallControl(camera three.PerspectiveCamera, renderer *three.WebGLRenderer) TrackBallControl {
	dom := renderer.Get("domElement")
	return TrackBallControl{
		Object: js.Global.Get("THREE").Get("TrackballControls").New(camera, dom),
	}
}

func (t TrackBallControl) Update() {
	t.Call("update")
}
