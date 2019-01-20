package object

import "github.com/m-tree/coordinate"

// Object is temporary interface, as the object entry of m-tree
type Object interface {
	GetCoordinate() *coordinate.Coordinate
}

// List is just list of object
type List interface {
	GetCoordinateList() []*coordinate.Coordinate
	Len() int
}
