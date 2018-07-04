var schemePaired = require('d3-scale-chromatic').schemePaired;
var tinyColor = require('tinycolor2');

const colorStr2Hex = str => isNaN(str) ? parseInt(tinyColor(str).toHex(), 16) : str;

function autoColorNodes(nodes) {
    const colors = schemePaired; // Paired color set from color brewer

    const uncoloredNodes = nodes.filter(node => !node.color);
    const nodeGroups = {};

    uncoloredNodes.forEach(node => { nodeGroups[node["group"]] = null });
    Object.keys(nodeGroups).forEach((group, idx) => { nodeGroups[group] = idx });

    uncoloredNodes.forEach(node => {
        node.color = colorStr2Hex(colors[nodeGroups[node["group"]] % colors.length]);
    });
}

module.exports = { colorStr2Hex, autoColorNodes };
