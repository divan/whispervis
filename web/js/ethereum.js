function NewEthereumGeometry(scale) {
	let geom = new THREE.Geometry();
	geom.vertices.push(
		new THREE.Vector3(  scale*1,  0,  0 ),
		new THREE.Vector3( -scale*1,  0,  0 ),
		new THREE.Vector3(  0,  scale*1.5,  0 ),
		new THREE.Vector3(  0, scale*-1.5,  0 ),
		new THREE.Vector3(  0,  0,  scale*1 ),
		new THREE.Vector3(  0,  0, -scale*1 )
	);
	geom.faces.push(
		new THREE.Face3( 0, 2, 4 ),
		new THREE.Face3( 0, 4, 3 ),
		new THREE.Face3( 0, 3, 5 ),
		new THREE.Face3( 0, 5, 2 ),
		new THREE.Face3( 1, 2, 5 ),
		new THREE.Face3( 1, 5, 3 ),
		new THREE.Face3( 1, 3, 4 ),
		new THREE.Face3( 1, 4, 2 )
	);
	geom.computeBoundingSphere();
	geom.computeFaceNormals();
	return geom;
}

module.exports = { NewEthereumGeometry };
