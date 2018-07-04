package main

import "github.com/gorilla/websocket"

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func (ws *WSServer) sendPositions(c *websocket.Conn) {
	msg := &WSResponse{
		Type:      RespPositions,
		Positions: ws.Positions,
	}

	ws.sendMsg(c, msg)
}

func (ws *WSServer) updatePositions() {
	// positions
	nodes := ws.layout.Nodes()
	positions := []*position{}
	for i := 0; i < len(nodes); i++ {
		pos := &position{
			X: nodes[i].X,
			Y: nodes[i].Y,
			Z: nodes[i].Z,
		}
		positions = append(positions, pos)
	}
	ws.Positions = positions

	ws.broadcastPositions()
}

func (ws *WSServer) broadcastPositions() {
	for i := 0; i < len(ws.hub); i++ {
		ws.sendPositions(ws.hub[i])
	}
}
