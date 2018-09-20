package widgets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"github.com/status-im/simulation/propagation"
)

// Simulation represents configuration panel for propagation simulation.
type Simulation struct {
	vecty.Core
	networkFn    func() []byte // function that returns current network JSON description
	simulationFn func(*propagation.Log)

	address string           // backend host address
	plog    *propagation.Log // last simulation result
}

// NewSimulation creates new simulation configuration panel. If simulation
// backend host address is not specified, it'll use 'localhost:8084' as a default.
func NewSimulation(address string, networkFn func() []byte, simulationFn func(*propagation.Log)) *Simulation {
	if address == "" {
		address = "http://localhost:8084"
	}
	return &Simulation{
		address:      address,
		networkFn:    networkFn,
		simulationFn: simulationFn,
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
	if s.networkFn == nil {
		return
	}

	go s.runSimulation()
}

func (s *Simulation) onRestartClick(e *vecty.Event) {
	go s.play()
}

// runSimulation starts whisper message propagation simulation,
// remotely talking to simulation backend.
func (s *Simulation) runSimulation() {
	payload := s.networkFn()
	buf := bytes.NewBuffer(payload)
	url := "http://" + s.address + "/"
	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		fmt.Println("[ERROR] POST request to simulation backend:", err)
		return
	}

	var plog propagation.Log
	err = json.NewDecoder(resp.Body).Decode(&plog)
	if err != nil {
		fmt.Println("[ERROR] decoding response from simulation backend:", err)
		return
	}

	var max int
	for _, ts := range plog.Timestamps {
		if ts > max {
			max = ts
		}
	}

	timespan := time.Duration(max) * time.Millisecond
	fmt.Printf("Whoa! Got results! %d timestamps over %v\n", len(plog.Timestamps), timespan)
	s.plog = &plog
	s.play()
}

func (s *Simulation) play() {
	s.simulationFn(s.plog)
}
