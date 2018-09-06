package widgets

import (
	"fmt"
	"sync"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Loader struct {
	vecty.Core

	mx      sync.RWMutex
	steps   int
	current int
}

func (l *Loader) Render() vecty.ComponentOrHTML {
	text := l.text()
	return elem.Div(
		vecty.Markup(
			vecty.Style("text-align", "center"),
			vecty.Style("position", "relative"),
			vecty.Style("top", "50%"),
		),
		elem.Heading1(
			vecty.Text(text),
		),
	)
}

func NewLoader(steps int) *Loader {
	return &Loader{
		steps: steps,
	}
}

func (l *Loader) Inc() {
	l.mx.Lock()
	l.current++
	l.mx.Unlock()
}

func (l *Loader) Steps() int {
	l.mx.RLock()
	defer l.mx.RUnlock()
	return l.steps
}

func (l *Loader) Reset() {
	l.mx.Lock()
	l.current = 0
	l.mx.Unlock()
}

// Progress reports loader's progress in percentage.
func (l *Loader) Progress() float64 {
	l.mx.RLock()
	defer l.mx.RUnlock()
	return 100 * float64(l.current) / float64(l.steps)
}

// Text returns formatted string to display.
func (l *Loader) text() string {
	var text string
	progress := l.Progress()
	if progress > 99.9 {
		text = "Completed"
	} else {
		text = fmt.Sprintf("Loading %.0f%%...", progress)
	}
	return text
}
