package distance

import (
	"math"

	"github.com/m-tree/coordinate"
	"github.com/m-tree/object"
)

// ManhattanDistance method to calculate manhattan distance of 2 objects
type ManhattanDistance struct {
}

// GetDistance return manhattan distanc of 2 objects
func (md *ManhattanDistance) GetDistance(objectOrigin, objectDestination object.Object) float64 {
	coordOrigin, coordDestination := objectOrigin.GetCoordinate(), objectDestination.GetCoordinate()
	return md.calcDistance(coordOrigin, coordDestination)
}

// GetDistanceMatrix returns distance matrix
func (md *ManhattanDistance) GetDistanceMatrix(objectList object.List) [][]float64 {
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
			distance := md.calcDistance(coord1, coord2)
			distanceMatrix[idx1][idx2] = distance
			distanceMatrix[idx2][idx1] = distance
		}
	}
	return distanceMatrix
}

func (md *ManhattanDistance) calcDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	distance := math.Abs(coordOrigin.Latitude-coordDestination.Latitude) + math.Abs(coordOrigin.Longitude-coordDestination.Longitude)
	return distance
}
