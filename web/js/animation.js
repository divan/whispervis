var gradient = require('d3-scale-chromatic').interpolateCool;

function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}

let mat = new THREE.LineBasicMaterial({
	color: '#cc0000',
	transparent: true,
	opacity: 0.7
});

// Params
var restoreTimeout = 250; // ms
function updateRestoreTimeout(value) {
	console.log("Updating restore timeout to ", value);
	restoreTimeout = value;
}
var delayFactor = 3;     // multiplication factor for blink timeout
function updateDelayFactor(value) {
	delayFactor = value;
}

function blinkLink(links, indices) {
	indices.forEach(idx => {
		if (links.children[idx].material === mat) {
			return
		}
		let oldMat = links.children[idx].material;
		links.children[idx].material = mat;

		console.log("Blinking with restore timeout", restoreTimeout);
		setTimeout(function() {
			links.children[idx].material = oldMat;
		}, restoreTimeout);
	});
}

var nodeMaterials = {};

// blinkNodes updates nodes color 
function blinkNodes(nodes, indices) {
	// TODO: implement proper node blink
	/*
	Object.keys(nodeCounters).forEach(idx => {
		let c = nodeCounters[idx];
		let scale = c / maxCounter;
		let color = gradient(scale);
		if (nodeMaterials[color] === undefined) {
			nodeMaterials[color] = new THREE.MeshStandardMaterial({color: new THREE.Color(color)});
		}
		nodes.children[idx].material = nodeMaterials[color];
	});
	*/
}

function animatePropagation(nodes, links, plog) {
	maxCounter = 0;
	nodeCounters = {};
	plog.Timestamps.forEach((ts, idx) => {
		setTimeout(function() {
			blinkLink(links, plog.Indices[idx]);
		}, ts*delayFactor);
		setTimeout(function() {
			blinkNodes(nodes, plog.Nodes[idx]);
		}, ts*delayFactor);
	});
}

module.exports = { animatePropagation,
	restoreTimeout, updateRestoreTimeout,
	delayFactor, updateDelayFactor }
