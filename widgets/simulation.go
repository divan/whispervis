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
	inProgress bool
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
	return Widget(
		elem.Div(
			Header("Simulation backend:"),
			elem.Div(
				elem.Input(
					vecty.Markup(
						vecty.MarkupIf(s.inProgress,
							vecty.Attribute("disabled", "true"),
						),
						prop.Value(s.address),
						vecty.Attribute("placehoder", "backend url"),
						event.Input(s.onEditInput),
						vecty.Class("input", "is-small"),
					),
				),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("pure-markup-group", "pure-u-1"),
			),
			elem.Button(
				vecty.Markup(
					vecty.MarkupIf(s.inProgress,
						vecty.Attribute("disabled", "true"),
						vecty.Class("is-loading"),
					),
					vecty.Class("button", "is-info", "is-small"),
					event.Click(s.onSimulateClick),
				),
				vecty.Text("Start simulation"),
			),
			vecty.If(s.hasResults,
				elem.Button(
					vecty.Markup(
						vecty.Class("button", "is-success", "is-small"),
						event.Click(s.onRestartClick),
					),
					vecty.Text("Replay"),
				),
			),
			elem.Break(),
			vecty.If(s.inProgress, elem.Div(
				vecty.Markup(
					vecty.Class("notification", "is-success"),
				),
				vecty.Text("Running simulation..."),
			)),
			elem.Div(
				vecty.If(s.errMsg != "", elem.Div(
					vecty.Markup(
						vecty.Class("notification", "is-danger"),
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
		s.inProgress = true
		vecty.Rerender(s)

		err := s.startSimulation()
		if err != nil {
			s.errMsg = err.Error()
		}

		s.hasResults = err == nil
		s.inProgress = false
		vecty.Rerender(s)
	}()
}

func (s *Simulation) onRestartClick(e *vecty.Event) {
	go s.replay()
}
