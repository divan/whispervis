package main

import (
	"fmt"
	"path"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/graph"
)

// Network represents network graph and information, used for
// for simulation and visualization.
type Network struct {
	Name        string
	Description string
	Data        *graph.Graph

	NodesCount int
}

// LoadNetwork loads network information from the JSON file.
// JSON format is specified in graphx/formats package.
func LoadNetwork(file string) (*Network, error) {
	name := path.Base(file)
	desc := ""

	g, err := formats.FromD3JSON(file)
	if err != nil {
		return nil, fmt.Errorf("open '%s': %v", file, err)
	}

	return &Network{
		Name:        name,
		Description: desc,
		Data:        g,
	}, nil
}
