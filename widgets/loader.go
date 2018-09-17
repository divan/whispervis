package widgets

import (
	"fmt"
	"sync"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// Loader implements prograss-bar style loader widget.
type Loader struct {
	vecty.Core

	mx      sync.RWMutex
	steps   int
	current int
}

// Render implements Component interface for Loader.
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

// NewLoader creates new Loader.
func NewLoader() *Loader {
	return &Loader{}
}

// Inc increases current value of loader by one.
func (l *Loader) Inc() {
	l.mx.Lock()
	if l.current < l.steps {
		l.current++
	}
	l.mx.Unlock()
}

// Steps returns the total number of steps.
func (l *Loader) Steps() int {
	l.mx.RLock()
	defer l.mx.RUnlock()
	return l.steps
}

// Reset resets the current state of Loader.
func (l *Loader) Reset() {
	l.mx.Lock()
	l.current = 0
	l.mx.Unlock()
}

// SetSteps updates the number of steps for loader.
func (l *Loader) SetSteps(steps int) {
	l.mx.Lock()
	l.steps = steps
	l.mx.Unlock()
}

// Progress reports Loader's progress in percentage.
func (l *Loader) Progress() float64 {
	l.mx.RLock()
	defer l.mx.RUnlock()
	return 100 * float64(l.current) / float64(l.steps)
}

// text returns formatted string to display.
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
