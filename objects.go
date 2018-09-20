package main

import (
	"fmt"

	"github.com/divan/graphx/graph"
	"github.com/divan/graphx/layout"
	"github.com/divan/three"
)

// CreateObjects creates WebGL primitives from layout/graphGroup data.
// TODO(divan): change positions and links types to something more clear and readable
func (w *WebGLScene) CreateObjects(positions map[string]*layout.Object, links []*graph.Link) {
	w.graphGroup = three.NewGroup()
	w.scene.Add(w.graphGroup)

	w.nodesGroup = three.NewGroup()
	w.graphGroup.Add(w.nodesGroup)

	w.edgesGroup = three.NewGroup()
	w.graphGroup.Add(w.edgesGroup)

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
		w.nodesGroup.Add(mesh)
		w.nodes = append(w.nodes, mesh)
	}
	obj := w.nodesGroup.GetObjectById(100)
	fmt.Println("Moving mesh", obj)
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
		w.edgesGroup.Add(line)
		w.lines = append(w.lines, line)
	}
}

// RemoveObjects removes WebGL primitives, cleaning up scene.
func (w *WebGLScene) RemoveObjects() {
	if w.nodesGroup != nil {
		for _, child := range w.nodesGroup.Children {
			w.nodesGroup.Remove(child)
		}
	}
	if w.edgesGroup != nil {
		for _, child := range w.edgesGroup.Children {
			w.edgesGroup.Remove(child)
		}
	}

	w.nodes, w.lines = nil, nil
	w.graphGroup, w.nodesGroup, w.edgesGroup = nil, nil, nil
}
