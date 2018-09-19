//go:generate go-bindata data/
package main

import (
	"bytes"
	"fmt"
	"io"

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
	content, err := Asset(file)
	if err != nil {
		return nil, fmt.Errorf("open bindata '%s': %v", file, err)
	}

	r := bytes.NewReader(content)

	n, err := LoadNetworkFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("open file '%s': %v", file, err)
	}
	n.Name = file
	return n, nil
}

// LoadNetworkFromReader loads network information from the io.Reader.
func LoadNetworkFromReader(r io.Reader) (*Network, error) {
	g, err := formats.FromD3JSONReader(r)
	if err != nil {
		return nil, fmt.Errorf("parse JSON: %v", err)
	}

	desc := "TBD"
	return &Network{
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
