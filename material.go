package main

import "github.com/lngramos/three"

// NewNodeMaterial creates a new default material for the graph node lines.
func NewNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(0, 255, 0)
	return three.NewMeshPhongMaterial(params)
}

// NewEdgeMaterial creates a new default material for the graph edge lines.
func NewEdgeMatherial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(200, 200, 255)
	return three.NewLineBasicMaterial(params)
}
