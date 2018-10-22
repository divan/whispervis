package main

import (
	"encoding/json"
	"errors"
	"io"
	"runtime"

	"github.com/divan/graphx/graph"
	"github.com/gopherjs/vecty"
)

// UpdateGraph starts graph layout simulation.
func (p *Page) UpdateGraph() {
	p.loader.Reset()
	p.loaded = false
	vecty.Rerender(p)

	config := p.forceEditor.Config()
	p.loader.SetSteps(config.Steps)
	for i := 0; i < config.Steps; i++ {
		p.layout.UpdatePositions()
		p.loader.Inc()
		vecty.Rerender(p.loader)
		runtime.Gosched()
	}
	p.loaded = true
	// TODO(divan): remove previous objects
	p.webgl.RemoveObjects()
	p.webgl.CreateObjects(p.layout.Positions(), p.layout.Links())

	vecty.Rerender(p)
}

// ApplyForces applies current forces to the objects, and runs
// a single simulation run to update positions.
func (p *Page) ApplyForces() {
	fc := p.forceEditor.Config()
	p.layout.SetConfig(fc.Config)
	p.layout.UpdatePositions()
	p.webgl.updatePositions()
	p.webgl.rt.Disable()
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
