package main

import (
	"time"

	"github.com/divan/three"
	"github.com/status-im/simulation/propagation"
)

var (
	BlinkedEdgeMaterial = NewBlinkedEdgeMaterial()
	BlinkedNodeMaterial = NewBlinkedNodeMaterial()
)

// AnimatePropagation visualizes propagation of message based on plog.
func (w *WebGLScene) AnimatePropagation(plog *propagation.Log) {
	w.rt.EnableRendering()

	for i, ts := range plog.Timestamps {
		duration := time.Duration(time.Duration(ts) * time.Millisecond)

		nodes := plog.Nodes[i]
		edges := plog.Links[i]
		delay := time.Duration(w.blink) * time.Millisecond
		fn := func() {
			// blink nodes for this timestamp
			for _, idx := range nodes {
				w.BlinkNode(idx)
				go time.AfterFunc(delay, func() { w.UnblinkNode(idx) })
			}
			// blink links for this timestamp
			for _, idx := range edges {
				w.BlinkEdge(idx)
				go time.AfterFunc(delay, func() { w.UnblinkEdge(idx) })
			}
		}
		go time.AfterFunc(duration, fn)

		w.rt.EnableRendering() // prevent thorttler from enabling during long animations
	}
}

// AnimateOneStep blinks permantently nodes and edges for the given step of plog.
// TODO(divan): "Animate-" is probably not the best name here, come up with something
// better (this function doesn't *animate* anything so far)
func (w *WebGLScene) AnimateOneStep(plog *propagation.Log, step int) {
	nodes := plog.Nodes[step]
	edges := plog.Links[step]

	nodesToBlink := make(map[int]struct{})
	for _, idx := range nodes {
		nodesToBlink[idx] = struct{}{}
	}
	// blink nodes for this timestamp
	for i, _ := range w.nodes {
		if _, ok := nodesToBlink[i]; ok {
			w.BlinkNode(i)
		} else {
			w.UnblinkNode(i)
		}
	}

	edgesToBlink := make(map[int]struct{})
	for _, idx := range edges {
		edgesToBlink[idx] = struct{}{}
	}
	for i, _ := range w.lines {
		if _, ok := edgesToBlink[i]; ok {
			w.BlinkEdge(i)
		} else {
			w.UnblinkEdge(i)
		}
	}
}

// BlinkNode animates a single node blinking. Node specified by its idx.
// TODO(divan): consider renaming it to Highlight or something.
func (w *WebGLScene) BlinkNode(id int) {
	node := w.nodes[id]
	node.Set("material", BlinkedNodeMaterial)
}

func (w *WebGLScene) UnblinkNode(id int) {
	node := w.nodes[id]
	node.Object.Set("material", DefaultNodeMaterial)
}

// BlinkEdge animates a single edge blinking. Edge specified by its idx.
func (w *WebGLScene) BlinkEdge(id int) {
	edge := w.lines[id]
	edge.Set("material", BlinkedEdgeMaterial)
}

func (w *WebGLScene) UnblinkEdge(id int) {
	node := w.lines[id]
	node.Object.Set("material", DefaultEdgeMaterial)
}

// NewBlinkedEdgeMaterial creates a new default material for the graph blinked edge lines.
func NewBlinkedEdgeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(255, 0, 0)
	params.Transparent = false
	return three.NewLineBasicMaterial(params)
}

// NewBlinkedNodeMaterial creates a new default material for the graph blinked node.
func NewBlinkedNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(255, 0, 0) // red
	params.Transparent = false
	return three.NewMeshPhongMaterial(params)
}
