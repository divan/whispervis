package main

import (
	"fmt"
	"sort"

	charts "github.com/cnguy/gopherjs-frappe-charts"
	"github.com/divan/graphx/graph"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/status-im/simulation/propagation"
	"github.com/status-im/whispervis/widgets"
)

// StatsPage is stats view component.
type StatsPage struct {
	vecty.Core

	width, height string

	chart1Data *charts.ChartData
	chart2Data *charts.ChartData
}

// NewStatsPage creates and inits new stats page.
func NewStatsPage() *StatsPage {
	width, height := PageViewSize()
	return &StatsPage{
		width:  fmt.Sprintf("%dpx", width),
		height: fmt.Sprintf("%dpx", height),
	}
}

// Render implements the vecty.Component interface.
func (s *StatsPage) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("width", s.width),
			vecty.Style("height", s.height),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("title", "has-text-centered"),
			),
			vecty.Text("Stats page"),
		),
		// consult this tile madness here https://bulma.io/documentation/layout/tiles/
		elem.Div(vecty.Markup(vecty.Class("tile", "is-anscestor")),
			elem.Div(vecty.Markup(vecty.Class("tile", "is-parent", "is-4", "is-vertical")),
				Tile(vecty.Text("Part left")),
				Tile(vecty.Text("Part left 2")),
			),

			elem.Div(vecty.Markup(vecty.Class("tile", "is-parent", "is-vertical")),
				vecty.If(s.chart1Data != nil, Tile(widgets.NewChart("nodes", s.chart1Data))),
				vecty.If(s.chart2Data != nil, Tile(widgets.NewChart("cumulative", s.chart2Data))),
			),
		),
	)
}

// UpdateStats updates stats page with a new data.
func (s *StatsPage) UpdateStats(g *graph.Graph, plog *propagation.Log) {
	sort.Sort(plog)

	labels := make([]string, len(plog.Timestamps))

	nodeCounts := make([]float64, len(plog.Timestamps))
	linkCounts := make([]float64, len(plog.Timestamps))
	nodeCum := make([]float64, len(plog.Timestamps))
	linkCum := make([]float64, len(plog.Timestamps))
	var totalNode, totalLink int64
	for i, ts := range plog.Timestamps {
		labels[i] = fmt.Sprintf("%d", ts)
		nodes := len(plog.Nodes[i])
		links := len(plog.Links[i])
		totalNode += int64(nodes)
		totalLink += int64(links)
		nodeCounts[i] = float64(nodes)
		linkCounts[i] = float64(links)
		nodeCum[i] = float64(totalNode)
		linkCum[i] = float64(totalLink)
	}

	data := charts.NewChartData()
	data.Labels = labels
	dataset1 := charts.NewDataset("Nodes", nodeCounts)
	dataset2 := charts.NewDataset("Links", linkCounts)
	data.Datasets = []*charts.Dataset{
		dataset1,
		dataset2,
	}
	s.chart1Data = data

	data2 := charts.NewChartData()
	data2.Labels = labels
	dataset3 := charts.NewDataset("Nodes", nodeCum)
	dataset4 := charts.NewDataset("Links", linkCum)
	data2.Datasets = []*charts.Dataset{
		dataset3,
		dataset4,
	}
	s.chart2Data = data2
}
