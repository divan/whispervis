package main

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

// uploadButton renders a upload button for network data.
func (p *Page) uploadButton() *vecty.HTML {
	return elem.Div(
		elem.Input(
			vecty.Markup(
				prop.ID("file"),
				prop.Type("file"),
				event.Input(p.onUploadClick),
			),
			vecty.Text("Upload network.json"),
		),
	)
}

// onUploadClick implements callback for "Upload" button clicked event.
func (p *Page) onUploadClick(e *vecty.Event) {
	// FIXME: run as a gorotine because GopherJS can't block JS thread in callback
	go func() {
		file := e.Target.Get("files").Index(0)
		fmt.Println("File size:", file.Get("size"))

		data := NewFileReader().ReadAll(file)
		fmt.Println("Input", string(data))
	}()
}
