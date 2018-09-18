package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

// NetworkSelector represents widget for choosing or uploading network topology
// to be used for visualization.
type NetworkSelector struct {
	vecty.Core

	description string
	value       string
}

// NewNetworkSelector creates new NetworkSelector.
func NewNetworkSelector(defaultValue string) *NetworkSelector {
	return &NetworkSelector{
		value: defaultValue,
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
			elem.Option(
				vecty.Markup(
					vecty.Property("value", "net100"),
					vecty.Property("selected", n.value == "net100"),
					vecty.Property("description", "Random network consisting from 100 nodes, 4-5 connections each"),
				),
				vecty.Text("Random network: 100 nodes"),
			),
			elem.Option(
				vecty.Markup(
					vecty.Property("value", "net300"),
					vecty.Property("selected", n.value == "net300"),
					vecty.Property("description", "Random network consisting from 300 nodes, 4-5 connections each"),
				),
				vecty.Text("Random network: 300 nodes"),
			),
			elem.Option(
				vecty.Markup(
					vecty.Property("value", "3dgrid"),
					vecty.Property("selected", n.value == "3dgrid"),
					vecty.Property("description", "5x5 3D cube, 125 nodes in total"),
				),
				vecty.Text("3D cube graph: 125 nodes"),
			),
			elem.Option(
				vecty.Markup(
					vecty.Property("value", "upload"),
					vecty.Property("selected", n.value == "upload"),
					vecty.Property("description", "Upload custom network topology..."),
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
		vecty.Text(n.description),
	)
}

// onChange implements handler for select input changed value
func (n *NetworkSelector) onChange(e *vecty.Event) {
	var desc string
	value := e.Target.Get("value").String()
	for i := 0; i < e.Target.Length(); i++ {
		optValue := e.Target.Index(i).Get("value").String()
		if optValue == value {
			desc = e.Target.Index(i).Get("description").String()
		}
	}

	n.value = value
	n.description = desc
	vecty.Rerender(n)
}
