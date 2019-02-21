package mtree

type leafEntry struct {
	parent             node
	distanceFromParent float64
	object             Object
}

func (lfe *leafEntry) getRadius() float64 {
	return 0
}

func (lfe *leafEntry) getCentroidObject() Object {
	return lfe.object
}

func (lfe *leafEntry) getDistanceFromParent() float64 {
	return lfe.distanceFromParent
}

func (lfe *leafEntry) setDistanceFromParent(distance float64) {
	lfe.distanceFromParent = distance
}

func (lfe *leafEntry) getParent() node {
	return lfe.parent
}

func (lfe *leafEntry) setParent(parent node) {
	lfe.parent = parent
}
