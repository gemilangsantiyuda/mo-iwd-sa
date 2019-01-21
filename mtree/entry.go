package mtree

import "github.com/vroup/mo-iwd-sa/coordinate"

// Entry interface
type Entry interface {
	GetParent() *Node
	SetParent(*Node)
	GetCoordinate() *coordinate.Coordinate
	GetRadius() float64
	GetDistanceFromParent() float64
	SetDistanceFromParent(float64)
}

// EntryList i don't know sounds silly but migh work on passing the getDistanceMatrix function to distance instead
type EntryList []Entry

// GetCoordinateList gets list of coordinates from each entry
func (el EntryList) GetCoordinateList() []*coordinate.Coordinate {
	var coordList []*coordinate.Coordinate
	for idx := range el {
		entry := el[idx]
		coord := entry.GetCoordinate()
		coordList = append(coordList, coord)
	}
	return coordList
}

// Len return the length of entrylist
func (el EntryList) Len() int {
	return len(el)
}
