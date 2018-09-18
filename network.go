//go:generate go-bindata data/
package main

import (
	"bytes"
	"fmt"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/graph"
)

// Network represents network graph and information, used for
// for simulation and visualization.
type Network struct {
	Name        string
	Description string
	Data        *graph.Graph
}

// LoadNetwork loads network information from the JSON file.
// JSON format is specified in graphx/formats package.
func LoadNetwork(file string) (*Network, error) {
	desc := ""

	content, err := Asset(file)
	if err != nil {
		return nil, fmt.Errorf("open bindata '%s': %v", file, err)
	}

	r := bytes.NewReader(content)

	g, err := formats.FromD3JSONReader(r)
	if err != nil {
		return nil, fmt.Errorf("open '%s': %v", file, err)
	}

	return &Network{
		Name:        file,
		Description: desc,
		Data:        g,
	}, nil
}

// String implements Stringer for Network.
func (n *Network) String() string {
	return fmt.Sprintf("[%s: %s] - %d nodes, %d links", n.Name, n.Description, n.NodesCount(), n.LinksCount())
}

// NodesCount returns number of the nodes in the network.
func (n *Network) NodesCount() int {
	if n.Data == nil {
		return 0
	}
	return len(n.Data.Nodes())
}

// LinksCount returns number of the links in the network.
func (n *Network) LinksCount() int {
	if n.Data == nil {
		return 0
	}
	return len(n.Data.Links())
}
