package main

import (
	"encoding/json"
	"os"
)

// PropagationLog represnts log of p2p message propagation
// with relative timestamps (starting from T0).
type PropagationLog struct {
	Timestamps []int   // timestamps in milliseconds starting from T0
	Indices    [][]int // indices of links for each step, len should be equal to len of Timestamps field
	Nodes      [][]int // indices of nodes involved
}

func LoadPropagationData(filename string) (*PropagationLog, error) {
	fd, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	var plog PropagationLog
	err = json.NewDecoder(fd).Decode(&plog)
	if err != nil {
		return nil, err
	}

	return &plog, nil
}
