package main

import (
	"fmt"
	"time"

	"github.com/status-im/simulation/propagation"
)

const AnimationSlowdown = 1

// AnimatePropagation visualizes propagation of message based on plog.
func (w *WebGLScene) AnimatePropagation(plog *propagation.Log) {
	fmt.Println("Animating plog")
	for i, ts := range plog.Timestamps {
		duration := time.Duration(time.Duration(ts) * time.Millisecond)
		duration = duration * AnimationSlowdown

		nodes := plog.Nodes[i]
		edges := plog.Indices[i]
		fn := func() {
			// blink nodes for this timestamp
			for _, idx := range nodes {
				w.BlinkNode(idx)
			}
			// blink links for this timestamp
			for _, idx := range edges {
				w.BlinkEdge(idx)
			}
		}
		time.AfterFunc(duration, fn)
	}
}
