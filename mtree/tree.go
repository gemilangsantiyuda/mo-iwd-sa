package mtree

import (
	"github.com/m-tree/object"
)

// SplitMechanism to make split mechanism open for other option
type SplitMechanism interface {
	Split(*Node, Entry) *Node
}

// DistanceCalculator , interfacing it so that the testing can simply use euclidean to simplify test
type DistanceCalculator interface {
	GetDistance(object.Object, object.Object) float64
	GetDistanceMatrix(objectList object.List) [][]float64
}

// Tree it's the tree, it has root and maximum entry of each node
type Tree struct {
	Root           *Node
	MaxEntry       int
	SplitMechanism SplitMechanism
	DistCalc       DistanceCalculator
}
