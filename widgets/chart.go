package widgets

import (
	"strings"

	charts "github.com/cnguy/gopherjs-frappe-charts"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Chart represents a wrapper for Frapper Charts library.
type Chart struct {
	vecty.Core

	id   string
	name string
	data *charts.ChartData
}

func NewChart(name string, data *charts.ChartData) *Chart {
	id := strings.ToLower(name) // TODO(divan): make it unique
	return &Chart{
		id:   id,
		name: name,
		data: data,
	}
}

// Render implements the vecty.Component interface for Chart.
func (c *Chart) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			prop.ID(c.id),
		),
	)
}

// Mount implements the vecty.Mounter interface for Chart. Triggers when DOM has been created for the component.
func (c *Chart) Mount() {
	charts.NewLineChart("#"+c.id, c.data).
		WithTitle(c.name).
		WithColors([]string{"blue"}).
		SetShowDots(false).
		SetHeatline(true).
		SetRegionFill(true).
		SetXAxisMode("tick").
		SetYAxisMode("span").
		SetIsSeries(true).
		Render()
}
