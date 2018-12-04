package main

import (
	"github.com/divan/simulation/propagation"
	"github.com/divan/simulation/stats"
)

func (p *Page) RecalculateStats(plog *propagation.Log) *stats.Stats {
	net := p.network.Current()

	return stats.Analyze(plog, net.Data.NumNodes(), net.Data.NumLinks())
}
