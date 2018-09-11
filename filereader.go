package main

import "github.com/gopherjs/gopherjs/js"

// FileReader implements FileReader API for js.Object.
type FileReader struct {
	*js.Object // FileReader object
}

// NewFileReader inits new FileReader.
func NewFileReader() *FileReader {
	return &FileReader{
		Object: js.Global.Get("FileReader").New(),
	}
}

// ReadAll reads all data from blob using FileReader API.
func (fr *FileReader) ReadAll(blob *js.Object) []byte {
	ch := make(chan []byte)
	fr.Set("onload", func() {
		ch <- js.Global.Get("Uint8Array").New(fr.Get("result")).Interface().([]byte)
	})
	fr.Call("readAsArrayBuffer", blob)
	return <-ch
}
