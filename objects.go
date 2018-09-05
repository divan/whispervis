package main

import (
	"github.com/lngramos/three"
)

func (p *Page) CreateObjects() {
	p.group = three.NewGroup()
	p.scene.Add(p.group)

	geometry := NewEthereumGeometry(2)

	materialParams := three.NewMaterialParameters()
	materialParams.Color = three.NewColor(0, 255, 0)
	material := three.NewMeshPhongMaterial(materialParams)

	for _, node := range p.layout.Positions() {
		mesh := three.NewMesh(geometry, material)
		mesh.Position.Set(node.X, node.Y, node.Z)
		p.group.Add(mesh)
	}

	// Lines
	lineMaterialParams := three.NewMaterialParameters()
	lineMaterialParams.Color = three.NewColor(200, 200, 255)
	lineMaterial := three.NewLineBasicMaterial(lineMaterialParams)

	for _, link := range p.layout.Links() {
		from := link.From()
		to := link.To()
		start := p.layout.Positions()[from]
		end := p.layout.Positions()[to]

		var geom = three.NewBasicGeometry(three.BasicGeometryParams{})
		geom.AddVertice(start.X, start.Y, start.Z)
		geom.AddVertice(end.X, end.Y, end.Z)

		line := three.NewLine(geom, lineMaterial)
		p.group.Add(line)
	}
}
