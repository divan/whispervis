package main

import (
	"flag"
	"log"

	"github.com/divan/graph-experiments/graph"
	"github.com/divan/graph-experiments/layout"
)

func main() {
	bind := flag.String("bind", ":20002", "Port to bind to")
	iterations := flag.Int("i", 600, "Graph layout iterations to run (0 = auto, buggy)")
	flag.Parse()

	data, err := graph.NewGraphFromJSON("network.json")
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
	repelling := layout.NewGravityForce(-100.0, layout.BarneHutMethod)
	springs := layout.NewSpringForce(0.01, 5.0, layout.ForEachLink)
	drag := layout.NewDragForce(0.4, layout.ForEachNode)
	layout3D := layout.New(data, repelling, springs, drag)

	ws := NewWSServer(layout3D)
	if *iterations == 0 {
		ws.layout.Calculate()
	} else {
		ws.layout.CalculateN(*iterations)
	}
	ws.updateGraph(data)
	ws.updatePropagationData(plog)

	log.Printf("Starting web server on %s...", *bind)
	startWeb(ws, *bind)
	select {}
}
