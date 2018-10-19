package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/status-im/simulation/propagation"
)

// Simulation represents configuration panel for propagation simulation.
type Simulation struct {
	vecty.Core
	startSimulation func()
	replay          func()

	address string           // backend host address
	plog    *propagation.Log // last simulation result
}

// NewSimulation creates new simulation configuration panel. If simulation
// backend host address is not specified, it'll use 'localhost:8084' as a default.
func NewSimulation(address string, startSimulation, replay func()) *Simulation {
	if address == "" {
		address = "http://localhost:8084"
	}
	return &Simulation{
		address:         address,
		startSimulation: startSimulation,
		replay:          replay,
	}
}

// Render implements vecty.Component interface for Simulation.
func (s *Simulation) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Div(
			elem.Heading3(vecty.Text("Simulation backend:")),
			vecty.Markup(
				vecty.Class("pure-markup-group", "pure-u-1"),
			),
			elem.Label(
				vecty.Markup(
					vecty.Class("pure-u-1-2"),
				),
				vecty.Text("Host address:"),
			),
			elem.Input(
				vecty.Markup(
					prop.Value(s.address),
					event.Input(s.onEditInput),
					vecty.Class("pure-input-1-3"),
					vecty.Style("float", "right"),
					vecty.Style("margin-right", "10px"),
					vecty.Style("text-align", "right"),
				),
			),
			elem.Break(),
			elem.Button(
				vecty.Markup(
					vecty.Class("pure-button"),
					vecty.Style("background", "rgb(28, 184, 65)"),
					vecty.Style("color", "white"),
					vecty.Style("border-radius", "4px"),
					event.Click(s.onSimulateClick),
				),
				vecty.Text("Start simulation"),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("pure-button"),
					vecty.Style("background", "rgb(28, 184, 65)"),
					vecty.Style("color", "white"),
					vecty.Style("border-radius", "4px"),
					event.Click(s.onRestartClick),
				),
				vecty.Text("Replay"),
			),
		),
	)
}

func (s *Simulation) onEditInput(event *vecty.Event) {
	value := event.Target.Get("value").String()

	s.address = value
}

// Address returns the current backend address.
func (s *Simulation) Address() string {
	return s.address
}

func (s *Simulation) onSimulateClick(e *vecty.Event) {
	go s.startSimulation()
}

func (s *Simulation) onRestartClick(e *vecty.Event) {
	go s.replay()
}
