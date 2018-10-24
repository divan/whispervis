package storage

import (
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

// TODO: add forces editor values

// this app can be served from "file://", so let's not
// mix with other local apps
const prefix = "whisperviz_"

// Get returns value by key from localStorage.
func Get(key string) string {
	return js.Global.Get("localStorage").Get(prefix + key).String()
}

// Set stores value by key in localStorage.
func Set(key, value string) {
	js.Global.Get("localStorage").Set(prefix+key, value)
}

// SetNetwork saves network preset value to local storage.
func SetNetwork(name string) {
	Set("network", name)
}

func Network() string {
	net := Get("network")
	if net == "" {
		net = "3dgrid125.json"
	}
	return net
}

// SetFPS saves FPS value to local storage.
func SetFPS(fps int) {
	Set("fps", strconv.Itoa(fps))
}

func FPS() int {
	var fps int = 60

	str := Get("fps")
	if str != "" {
		value, err := strconv.Atoi(str)
		if err == nil {
			fps = value
		}
	}
	return fps
}

// SetRT saves Render Throttler value to local storage.
func SetRT(val bool) {
	str := "1"
	if !val {
		str = "0"
	}
	Set("render_throttler", str)
}

// RT returns stored value for Render Throttler option.
func RT() bool {
	var rt bool = true

	str := Get("render_throttler")
	if str == "0" {
		rt = false
	}
	return rt
}
