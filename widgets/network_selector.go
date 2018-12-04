package widgets

import (
	"bytes"
	"fmt"

	"github.com/divan/whispervis/network"
	"github.com/divan/whispervis/storage"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// NetworkSelector represents widget for choosing or uploading network topology
// to be used for visualization.
type NetworkSelector struct {
	vecty.Core

	current  *network.Network
	isCustom bool
	networks map[string]*network.Network

	upload *UploadWidget

	handler func(*network.Network) // executed on network change
}

// NewNetworkSelector creates new NetworkSelector.
func NewNetworkSelector(handler func(*network.Network)) *NetworkSelector {
	current := &network.Network{}
	networks, err := network.LoadNetworks()
	if err != nil {
		fmt.Println("No networks loaded:", err)
	} else {
		current = networks[storage.Network()]
	}

	ns := &NetworkSelector{
		networks: networks,
		current:  current,
		handler:  handler,
	}
	ns.upload = NewUploadWidget(ns.onUpload)
	ns.setCurrentNetwork(current)
	return ns
}

// Render implements the vecty.Component interface.
func (n *NetworkSelector) Render() vecty.ComponentOrHTML {
	return Widget(
		Header("Network graph:"),
		elem.Div(
			vecty.Markup(
				vecty.Class("select", "is-fullwidth"),
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
			vecty.If(n.isCustom, vecty.Text("Upload custom graph...")),
			vecty.If(!n.isCustom, vecty.Text(n.current.Description)),
		),
	)
}

// networkOptions renders 'option' elements for network 'select' input tag.
func (n *NetworkSelector) networkOptions() vecty.List {
	var options vecty.List
	for name := range n.networks {
		currentName := "none"
		if n.current != nil {
			currentName = n.current.Name
		}

		opt := elem.Option(
			vecty.Markup(
				vecty.Property("value", name),
				vecty.Property("selected", currentName == "data/"+name), // TODO(divan): get rid of "data"
			),
			vecty.Text(name),
		)
		options = append(options, opt)
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
	n.handler(n.current)

	// save to localstorage
	storage.SetNetwork(value)

	vecty.Rerender(n)
}

// onUpload implements callback for "Upload" button clicked event.
func (n *NetworkSelector) onUpload(json []byte) {
	r := bytes.NewReader(json)
	net, err := network.LoadNetworkFromReader(r)
	if err != nil {
		fmt.Printf("[ERROR] Load network: %v", err)
	}

	net.Name = fmt.Sprintf("Uploaded (%d nodes)", net.NodesCount())

	n.networks[net.Name] = net
	n.setCurrentNetwork(net)

	if n.handler != nil {
		go n.handler(n.current)
	}

	vecty.Rerender(n)
}

// Current returns the currently selected network.
func (n *NetworkSelector) Current() *network.Network {
	return n.current
}

// setCurrentNetwork changes current network and runs needed update handlers.
func (n *NetworkSelector) setCurrentNetwork(net *network.Network) {
	n.current = net
}
