package mtree

import (
	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/object"
)

// LeafEntry stores the spatial object
type LeafEntry struct {
	Object             object.Object
	ObjectID           string
	Parent             *Node
	DistanceFromParent float64
}

// GetCoordinate return the spatial object lat lon coordinate
func (le *LeafEntry) GetCoordinate() *coordinate.Coordinate {
	return le.Object.GetCoordinate()
}

// GetParent return this entry parent
func (le *LeafEntry) GetParent() *Node {
	return le.Parent
}

// SetParent set this entry's parent, and update its distance from parent
func (le *LeafEntry) SetParent(newParent *Node) {
	le.Parent = newParent
}

// GetRadius return 0 because spatial object has no radius
func (le *LeafEntry) GetRadius() float64 {
	return 0.
}

// GetDistanceFromParent return this entry's distance from its parent
func (le *LeafEntry) GetDistanceFromParent() float64 {
	return le.DistanceFromParent
}

// SetDistanceFromParent set this leaf entry's distance to its parent node
func (le *LeafEntry) SetDistanceFromParent(dist float64) {
	le.DistanceFromParent = dist
}
