package main

import (
	"bytes"
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/status-im/whispervis/widgets"
)

// NetworkSelector represents widget for choosing or uploading network topology
// to be used for visualization.
type NetworkSelector struct {
	vecty.Core

	current  *Network
	isCustom bool
	networks map[string]*Network

	upload *widgets.UploadWidget

	handler func(*Network) // executed on network change
}

// NewNetworkSelector creates new NetworkSelector.
func NewNetworkSelector(handler func(*Network)) *NetworkSelector {
	current := &Network{}
	networks, err := LoadNetworks()
	if err != nil {
		fmt.Println("No networks loaded:", err)
	} else {
		current = networks["net100.json"]
	}

	ns := &NetworkSelector{
		networks: networks,
		current:  current,
		handler:  handler,
	}
	ns.upload = widgets.NewUploadWidget(ns.onUpload)
	ns.setCurrentNetwork(current)
	return ns
}

// Render implements the vecty.Component interface.
func (n *NetworkSelector) Render() vecty.ComponentOrHTML {
	return widgets.Widget(
		widgets.Header("Network graph:"),
		elem.Div(
			vecty.Markup(
				vecty.Class("select"),
				event.Change(n.onChange),
			),
			elem.Select(
				vecty.Markup(
					event.Change(n.onChange),
				),
				n.networkOptions(),
				elem.Option(
					vecty.Markup(
						vecty.Property("value", "upload"),
						vecty.Property("selected", n.isCustom),
					),
					vecty.Text("Upload custom..."),
				),
			),
		),
		n.descriptionBlock(),
		vecty.If(n.isCustom, n.upload),
	)
}

// descriptionBlock renders the block with network description.
func (n *NetworkSelector) descriptionBlock() *vecty.HTML {
	if n.current == nil || n.current.Description == "" {
		return elem.Span()
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("is-small", "is-marginless"),
		),
		elem.Div(
			vecty.Text(n.current.Description),
		),
	)
}

// LoadNetworks imports preloaded neworks from the directory with JSON files.
func LoadNetworks() (map[string]*Network, error) {
	files, err := AssetDir("data")
	if err != nil {
		return nil, err
	}

	networks := map[string]*Network{}
	for _, file := range files {
		network, err := LoadNetwork("data/" + file)
		if err != nil {
			return nil, fmt.Errorf("load network: %v", err)
		}

		networks[file] = network
	}
	return networks, nil
}

// networkOptions renders 'option' elements for network 'select' input tag.
func (n *NetworkSelector) networkOptions() vecty.List {
	var options vecty.List
	for name := range n.networks {
		options = append(options, elem.Option(
			vecty.Markup(
				vecty.Property("value", name),
				vecty.Property("selected", n.current.Name == "data/"+name), // TODO(divan): get rid of "data"
			),
			vecty.Text(name),
		))
	}
	return options
}

// onChange implements handler for select input changed value
func (n *NetworkSelector) onChange(e *vecty.Event) {
	value := e.Target.Get("value").String()

	if value == "upload" {
		n.isCustom = true
		vecty.Rerender(n)
		return
	}

	n.isCustom = false

	net := n.networks[value]
	n.setCurrentNetwork(net)

	vecty.Rerender(n)
}

// onUpload implements callback for "Upload" button clicked event.
func (n *NetworkSelector) onUpload(json []byte) {
	r := bytes.NewReader(json)
	net, err := LoadNetworkFromReader(r)
	if err != nil {
		fmt.Printf("[ERROR] Load network: %v", err)
	}

	net.Name = fmt.Sprintf("Uploaded (%d nodes)", net.NodesCount())

	n.networks[net.Name] = net
	n.setCurrentNetwork(net)

	vecty.Rerender(n)
}

// setCurrentNetwork changes current network and runs needed update handlers.
func (n *NetworkSelector) setCurrentNetwork(net *Network) {
	n.current = net

	if n.handler != nil {
		go n.handler(n.current)
	}
}
