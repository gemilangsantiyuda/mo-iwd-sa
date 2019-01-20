package distance

import (
	"math"

	"github.com/m-tree/coordinate"
	"github.com/m-tree/object"
)

// HaversineDistance repo for haversine distance calculator
type HaversineDistance struct {
}

// GetDistance return the haversine distance of the 2 objects based on their latitude longitude location
func (hd *HaversineDistance) GetDistance(objectOrigin, objectDestination object.Object) float64 {
	coordOrigin := objectOrigin.GetCoordinate()
	coordDestination := objectDestination.GetCoordinate()
	haversineDistance := hd.calculateCoordHaversineDistance(coordOrigin, coordDestination)
	return haversineDistance
}

func (hd *HaversineDistance) calculateCoordHaversineDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	DY := math.Abs(coordOrigin.Latitude-coordDestination.Latitude) / 180 * math.Pi
	DX := math.Abs(coordOrigin.Longitude-coordDestination.Longitude) / 180 * math.Pi
	Y1 := coordOrigin.Latitude / 180 * math.Pi
	Y2 := coordDestination.Latitude / 180 * math.Pi
	R := 6372800.00000000 // Approximation of earth radius in meter
	a := math.Sin(DY/2)*math.Sin(DY/2) + math.Cos(Y1)*math.Cos(Y2)*math.Sin(DX/2)*math.Sin(DX/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

// GetDistanceMatrix returns distance matrix
func (hd *HaversineDistance) GetDistanceMatrix(objectList object.List) [][]float64 {
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
			distance := hd.calculateCoordHaversineDistance(coord1, coord2)
			distanceMatrix[idx1][idx2] = distance
			distanceMatrix[idx2][idx1] = distance
		}
	}
	return distanceMatrix
}
