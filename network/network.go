//go:generate go-bindata -pkg network data/
package network

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"

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

// LoadNetworks imports preloaded neworks from the directory with JSON files.
func LoadNetworks() (map[string]*Network, error) {
	files, err := AssetDir("data")
	if err != nil {
		return nil, err
	}

	networks := map[string]*Network{}
	for _, file := range files {
		network, err := LoadNetwork("data/" + file)
		if err != nil {
			return nil, fmt.Errorf("load network: %v", err)
		}

		networks[file] = network
	}
	return networks, nil
}

// LoadNetworkFromReader loads network information from the io.Reader.
func LoadNetworkFromReader(r io.Reader) (*Network, error) {
	g, desc, err := GraphFromJSON(r)
	if err != nil {
		return nil, fmt.Errorf("parse JSON: %v", err)
	}

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

// GraphFromJSON is a custom version of graphx JSON importer, as we want to use
// some additional fields (Description).
// TODO(divan): that's probably can be done better within the limits of graphx library.
func GraphFromJSON(r io.Reader) (*graph.Graph, string, error) {
	// decode into temporary struct to process
	var res struct {
		Description string             `json:"description"`
		Nodes       []*graph.BasicNode `json:"nodes"`
		Links       []*struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"links"`
	}
	err := json.NewDecoder(r).Decode(&res)
	if err != nil {
		return nil, "", err
	}

	if len(res.Nodes) == 0 {
		return nil, "", errors.New("empty graph")
	}

	// convert links IDs into indices
	g := graph.NewGraphMN(len(res.Nodes), len(res.Links))

	for _, node := range res.Nodes {
		g.AddNode(node)
	}

	for _, link := range res.Links {
		err := g.AddLink(link.Source, link.Target)
		if err != nil {
			return nil, "", err
		}
	}

	g.UpdateCache()

	return g, res.Description, nil
}
