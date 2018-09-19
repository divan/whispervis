package main

import (
	"fmt"
)

// onUpload implements callback for "Upload" button clicked event.
func (p *Page) onUpload(json []byte) {
	if err := p.UpdateNetworkGraph(json); err != nil {
		fmt.Println("[ERROR] Failed to process network.json:", err)
		return
	}
	fmt.Println("Uploaded new network graph")
}
