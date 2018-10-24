package main

import (
	"time"
)

// DefaultRenderThrottleDecay defines a decay period after which throttling should reenabled.
const DefaultRenderThrottleDecay = 5 // sec

// RenderThrottler is a helper to enable rendering only when it's really need to.
// Rendering can be enabled/requested externally, but will be turned off/slowed down
// after some period of time.
// TODO(divan): also, instead of disabling, might be a good idea to just decrease FPS. Explore this.
type RenderThrottler struct {
	needRendering bool
	lastUpdate    int64
	decay         int64
}

// NewRenderThrottler returns a new render throttler.
func NewRenderThrottler() *RenderThrottler {
	r := &RenderThrottler{
		needRendering: true,
		lastUpdate:    time.Now().Unix(),
		decay:         DefaultRenderThrottleDecay,
	}

	return r
}

// EnableRendering enables next frames to render.
func (r *RenderThrottler) EnableRendering() {
	r.needRendering = true
	r.lastUpdate = time.Now().Unix()
}

// DisableRendering disables next frames from rendering.
func (r *RenderThrottler) DisableRendering() {
	r.needRendering = false
}

// ReenableIfNeeded checks if sufficient time has passed since throttling has been
// disabled, and enables throttling back if so.
func (r *RenderThrottler) ReenableIfNeeded() {
	now := time.Now().Unix()
	if r.lastUpdate+r.decay < now {
		r.DisableRendering()
	}
}

// NeedRendering returns true if next frame should be rendered.
func (r *RenderThrottler) NeedRendering() bool {
	return r.needRendering
}
