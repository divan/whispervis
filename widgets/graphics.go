package widgets

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/status-im/whispervis/storage"
)

// Graphics represents graphics configuration widget.
type Graphics struct {
	vecty.Core

	rtSwitch    *Switch
	fpsRadio    *RadioGroup
	blinkRange  *Range
	collapsable *Collapsable

	conf SceneConfigurator
}

// NewGraphics creates a new Graphics widget. It needs to have
// access to scene configuration, as it configures mostly things from it.
func NewGraphics(conf SceneConfigurator) *Graphics {
	g := &Graphics{
		conf: conf,
	}
	g.rtSwitch = NewSwitch("Render throttler", storage.RT(), conf.ToggleRenderThrottler)
	g.fpsRadio = NewRadioGroup("FPS", storage.FPS(), conf.ChangeFPS, []int{60, 30, 20, 15})
	g.blinkRange = NewRange("Blink:", BlinkDescription, storage.BlinkTime(), 10, 1500, conf.ChangeBlinkTime)
	g.collapsable = NewCollapsable("Graphics:", false,
		g.applyButton,
		g.fpsRadio,
		g.rtSwitch,
		g.blinkRange,
	)
	return g
}

// Render implements vecty's Component interface for Graphics.
func (g *Graphics) Render() vecty.ComponentOrHTML {
	return Widget(
		g.collapsable,
	)
}

// FIXME: it exists due to limitations of collapsible.
func (g *Graphics) applyButton() vecty.ComponentOrHTML {
	return elem.Span()
}

const (
	BlinkDescription = "How many milliseconds should blink last. You may want to change this value to make animations more interesting depending on the length of the total simulation"
)
