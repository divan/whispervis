package main

import (
	"fmt"
	"time"

	"github.com/divan/three"
	"github.com/status-im/simulation/propagation"
)

const (
	BlinkDecay        = 200 * time.Millisecond // time for highlighted node/link to be active
	AnimationSlowdown = 1                      // slowdown factor for propagation animation
)

var (
	BlinkedEdgeMaterials = NewBlinkedEdgeMaterials()
	BlinkedNodeMaterials = NewBlinkedNodeMaterials()
)

// AnimatePropagation visualizes propagation of message based on plog.
func (w *WebGLScene) AnimatePropagation(plog *propagation.Log) {
	fmt.Println("Animating plog")
	w.rt.Disable()

	maxTs := plog.Timestamps[len(plog.Timestamps)-1]
	for i, ts := range plog.Timestamps {
		duration := time.Duration(time.Duration(ts) * time.Millisecond)
		duration = duration * AnimationSlowdown

		percentage := (ts * 100) / maxTs // % of plog
		if percentage > 99 {
			percentage = 99
		}

		nodes := plog.Nodes[i]
		edges := plog.Links[i]
		fn := func() {
			// blink nodes for this timestamp
			for _, idx := range nodes {
				w.BlinkNode(idx, percentage)
			}
			// blink links for this timestamp
			for _, idx := range edges {
				w.BlinkEdge(idx, percentage)
			}
		}
		go time.AfterFunc(duration, fn)

		w.rt.Disable() // prevent thorttler from enabling during long animations
	}
}

// BlinkNode animates a single node blinking. Node specified by its idx.
func (w *WebGLScene) BlinkNode(id, percentage int) {
	node := w.nodes[id]
	node.Set("material", BlinkedNodeMaterials[percentage/10]) // choose material depending on percentage of propagation
	restore := func() { node.Object.Set("material", DefaultNodeMaterial) }
	go time.AfterFunc(BlinkDecay, restore)

}

// BlinkEdge animates a single edge blinking. Edge specified by its idx.
func (w *WebGLScene) BlinkEdge(id, percentage int) {
	edge := w.lines[id]
	edge.Set("material", BlinkedEdgeMaterials[percentage/10]) // choose material depending on percentage of propagation
	restore := func() { edge.Object.Set("material", DefaultEdgeMaterial) }
	go time.AfterFunc(BlinkDecay, restore)
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
