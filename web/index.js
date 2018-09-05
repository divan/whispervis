var { colorStr2Hex, autoColorNodes } = require('./js/colors.js');
require('./js/keys.js');
var accessorFn = require('./js/shitty_hacks.js');
var { animatePropagation, restoreTimeout, updateRestoreTimeout, delayFactor, updateDelayFactor } = require('./js/animation.js');
var { NewEthereumGeometry } = require('./js/ethereum.js');

var Stats = require('stats-js');
const dat = require('dat.gui');


// WebGL
let canvas = document.getElementById("preview");
var renderer = new THREE.WebGLRenderer({ canvas: canvas });

var graphData, plog;
var positions = Array();

function setGraphData(data) {
	graphData = data;

	initGraph();
}

function setPropagation(plogData) {
	plog = plogData;
	animatePropagation(nodesGroup, linksGroup, plogData);
}

function updatePositions(data) {
	positions = [];
	Object.keys(data).forEach((k) => {
		let v = data[k];
		positions[k] = {
			x: v.X,
			y: v.Y,
			z: v.Z,
		};
	})
	console.log("Positions", positions);

	redrawGraph();
}

module.exports = { updatePositions, setGraphData, setPropagation };

// Setup scene
const scene = new THREE.Scene();
scene.background = new THREE.Color(0x000011);

// Add lights
scene.add(new THREE.AmbientLight(0xbbbbbb));
scene.add(new THREE.DirectionalLight(0xffffff, 0.6));

var linksGroup = new THREE.Group();
scene.add(linksGroup);
var nodesGroup = new THREE.Group();
scene.add(nodesGroup);

// Setup camera
var camera = new THREE.PerspectiveCamera();
camera.far = 20000;

var tbControls = new THREE.TrackballControls(camera, renderer.domElement);
var flyControls = new THREE.FlyControls(camera, renderer.domElement);

var animate = function () {
	// frame cycle
	tbControls.update();
	flyControls.update(1);

	renderer.render(scene, camera);
	stats.update();
	requestAnimationFrame( animate );
};

var width = window.innerWidth * 80 / 100 - 20;
var height = window.innerHeight - 20;
var nodeRelSize = 1;
var nodeResolution = 8;

// Stats
var stats = new Stats();
document.body.appendChild( stats.domElement );
stats.domElement.style.position = 'absolute';
stats.domElement.style.right = '15px';
stats.domElement.style.bottom = '20px';

// Dat GUI
const gui = new dat.GUI();
var f1 = gui.addFolder('Animation');
var restoreCtl = f1.add({ restoreTimeout: restoreTimeout }, 'restoreTimeout').name('Restore timeout');
restoreCtl.onFinishChange(function(value) {
	updateRestoreTimeout(value);
});
var factorCtl = f1.add({ delayFactor: delayFactor }, 'delayFactor').name('Delay factor');
factorCtl.onFinishChange(function(value) {
	updateDelayFactor(value);
});

var initGraph = function () {
	resizeCanvas();

	// parse links
	graphData.links.forEach(link => {
		link.source = link["source"];
		link.target = link["target"];
	});

	// Add WebGL objects
	// Clear the place
	while (nodesGroup.children.length) {
		nodesGroup.remove(nodesGroup.children[0])
	} 
	while (linksGroup.children.length) {
		linksGroup.remove(linksGroup.children[0])
	}

	// Render nodes
	const nameAccessor = accessorFn("name");
	const valAccessor = accessorFn("weight");
	const colorAccessor = accessorFn("color");
	let sphereGeometries = {}; // indexed by node value
	let sphereMaterials = {}; // indexed by color

	autoColorNodes(graphData.nodes);
	graphData.nodes.forEach((node, idx) => {
		let val = valAccessor(node) || 1;
		if (!sphereGeometries.hasOwnProperty(val)) {
			sphereGeometries[val] = NewEthereumGeometry(val);
		}

		const color = colorAccessor(node);
		if (!sphereMaterials.hasOwnProperty(color)) {
			sphereMaterials[color] = new THREE.MeshStandardMaterial({
				color: colorStr2Hex(color || '#00ff00'),
				transparent: false,
				opacity: 0.75
			});
		}

		const sphere = new THREE.Mesh(sphereGeometries[val], sphereMaterials[color]);

		sphere.name = nameAccessor(node); // Add label
		sphere.__data = node; // Attach node data

		nodesGroup.add(node.__sphere = sphere);
		if (positions[node.id] !== undefined) {
			sphere.position.set(positions[node.id].x, positions[node.id].y, positions[node.id].z);
		}
	});

	const linkColorAccessor = accessorFn("color");
	let lineMaterials = {}; // indexed by color
	console.log("Adding links", graphData.links.lengh);
	graphData.links.forEach(link => {
		const color = linkColorAccessor(link);
		if (!lineMaterials.hasOwnProperty(color)) {
			lineMaterials[color] = new THREE.LineBasicMaterial({
				color: /*colorStr2Hex(color || '#f0f0f0')*/ '#f0f0f0',
				transparent: true,
				opacity: 0.4,
			});
		}

		const geometry = new THREE.BufferGeometry();
		geometry.addAttribute('position', new THREE.BufferAttribute(new Float32Array(2 * 3), 3));
		const lineMaterial = lineMaterials[color];
		const line = new THREE.Line(geometry, lineMaterial);

		line.renderOrder = 10; // Prevent visual glitches of dark lines on top of spheres by rendering them last

		linksGroup.add(link.__line = line);
	});

	// correct camera position
	if (camera.position.x === 0 && camera.position.y === 0) {
		// If camera still in default position (not user modified)
		camera.lookAt(nodesGroup.position);
		camera.position.z = Math.cbrt(graphData.nodes.length) * 50;
	}

	function resizeCanvas() {
		if (width && height) {
			renderer.setSize(width, height);
			camera.aspect = width/height;
			camera.updateProjectionMatrix();
		}
	}
};

