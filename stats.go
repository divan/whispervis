package main

import (
	"github.com/status-im/simulation/propagation"
	"github.com/status-im/simulation/stats"
)

func (p *Page) RecalculateStats(plog *propagation.Log) *stats.Stats {
	net := p.network.Current()
	nodes := len(net.Data.Nodes())
	links := len(net.Data.Links())

	return stats.Analyze(plog, nodes, links)
}
