// accessorFn
function accessorFn(p) {
	if (p instanceof Function) {
		return p                     // fn
	}

    if (typeof p === 'string') {
        return function(obj) {
			return obj[p];     // property name
		}
	}
    
	return function (obj){
		return p;         // constant
	}
}

module.exports = accessorFn;
