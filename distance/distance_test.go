package distance_test

import (
	"math"
	"testing"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/distance"
	"github.com/vroup/mo-iwd-sa/mtree"
)

type DistCalculator interface {
	GetDistance(mtree.Object, mtree.Object) float64
}

type TestCase struct {
	lat1, lon1         float64
	lat2, lon2         float64
	distance           float64
	tolerance          float64
	distanceCalculator DistCalculator
}

var testCases = map[string]TestCase{
	"haversine": {-84.412977, 39.152501, -84.412946, 39.152505, 2.7098232942902385, 2, &distance.HaversineDistance{}},
	"euclidean": {-7, -4, 17, 6.5, 26.196374, 0.01, &distance.EuclideanDistance{}},
	"manhattan": {5, 5, 12, 13, 15, 0.01, &distance.ManhattanDistance{}},
}

func TestDistance(t *testing.T) {
	t.Run("haversine", func(t *testing.T) {
		testDistance("haversine", testCases["haversine"], t)
	})
	t.Run("euclidean", func(t *testing.T) {
		testDistance("euclidean", testCases["euclidean"], t)
	})
	t.Run("manhattan", func(t *testing.T) {
		testDistance("manhattan", testCases["manhattan"], t)
	})

}

func testDistance(name string, tc TestCase, t *testing.T) {
	// Arange
	distCalc := tc.distanceCalculator
	coord1 := &coordinate.Coordinate{
		Latitude:  tc.lat1,
		Longitude: tc.lon1,
	}
	coord2 := &coordinate.Coordinate{
		Latitude:  tc.lat2,
		Longitude: tc.lon2,
	}
	distExpected := tc.distance

	// Act
	distResult := distCalc.GetDistance(coord1, coord2)
	diff := math.Abs(distResult - distExpected)

	// Assert
	if diff > tc.tolerance {
		t.Errorf("Error! The difference (error) is too big, error = %f, expected error <= %f", diff, tc.tolerance)
	}
}
