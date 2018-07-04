var graph = require('../index.js');

var ws = new WebSocket('ws://' + window.location.host + '/ws');

// request graphData and initial positions from websocket connection
ws.onopen = function (event) {
	ws.send('{"cmd": "init"}'); 
};

ws.onmessage = function (event) {
	let msg = JSON.parse(event.data);
	switch(msg.type) {
		case "graph":
			graph.setGraphData(msg.graph);
			break;
		case "propagation":
			graph.setPropagation(msg.propagation);
			break;
		case "positions":
			console.log("Updating positions...");
			graph.updatePositions(msg.positions);
			break;
	}
}

module.exports = { ws };
