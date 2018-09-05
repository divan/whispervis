package main

import (
	"github.com/lngramos/three"
)

func (p *Page) CreateObjects() {
	p.graph = three.NewGroup()
	p.scene.Add(p.graph)

	p.nodes = three.NewGroup()
	p.graph.Add(p.nodes)

	p.edges = three.NewGroup()
	p.graph.Add(p.edges)

	p.createNodes()
	p.createEdges()
}

func (p *Page) createNodes() {
	scale := 2.0
	geometry := NewEthereumGeometry(scale)
	material := NewNodeMaterial()
	for _, node := range p.layout.Positions() {
		mesh := three.NewMesh(geometry, material)
		mesh.Position.Set(node.X, node.Y, node.Z)
		p.nodes.Add(mesh)
	}
}

func (p *Page) createEdges() {
	material := NewEdgeMatherial()
	for _, link := range p.layout.Links() {
		from := link.From()
		to := link.To()
		start := p.layout.Positions()[from]
		end := p.layout.Positions()[to]

		var geom = three.NewBasicGeometry(three.BasicGeometryParams{})
		geom.AddVertice(start.X, start.Y, start.Z)
		geom.AddVertice(end.X, end.Y, end.Z)

		line := three.NewLine(geom, material)
		p.edges.Add(line)
	}
}
