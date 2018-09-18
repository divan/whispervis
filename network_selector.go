package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// NetworkSelector represents widget for choosing or uploading network topology
// to be used for visualization.
type NetworkSelector struct {
	vecty.Core

	current  *Network
	isCustom bool
	networks map[string]*Network

	handler func(*Network)
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

	return &NetworkSelector{
		networks: networks,
		current:  current,
		handler:  handler,
	}
}

// Render implements the vecty.Component interface.
func (n *NetworkSelector) Render() vecty.ComponentOrHTML {
	return elem.Div(
		elem.Heading3(vecty.Text("Choose network:")),
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
		n.descriptionBlock(),
		elem.HorizontalRule(),
	)
}

func (n *NetworkSelector) descriptionBlock() *vecty.HTML {
	return elem.Div(
		vecty.Markup(
			vecty.Style("padding", "7px"),
			vecty.Style("font-style", "italic"),
			vecty.Style("color", "blue"),
		),
		vecty.Text(n.current.Description),
	)
}

// onChange implements handler for select input changed value
func (n *NetworkSelector) onChange(e *vecty.Event) {
	value := e.Target.Get("value").String()

	if value == "upload" {
		n.isCustom = true
		// TODO(divan): handle uploaded network
	} else {
		n.isCustom = false
		n.current = n.networks[value]
	}

	if n.handler != nil {
		go n.handler(n.current)
	}

	vecty.Rerender(n)
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

func (n *NetworkSelector) networkOptions() vecty.List {
	var options vecty.List
	for name, _ := range n.networks {
		options = append(options, elem.Option(
			vecty.Markup(
				vecty.Property("value", name),
				vecty.Property("selected", n.current.Name == name),
			),
			vecty.Text(name),
		))
	}
	return options
}
