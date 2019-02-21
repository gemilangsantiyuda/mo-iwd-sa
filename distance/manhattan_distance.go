package distance

import (
	"math"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/mtree"
	iwdObj "github.com/vroup/mo-iwd-sa/object"
)

// ManhattanDistance method to calculate manhattan distance of 2 objects
type ManhattanDistance struct {
}

// GetDistance return manhattan distanc of 2 objects
func (md *ManhattanDistance) GetDistance(objectOrigin, objectDestination mtree.Object) float64 {
	rOrigin := objectOrigin.(iwdObj.Object)
	rDest := objectDestination.(iwdObj.Object)
	coordOrigin, coordDestination := rOrigin.GetCoordinate(), rDest.GetCoordinate()
	return md.calcDistance(coordOrigin, coordDestination)
}
func (md *ManhattanDistance) calcDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	return math.Abs(coordOrigin.Latitude-coordDestination.Latitude) + math.Abs(coordOrigin.Longitude-coordDestination.Longitude)
}
