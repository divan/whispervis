package main

import "github.com/divan/three"

var (
	DefaultNodeMaterial     = NewNodeMaterial()
	TransparentNodeMaterial = NewTransparentNodeMaterial()
	DefaultEdgeMaterial     = NewEdgeMaterial()
)

const DefaultTransparency = 0.9

// NewNodeMaterial creates a new default material for the graph node.
func NewNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(0, 255, 0)
	params.Transparent = false
	params.Opacity = DefaultTransparency
	return three.NewMeshPhongMaterial(params)
}

// NewTransparentNodeMaterial creates a new transparent material for the graph normal node.
func NewTransparentNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(0, 255, 0)
	params.Transparent = true
	params.Opacity = 0.7
	return three.NewMeshPhongMaterial(params)
}

// NewEdgeMaterial creates a new default material for the graph edge lines.
func NewEdgeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorHex(0xf0f0f0)
	params.Transparent = true
	params.Opacity = 0.7
	return three.NewLineBasicMaterial(params)
}