var redrawGraph = function () {
	graphData.nodes.forEach((node, idx) => {
		const sphere = node.__sphere;
		if (!sphere) return;

		sphere.position.x = positions[node.id].x;
		sphere.position.y = positions[node.id].y || 0;
		sphere.position.z = positions[node.id].z || 0;
	});


	graphData.links.forEach(link => {
		const line = link.__line;
		if (!line) return;

		linePos = line.geometry.attributes.position;

		let start = link.source;
		let end = link.target;

		linePos.array[0] = positions[start].x;
		linePos.array[1] = positions[start].y || 0;
		linePos.array[2] = positions[start].z || 0;
		linePos.array[3] = positions[end].x;
		linePos.array[4] = positions[end].y || 0;
		linePos.array[5] = positions[end].z || 0;

		linePos.needsUpdate = true;
		line.geometry.computeBoundingSphere();
	});
};

// replay restarts propagation animation.
function replay() {
	console.log(linksGroup, plog);
	animatePropagation(nodesGroup, linksGroup, plog);
}
// js functions after browserify cannot be accessed from html,
// so instead of using onclick="replay()" we need to attach listener
// here.
// Did I already say that whole frontend ecosystem is a one giant
// museum of hacks for hacks on top of hacks?
var replayButton = document.getElementById('replayButton');
replayButton.addEventListener('click', replay);

animate();

// Handle mouse hover
var INTERSECTED;

function onMouseMove( event ) {
	let canvasBounds = renderer.context.canvas.getBoundingClientRect();
    mouse.x = ( ( event.clientX - canvasBounds.left ) / ( canvasBounds.right - canvasBounds.left ) ) * 2 - 1;
	mouse.y = - ( ( event.clientY - canvasBounds.top ) / ( canvasBounds.bottom - canvasBounds.top) ) * 2 + 1;

	raycaster.setFromCamera( mouse, camera );
	var intersects = raycaster.intersectObjects( scene.children, true );

    let nodeInfo = document.getElementById('nodeInfo');
    let nodeID = document.getElementById('selectedNodeID');
	if (intersects.length > 0) {
		// if the closest object intersected is not the currently stored intersection object
		if (intersects[0].object != INTERSECTED) {
			// restore previous intersection object (if it exists) to its original color
			if (INTERSECTED)
				INTERSECTED.material.color.setHex(INTERSECTED.currentHex);
			// store reference to closest object as current intersection object
			
			// find the object representing node (has __data.id field)
			let obj = intersects.filter(x => x.object.__data !== undefined);
			if (obj.length == 0) { return }
			INTERSECTED = obj[0].object;
			if (INTERSECTED.__data !== undefined) {
				let id = INTERSECTED.__data.id;
				nodeInfo.hidden = false;
				nodeID.innerHTML = id;
				let stats = current();
				let nodeStats = stats.Nodes.filter(x => x.ID == id)[0];
				let count = nodeStats.ClientsNum + nodeStats.PeersNum;
			}
			// store color of closest object (for later restoration)
			INTERSECTED.currentHex = INTERSECTED.material.color.getHex();
			// set a new color for closest object
			INTERSECTED.material.color.setHex(0xffff00);
		}
	} else {
		// restore previous intersection object (if it exists) to its original color
		if (INTERSECTED)
		  INTERSECTED.material.color.setHex(INTERSECTED.currentHex);
		// remove previous intersection object reference
		//     by setting current intersection object to "nothing"
		INTERSECTED = null;
		nodeInfo.hidden = true;
	}
}

canvas.addEventListener( 'mousemove', onMouseMove, false );
