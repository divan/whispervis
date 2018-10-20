package main

import "github.com/divan/three"

var (
	DefaultNodeMaterial = NewNodeMaterial()
	BlinkedNodeMaterial = NewBlinkedNodeMaterial()
	DefaultEdgeMaterial = NewEdgeMaterial()
	BlinkedEdgeMaterial = NewBlinkedEdgeMaterial()
)

const (
	DefaultTransparency = 0.9
)

// NewNodeMaterial creates a new default material for the graph node.
func NewNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(0, 255, 0)
	params.Transparent = true
	params.Opacity = DefaultTransparency
	return three.NewMeshPhongMaterial(params)
}

// NewBlinkedNodeMaterial creates a new default material for the graph blinked node.
func NewBlinkedNodeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(255, 0, 0) // red
	params.Transparent = true
	params.Opacity = DefaultTransparency
	return three.NewMeshPhongMaterial(params)
}

// NewEdgeMaterial creates a new default material for the graph edge lines.
func NewEdgeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(200, 200, 255)
	params.Transparent = true
	params.Opacity = DefaultTransparency
	return three.NewLineBasicMaterial(params)
}

// NewBlinkedEdgeMaterial creates a new default material for the graph blinked edge lines.
func NewBlinkedEdgeMaterial() three.Material {
	params := three.NewMaterialParameters()
	params.Color = three.NewColor(255, 0, 0)
	params.Transparent = true
	params.Opacity = DefaultTransparency
	return three.NewLineBasicMaterial(params)
}
