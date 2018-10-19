package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/divan/graphx/formats"
	"github.com/status-im/simulation/propagation"
	"github.com/status-im/simulation/stats"
)

// Simulation represents the state of last simulation.
type Simulation struct {
	plog  *propagation.Log
	stats *stats.Stats
}

// runSimulation starts whisper message propagation simulation,
// remotely talking to simulation backend with given address.
func (p *Page) runSimulation(address string) (*Simulation, error) {
	payload := p.currentNetworkJSON()
	buf := bytes.NewBuffer(payload)
	host := "http://" + address + "/"
	resp, err := http.Post(host, "application/json", buf)
	if err != nil {
		fmt.Println("[ERROR] POST request to simulation backend:", err)
		return nil, fmt.Errorf("Backend error. Did you run backend?")
	}

	var plog propagation.Log
	err = json.NewDecoder(resp.Body).Decode(&plog)
	if err != nil {
		fmt.Println("[ERROR] decoding response from simulation backend:", err)
		return nil, fmt.Errorf("decoding JSON response error: %v", err)
	}

	return &Simulation{
		plog: &plog,
	}, nil
}

// currentNetworkJSON returns JSON encoded description of the current graph/network.
func (p *Page) currentNetworkJSON() []byte {
	net := p.network.current.Data
	var buf bytes.Buffer
	err := formats.NewD3JSON(&buf, true).ExportGraph(net)
	if err != nil {
		fmt.Println("[ERROR] Can't export graph:", err)
		return nil
	}
	return buf.Bytes()
}
