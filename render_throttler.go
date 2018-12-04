package main

import (
	"time"

	"github.com/divan/whispervis/storage"
)

// DefaultRenderThrottleDecay defines a decay period after which throttling should reenabled.
const (
	DefaultRenderThrottleDecay = 5 // sec
)

// RenderThrottler is a helper to enable rendering only when it's really need to.
// Rendering can be enabled/requested externally, but will be turned off/slowed down
// after some period of time.
// TODO(divan): also, instead of disabling, might be a good idea to just decrease FPS. Explore this.
type RenderThrottler struct {
	isActive bool // switch on/off

	needRendering bool // what throttler thinks about next frame render pass
	lastUpdate    int64
	decay         int64
}

// NewRenderThrottler returns a new render throttler.
func NewRenderThrottler() *RenderThrottler {
	r := &RenderThrottler{
		isActive:      storage.RT(),
		needRendering: true,
		lastUpdate:    time.Now().Unix(),
		decay:         DefaultRenderThrottleDecay,
	}

	return r
}

// Toggle switches render throttler state.
func (r *RenderThrottler) Toggle() {
	r.isActive = !r.isActive
	storage.SetRT(r.isActive) // TODO: this is bad. global local storage
}

// IsActive returns state of render throttler.
func (r *RenderThrottler) IsActive() bool {
	return r.isActive
}

// EnableRendering enables next frames to render.
func (r *RenderThrottler) EnableRendering() {
	if !r.isActive {
		return
	}

	r.needRendering = true
	r.lastUpdate = time.Now().Unix()
}

// DisableRendering disables next frames from rendering.
func (r *RenderThrottler) DisableRendering() {
	if !r.isActive {
		return
	}
	r.needRendering = false
}

// ReenableIfNeeded checks if sufficient time has passed since throttling has been
// disabled, and enables throttling back if so.
func (r *RenderThrottler) ReenableIfNeeded() {
	if !r.isActive {
		return
	}
	now := time.Now().Unix()
	if r.lastUpdate+r.decay < now {
		r.DisableRendering()
	}
}

// NeedRendering returns true if next frame should be rendered.
func (r *RenderThrottler) NeedRendering() bool {
	if !r.isActive {
		return true // always render when throttler is disabled
	}
	return r.needRendering
}
