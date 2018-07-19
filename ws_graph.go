package main

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/graph"
	"github.com/gorilla/websocket"
)

func (ws *WSServer) sendGraphData(c *websocket.Conn) {
	var buf bytes.Buffer
	d3json := formats.NewD3JSON(&buf, false)
	err := d3json.ExportGraph(ws.graph)
	if err != nil {
		log.Fatal("Can't marshal graph to JSON")
	}
	msg := &WSResponse{
		Type:  RespGraph,
		Graph: json.RawMessage(buf.Bytes()),
	}

	ws.sendMsg(c, msg)
}

func (ws *WSServer) updateGraph(g *graph.Graph) {
	ws.graph = g

	ws.broadcastGraphData()
}

func (ws *WSServer) broadcastGraphData() {
	for i := 0; i < len(ws.hub); i++ {
		ws.sendGraphData(ws.hub[i])
	}
}
