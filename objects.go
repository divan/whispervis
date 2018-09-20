package main

import (
	"github.com/divan/graphx/graph"
	"github.com/divan/graphx/layout"
	"github.com/divan/three"
)

// CreateObjects creates WebGL primitives from layout/graph data.
// TODO(divan): change positions and links types to something more clear and readable
func (w *WebGLScene) CreateObjects(positions map[string]*layout.Object, links []*graph.Link) {
	w.RemoveObjects()

	w.graph = three.NewGroup()
	w.scene.Add(w.graph)

	w.nodes = three.NewGroup()
	w.graph.Add(w.nodes)

	w.edges = three.NewGroup()
	w.graph.Add(w.edges)

	w.createNodes(positions)
	w.createEdges(positions, links)
}

func (w *WebGLScene) createNodes(positions map[string]*layout.Object) {
	scale := 2.0
	geometry := NewEthereumGeometry(scale)
	material := NewNodeMaterial()
	for _, node := range positions {
		mesh := three.NewMesh(geometry, material)
		mesh.Position.Set(node.X, node.Y, node.Z)
		w.nodes.Add(mesh)
	}
}

func (w *WebGLScene) createEdges(positions map[string]*layout.Object, links []*graph.Link) {
	material := NewEdgeMatherial()
	for _, link := range links {
		from := link.From()
		to := link.To()
		start := positions[from]
		end := positions[to]

		var geom = three.NewBasicGeometry(three.BasicGeometryParams{})
		geom.AddVertice(start.X, start.Y, start.Z)
		geom.AddVertice(end.X, end.Y, end.Z)

		line := three.NewLine(geom, material)
		w.edges.Add(line)
	}
}

// RemoveObjects removes WebGL primitives, cleaning up scene.
func (w *WebGLScene) RemoveObjects() {
	// TODO(divan): remove objects from scene via vthree
	w.graph, w.nodes, w.edges = nil, nil, nil
}
