package main

import (
	"flag"
	"log"

	"github.com/divan/graphx/formats"
	"github.com/divan/graphx/layout"
)

func main() {
	bind := flag.String("bind", ":20002", "Port to bind to")
	iterations := flag.Int("i", 600, "Graph layout iterations to run (0 = auto, buggy)")
	flag.Parse()

	data, err := formats.FromD3JSON("network.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded graph: %d nodes, %d links\n", len(data.Nodes()), len(data.Links()))

	plog, err := LoadPropagationData("propagation.json")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded propagation data: %d timestamps\n", len(plog.Timestamps))

	log.Printf("Initializing layout...")
	l := layout.NewAuto(data)

	ws := NewWSServer(l)
	if *iterations == 0 {
		ws.layout.Calculate()
	} else {
		ws.layout.CalculateN(*iterations)
	}
	ws.updatePositions()
	ws.updateGraph(data)
	ws.updatePropagationData(plog)

	log.Printf("Starting web server on %s...", *bind)
	startWeb(ws, *bind)
	select {}
}
