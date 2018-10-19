package main

import "github.com/status-im/simulation/stats"

func (p *Page) RecalculateStats() {
	sim := p.simulation
	if sim == nil || sim.plog == nil {
		return
	}

	net := p.network.current
	nodes := len(net.Data.Nodes())
	links := len(net.Data.Links())

	stats := stats.Analyze(sim.plog, nodes, links)
	p.simulation.stats = stats
}
