package main

import (
	"github.com/gorilla/websocket"
)

func (ws *WSServer) sendPropagationData(c *websocket.Conn) {
	msg := &WSResponse{
		Type:        RespPropagation,
		Propagation: ws.propagation,
	}

	ws.sendMsg(c, msg)
}

func (ws *WSServer) updatePropagationData(plog *PropagationLog) {
	ws.propagation = plog

	ws.broadcastPropagationData()
}

func (ws *WSServer) broadcastPropagationData() {
	for i := 0; i < len(ws.hub); i++ {
		ws.sendPropagationData(ws.hub[i])
	}
}
