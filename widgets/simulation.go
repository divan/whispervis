package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// Simulation represents configuration panel for propagation simulation.
type Simulation struct {
	vecty.Core
	startSimulation func() error
	replay          func()

	address string // backend host address

	errMsg     string
	hasResults bool
}

// NewSimulation creates new simulation configuration panel. If simulation
// backend host address is not specified, it'll use 'localhost:8084' as a default.
func NewSimulation(address string, startSimulation func() error, replay func()) *Simulation {
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
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("pure-markup-group", "pure-u-1"),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("pure-button"),
					vecty.Class("pure-u-1-2"),
					vecty.Style("background", "rgb(28, 184, 65)"),
					vecty.Style("color", "white"),
					vecty.Style("border-radius", "4px"),
					event.Click(s.onSimulateClick),
				),
				vecty.Text("Start simulation"),
			),
			vecty.If(s.hasResults,
				elem.Button(
					vecty.Markup(
						vecty.Class("pure-button"),
						vecty.Class("pure-u-1-3"),
						vecty.Style("background", "rgb(28, 184, 65)"),
						vecty.Style("color", "white"),
						vecty.Style("margin", "10px"),
						vecty.Style("border-radius", "4px"),
						event.Click(s.onRestartClick),
					),
					vecty.Text("Replay"),
				),
			),
			elem.Break(),
			elem.Div(
				vecty.If(s.errMsg != "", elem.Paragraph(
					vecty.Markup(
						vecty.Style("background", "rgb(202, 60, 60)"),
						vecty.Style("color", "white"),
						vecty.Style("border-radius", "4px"),
						vecty.Style("margin-right", "5px"),
						vecty.Style("padding", "5px"),
					),
					vecty.Text(s.errMsg),
				)),
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
	go func() {
		s.errMsg = ""
		s.hasResults = false
		vecty.Rerender(s)

		err := s.startSimulation()
		if err != nil {
			s.errMsg = err.Error()
		}

		s.hasResults = err == nil
		vecty.Rerender(s)
	}()
}

func (s *Simulation) onRestartClick(e *vecty.Event) {
	go s.replay()
}
