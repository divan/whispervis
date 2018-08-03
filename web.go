//go:generate browserify web/index.js web/js/ws.js -o web/bundle.js
//go:generate $GOPATH/bin/go-bindata -nocompress=false -nomemcopy=true -prefix=web -o "assets.go" web/index.html web/bundle.js web/node_modules/three/build/three.min.js web/js/controls web/css/...
package main

import (
	"fmt"
	"github.com/elazarl/go-bindata-assetfs"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func startWeb(ws *WSServer, port string) {
	go func() {
		// Handle static files
		fs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: ""}
		http.Handle("/", noCacheMiddleware(http.FileServer(fs)))
		http.HandleFunc("/ws", ws.Handle)
		log.Fatal(http.ListenAndServe(port, nil))
	}()
	url := "http://localhost" + port
	fmt.Println("Please go to this url:", url)
	/*
		time.Sleep(1 * time.Second)
		startBrowser(url)
	*/
}

// startBrowser tries to open the URL in a browser
// and reports whether it succeeds.
//
// Orig. code: golang.org/x/tools/cmd/cover/html.go
func startBrowser(url string) error {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	fmt.Println("If browser window didn't appear, please go to this url:", url)
	return cmd.Start()
}

func noCacheMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=0, no-cache")
		h.ServeHTTP(w, r)
	})
}
