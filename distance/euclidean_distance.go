package distance

import (
	"math"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/object"
)

// EuclideanDistance calculation method
type EuclideanDistance struct {
}

// GetDistance return euclidean distance of 2 objects
func (ed *EuclideanDistance) GetDistance(objectOrigin, objectDestination object.Object) float64 {
	coordOrigin, coordDestination := objectOrigin.GetCoordinate(), objectDestination.GetCoordinate()
	return ed.calcDistance(coordOrigin, coordDestination)
}

func (ed *EuclideanDistance) calcDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	distanceX := coordOrigin.Latitude - coordDestination.Latitude
	distanceY := coordOrigin.Longitude - coordDestination.Longitude
	distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
	return distance
}

// GetDistanceMatrix returns distance matrix
func (ed *EuclideanDistance) GetDistanceMatrix(objectList object.List) [][]float64 {
	coordList := objectList.GetCoordinateList()
	distanceMatrix := make([][]float64, len(coordList))
	for idx := range coordList {
		distanceList := make([]float64, len(coordList))
		distanceMatrix[idx] = distanceList
	}

	for idx1 := range coordList {
		coord1 := coordList[idx1]
		for idx2 := idx1 + 1; idx2 < len(coordList); idx2++ {
			coord2 := coordList[idx2]
			distance := ed.calcDistance(coord1, coord2)
			distanceMatrix[idx1][idx2] = distance
			distanceMatrix[idx2][idx1] = distance
		}
	}
	return distanceMatrix
}
