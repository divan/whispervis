package main

import (
	"github.com/gorilla/websocket"
)

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
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
	nodes := ws.layout.Positions()
	ws.Positions = nodes
	/*
		positions := []*position{}
		for _, node := range nodes {
			log.Println("node", node)
			pos := &position{
				X: node.X,
				Y: node.Y,
				Z: node.Z,
			}
			positions = append(positions, pos)
		}
		ws.Positions = positions
	*/

	ws.broadcastPositions()
}

func (ws *WSServer) broadcastPositions() {
	for i := 0; i < len(ws.hub); i++ {
		ws.sendPositions(ws.hub[i])
	}
}
