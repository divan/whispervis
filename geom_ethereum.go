package main

import "github.com/divan/three"

// NewEthereumGeometry creates a geometry for representing Ethereum node (non-regular 3D octahedron).
// It scales object according to the given scale.
func NewEthereumGeometry(scale float64) three.Geometry {
	var geom = three.NewBasicGeometry(three.BasicGeometryParams{})
	vertices := []struct {
		x, y, z float64
	}{
		{scale * 1.0, 0.0, 0.0},
		{scale * -1.0, 0.0, 0.0},
		{0.0, scale * 1.5, 0.0},
		{0.0, scale * -1.5, 0.0},
		{0.0, 0.0, scale * 1.0},
		{0.0, 0.0, scale * -1.0},
	}

	for _, v := range vertices {
		geom.AddVertice(v.x, v.y, v.z)
	}

	faces := []struct {
		a, b, c int
	}{
		{0, 2, 4},
		{0, 4, 3},
		{0, 3, 5},
		{0, 5, 2},
		{1, 2, 5},
		{1, 5, 3},
		{1, 3, 4},
		{1, 4, 2},
	}

	for _, f := range faces {
		geom.AddFace(f.a, f.b, f.c)
	}

	geom.ComputeBoundingSphere()
	geom.ComputeFaceNormals()

	return geom
}
