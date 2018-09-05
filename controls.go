package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lngramos/three"
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
