package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/divan/graphx/formats"
	"github.com/divan/simulation/propagation"
	"github.com/divan/simulation/stats"
)

// Simulation represents the state of last simulation.
type Simulation struct {
	plog  *propagation.Log
	stats *stats.Stats
}

// SimulationRequests defines a POST request payload for simulation backend.
type SimulationRequest struct {
	Algorithm string          `json:"algorithm"`
	SenderIdx int             `json:"senderIdx"` // index of the sender node (index of data.Nodes, in fact)
	TTL       int             `json:"ttl"`       // ttl in seconds
	MsgSize   int             `json:"msg_size"`  // msg size in bytes
	Network   json.RawMessage `json:"network"`   // current network graph
}

// runSimulation starts whisper message propagation simulation,
// remotely talking to simulation backend with given address.
func (p *Page) runSimulation(address string, ttl int) (*Simulation, error) {
	buf, err := p.newPOSTSimulationRequest(ttl)
	if err != nil {
		return nil, fmt.Errorf("Internal error. See console output.")
	}
	resp, err := http.Post(address, "application/json", buf)
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

// newPOSTSimulationRequest generates SimulationReqeust and
// prepares it as io.Reader for usage with http.Post.
func (p *Page) newPOSTSimulationRequest(ttl int) (io.Reader, error) {
	req := SimulationRequest{
		Algorithm: "whisperv6", // TODO(divan): move to UI configuration
		SenderIdx: 0,
		TTL:       ttl,
		MsgSize:   400,
		Network:   p.currentNetworkJSON(),
	}
	payload, err := json.Marshal(req)
	if err != nil {
		fmt.Println("[ERROR] Can't marshal SimulationRequest", err)
		return nil, err
	}

	buf := bytes.NewBuffer(payload)

	return buf, nil
}

// currentNetworkJSON returns JSON encoded description of the current graph/network.
func (p *Page) currentNetworkJSON() []byte {
	net := p.network.Current().Data
	var buf bytes.Buffer
	err := formats.NewD3JSON(&buf, true).ExportGraph(net)
	if err != nil {
		fmt.Println("[ERROR] Can't export graph:", err)
		return nil
	}
	return buf.Bytes()
}
