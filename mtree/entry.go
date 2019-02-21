package mtree

type entry interface {
	getParent() node
	getRadius() float64
	getDistanceFromParent() float64
	getCentroidObject() Object
	setParent(node)
	setDistanceFromParent(float64)
}
