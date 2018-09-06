package widgets

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Loader struct {
	vecty.Core

	steps   int
	current int
}

func (l *Loader) Render() vecty.ComponentOrHTML {
	var text string
	progress := l.Progress()
	if progress > 99.9 {
		text = "Completed"
	} else {
		text = fmt.Sprintf("Loading %.1f%%...", progress)
		fmt.Println("Loader.Render()", text, progress)
	}
	return elem.Div(
		elem.Heading1(vecty.Text(text)),
	)
}

func NewLoader(steps int) *Loader {
	return &Loader{
		steps: steps,
	}
}

func (l *Loader) Inc() {
	l.current++
	vecty.Rerender(l)
}

// Progress reports loader's progress in percentage.
func (l *Loader) Progress() float64 {
	fmt.Println("progress", 100*float64(l.current)/float64(l.steps))
	return 100 * float64(l.current) / float64(l.steps)
}
