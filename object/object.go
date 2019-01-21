package object

import "github.com/vroup/mo-iwd-sa/coordinate"

// Object is temporary interface, as the object entry of m-tree
type Object interface {
	GetCoordinate() *coordinate.Coordinate
}

// List is just list of object
type List interface {
	GetCoordinateList() []*coordinate.Coordinate
	Len() int
}
