package main

import (
	"time"

	"github.com/divan/three"
	"github.com/status-im/simulation/propagation"
)

var (
	BlinkedEdgeMaterials = NewBlinkedEdgeMaterials()
	BlinkedNodeMaterials = NewBlinkedNodeMaterials()
)

// AnimatePropagation visualizes propagation of message based on plog.
func (w *WebGLScene) AnimatePropagation(plog *propagation.Log) {
	w.rt.EnableRendering()

	maxTs := plog.Timestamps[len(plog.Timestamps)-1]
	for i, ts := range plog.Timestamps {
		duration := time.Duration(time.Duration(ts) * time.Millisecond)

		percentage := (ts * 100) / maxTs // % of plog
		if percentage > 99 {
			percentage = 99
		}

		nodes := plog.Nodes[i]
		edges := plog.Links[i]
		delay := time.Duration(w.blink) * time.Millisecond
		fn := func() {
			// blink nodes for this timestamp
			for _, idx := range nodes {
				w.BlinkNode(idx, percentage)
				go time.AfterFunc(delay, func() { w.UnblinkNode(idx) })
			}
			// blink links for this timestamp
			for _, idx := range edges {
				w.BlinkEdge(idx, percentage)
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
			w.BlinkNode(i, 99)
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
			w.BlinkEdge(i, 99)
		} else {
			w.UnblinkEdge(i)
		}
	}
}

// BlinkNode animates a single node blinking. Node specified by its idx.
// TODO(divan): consider renaming it to Highlight or something.
func (w *WebGLScene) BlinkNode(id, percentage int) {
	node := w.nodes[id]
	node.Set("material", BlinkedNodeMaterials[percentage/10]) // choose material depending on percentage of propagation
}

func (w *WebGLScene) UnblinkNode(id int) {
	node := w.nodes[id]
	node.Object.Set("material", DefaultNodeMaterial)
}

// BlinkEdge animates a single edge blinking. Edge specified by its idx.
func (w *WebGLScene) BlinkEdge(id, percentage int) {
	edge := w.lines[id]
	edge.Set("material", BlinkedEdgeMaterials[percentage/10]) // choose material depending on percentage of propagation

	delay := time.Duration(w.blink) * time.Millisecond
	go time.AfterFunc(delay, func() { w.UnblinkEdge(id) })
}

func (w *WebGLScene) UnblinkEdge(id int) {
	node := w.lines[id]
	node.Object.Set("material", DefaultEdgeMaterial)
}

// NewBlinkedEdgeMaterials creates a new default material for the graph blinked edge lines.
func NewBlinkedEdgeMaterials() []three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(255, 0, 0)
	params.Transparent = true
	ret := make([]three.Material, 0, 10)
	for i := 0; i < 10; i++ {
		params.Opacity = float64(1 - (float64(i) * 0.05)) // 1, 0.95, 0.90, 0.85...
		ret = append(ret, three.NewLineBasicMaterial(params))
	}
	return ret
}

// NewBlinkedNodeMaterials creates a new default material for the graph blinked node.
func NewBlinkedNodeMaterials() []three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(255, 0, 0) // red
	params.Transparent = true
	ret := make([]three.Material, 0, 10)
	for i := 0; i < 10; i++ {
		params.Opacity = float64(1 - (float64(i) * 0.05)) // 1, 0.95, 0.90, 0.85...
		ret = append(ret, three.NewMeshPhongMaterial(params))
	}
	return ret
}
