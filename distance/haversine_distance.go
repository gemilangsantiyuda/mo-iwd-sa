package distance

import (
	"math"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/mtree"
	iwdObj "github.com/vroup/mo-iwd-sa/object"
)

// HaversineDistance repo for haversine distance calculator
type HaversineDistance struct {
}

// GetDistance return the haversine distance of the 2 objects based on their latitude longitude location
func (hd *HaversineDistance) GetDistance(objectOrigin, objectDestination mtree.Object) float64 {
	rOrigin := objectOrigin.(iwdObj.Object)
	rDest := objectDestination.(iwdObj.Object)
	coordOrigin, coordDestination := rOrigin.GetCoordinate(), rDest.GetCoordinate()
	distance := hd.calculateCoordHaversineDistance(coordOrigin, coordDestination)
	return distance
}

func (hd *HaversineDistance) calculateCoordHaversineDistance(coordOrigin, coordDestination *coordinate.Coordinate) float64 {
	DY := math.Abs(coordOrigin.Latitude-coordDestination.Latitude) / 180 * math.Pi
	DX := math.Abs(coordOrigin.Longitude-coordDestination.Longitude) / 180 * math.Pi
	Y1 := coordOrigin.Latitude / 180 * math.Pi
	Y2 := coordDestination.Latitude / 180 * math.Pi
	R := 6372800.00000000 // Approximation of earth radius in meter
	a := math.Sin(DY/2)*math.Sin(DY/2) + math.Cos(Y1)*math.Cos(Y2)*math.Sin(DX/2)*math.Sin(DX/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c * 5. / 4.
}
