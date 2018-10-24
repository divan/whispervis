package main

import "github.com/divan/three"

var (
	DefaultNodeMaterial     = NewNodeMaterial()
	BlinkedNodeMaterial     = NewBlinkedNodeMaterial()
	TransparentNodeMaterial = NewTransparentNodeMaterial()
	DefaultEdgeMaterial     = NewEdgeMaterial()
	BlinkedEdgeMaterial     = NewBlinkedEdgeMaterial()
)

const (
	DefaultTransparency = 0.9
)

// NewNodeMaterial creates a new default material for the graph node.
func NewNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(0, 255, 0)
	params.Transparent = false
	params.Opacity = DefaultTransparency
	return three.NewMeshPhongMaterial(params)
}

// NewBlinkedNodeMaterial creates a new default material for the graph blinked node.
func NewBlinkedNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(255, 0, 0) // red
	params.Transparent = true
	params.Opacity = DefaultTransparency
	return three.NewMeshPhongMaterial(params)
}

// NewTransparentNodeMaterial creates a new transparent material for the graph normal node.
func NewTransparentNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(0, 255, 0)
	params.Transparent = true
	params.Opacity = 0.5
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

// NewBlinkedEdgeMaterial creates a new default material for the graph blinked edge lines.
func NewBlinkedEdgeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColorRGB(255, 0, 0)
	params.Transparent = true
	params.Opacity = DefaultTransparency
	return three.NewLineBasicMaterial(params)
}
