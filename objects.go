package main

import (
	"github.com/divan/graphx/graph"
	"github.com/divan/graphx/layout"
	"github.com/divan/three"
)

// Mesh represents three.Mesh, and holds some additional metadata.
type Mesh struct {
	ID string
	*three.Mesh
}

// Line represents three.Line, and holds some additional metadata.
type Line struct {
	From string
	To   string
	*three.Line
}

// CreateObjects creates WebGL primitives from layout/graphGroup data.
// TODO(divan): change positions and links types to something more clear and readable
func (w *WebGLScene) CreateObjects(positions map[string]*layout.Object, links []*graph.Link) {
	w.positions = positions

	w.graphGroup = three.NewGroup()
	w.graphGroup.MatrixAutoUpdate = false
	w.scene.Add(w.graphGroup)

	w.nodesGroup = three.NewGroup()
	w.nodesGroup.MatrixAutoUpdate = false
	w.graphGroup.Add(w.nodesGroup)

	w.edgesGroup = three.NewGroup()
	w.edgesGroup.MatrixAutoUpdate = false
	w.graphGroup.Add(w.edgesGroup)

	w.createNodes()
	w.createEdges(links)

	// once we know nodes, we can prepare fancy stuff
	w.wobbling = NewWobbling(w.positions)
}

func (w *WebGLScene) createNodes() {
	scale := 2.0
	geometry := NewEthereumGeometry(scale)
	material := NewNodeMaterial()
	for id, node := range w.positions {
		mesh := &Mesh{
			ID:   id,
			Mesh: three.NewMesh(geometry, material),
		}
		mesh.Position.Set(node.X, node.Y, node.Z)
		mesh.MatrixAutoUpdate = false
		mesh.UpdateMatrix()
		w.nodesGroup.Add(mesh.Mesh)
		w.nodes = append(w.nodes, mesh)
	}
}

func (w *WebGLScene) createEdges(links []*graph.Link) {
	material := NewEdgeMaterial()
	for _, link := range links {
		from := link.From()
		to := link.To()
		start := w.positions[from]
		end := w.positions[to]

		var geom = three.NewBufferGeometry()
		var positions = make([]float32, 2*3) // 2 positions (start and end), 3 coords per each
		var attr = three.NewBufferAttribute(positions, 3)

		geom.AddAttribute("position", attr)

		attr.SetXYZ(0, start.X, start.Y, start.Z)
		attr.SetXYZ(1, end.X, end.Y, end.Z)
		attr.NeedsUpdate = true

		line := &Line{
			From: from,
			To:   to,
			Line: three.NewLine(geom, material),
		}
		line.MatrixAutoUpdate = false
		w.edgesGroup.Add(line.Line)
		w.lines = append(w.lines, line)
	}
}

// updatePositions sets meshes/lines positions from positions (probably
// recalculated somewhere else)
func (w *WebGLScene) updatePositions() {
	for _, node := range w.nodes {
		pos := w.positions[node.ID]
		node.Position.Set(pos.X, pos.Y, pos.Z)
		node.UpdateMatrix()
	}
	for i := range w.lines {
		start := w.positions[w.lines[i].From]
		end := w.positions[w.lines[i].To]

		attr := w.lines[i].Geometry.GetAttribute("position")
		attr.SetXYZ(0, start.X, start.Y, start.Z)
		attr.SetXYZ(1, end.X, end.Y, end.Z)
		attr.NeedsUpdate = true
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
	w.positions = nil
	w.graphGroup, w.nodesGroup, w.edgesGroup = nil, nil, nil
}
