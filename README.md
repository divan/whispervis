# P2P visualization

Web application for visualizing network graphs and p2p message propagation algorithms.
It provides an UI for choosing/generating different network graphs, requesting simulation results from the simulation backend, displaying statistics and animating message(s) propagation.

## Demo
[![Demo](https://img.youtube.com/vi/z2Zrfz6xxng/0.jpg)](https://www.youtube.com/watch?v=z2Zrfz6xxng)


## Usage

Just open `index.html` in the modern browser.

```
git clone git@github.com:status-im/whispervis.git
cd whisperviz/

# on MacOS
open index.html
```

## Development
This app is written in Go, using [Vecty](https://github.com/gopherjs/vecty) and [GopherJS](https://github.com/gopherjs/gopherjs). Vecty is a framework for building components on top of [GopherJS](https://github.com/gopherjs/gopherjs), a bit similar to React. Unlike React, it's written in a programming language instead of JavaScript, so the code is maintanable.

If you want to contribute to the development, you will need Go and GopherJS installed:

```
go get -u github.com/gopherjs/gopherjs
```

Then, install source code:
```
go get github.com/status-im/whispervis
cd $GOPATH/github.com/status-im/whispervis
```
after you made your changes, just run:

```
gopherjs build
```

and reopen `index.html`.

## Licence
MIT
