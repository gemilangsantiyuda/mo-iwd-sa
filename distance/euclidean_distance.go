package distance

import (
	"math"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/mtree"
	iwdObj "github.com/vroup/mo-iwd-sa/object"
)

// EuclideanDistance calculation method
type EuclideanDistance struct {
}

// GetDistance return euclidean distance of 2 objects
func (ed *EuclideanDistance) GetDistance(objectOrigin, objectDestination mtree.Object) float64 {
	rOrigin := objectOrigin.(iwdObj.Object)
	rDest := objectDestination.(iwdObj.Object)
	coordOrigin, coordDestination := rOrigin.GetCoordinate(), rDest.GetCoordinate()
	distance := ed.calcDistance(coordOrigin, coordDestination)
	return distance
}

func (ed *EuclideanDistance) calcDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	distanceX := coordOrigin.Latitude - coordDestination.Latitude
	distanceY := coordOrigin.Longitude - coordDestination.Longitude
	distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
	return distance
}
