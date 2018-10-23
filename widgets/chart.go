package widgets

import (
	"fmt"

	charts "github.com/cnguy/gopherjs-frappe-charts"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Chart represents a wrapper for Frapper Charts library.
type Chart struct {
	vecty.Core

	data *charts.ChartData
}

func NewChart(data *charts.ChartData) *Chart {
	return &Chart{
		data: data,
	}
}

// Render implements the vecty.Component interface for Chart.
func (c *Chart) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			prop.ID("chart"),
		),
	)
}

// Mount implements the vecty.Mounter interface for Chart. Triggers when DOM has been created for the component.
func (c *Chart) Mount() {
	fmt.Println("Attaching chart...")
	charts.NewLineChart("#chart", c.data).
		WithTitle("Test chart").
		WithColors([]string{"blue"}).
		SetShowDots(false).
		SetHeatline(true).
		SetRegionFill(true).
		SetXAxisMode("tick").
		SetYAxisMode("span").
		SetIsSeries(true).
		Render()
}
