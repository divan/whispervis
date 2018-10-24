package widgets

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/status-im/simulation/stats"
)

// Stats represents widget with the statistics from the latest simulation.
type Stats struct {
	vecty.Core

	stats *stats.Stats
}

// NewStats create a new Stats widget.
func NewStats() *Stats {
	return &Stats{
		/*
			stats: &stats.Stats{
				Time:         1234 * time.Millisecond,
				NodeCoverage: stats.NewCoverage(100, 200),
				LinkCoverage: stats.NewCoverage(100, 200),
			},
		*/
	}
}

// Update updates the stats based on the propagation log info and current graph.
func (s *Stats) Update(stats *stats.Stats) {
	s.stats = stats
	vecty.Rerender(s)
}

// Render implements vecty.Component interface for Stats.
func (s *Stats) Render() vecty.ComponentOrHTML {
	if s.stats == nil {
		return elem.Div()
	}
	return Widget(
		Header("Stats:"),
		elem.Table(
			vecty.Markup(
				vecty.Class("table", "is-hoverable", "is-fullwidth"),
			),
			elem.TableBody(
				s.tableRow("Elapsed time:", s.stats.Time),
				s.tableRow("Nodes hit:", s.stats.NodeCoverage.Actual),
				s.tableRow("Links hit:", s.stats.LinkCoverage.Actual),
			),
		),
	)
}

func (s *Stats) tableRow(label string, value interface{}) *vecty.HTML {
	return elem.TableRow(
		elem.TableData(
			vecty.Markup(
				vecty.Style("font-weight", "bold"),
			),
			vecty.Text(label),
		),
		elem.TableData(
			vecty.Markup(
				vecty.Style("text-align", "right"),
			),
			vecty.Text(
				fmt.Sprintf("%v", value),
			),
		),
	)
}

func (s *Stats) Reset() {
	s.stats = nil
}
