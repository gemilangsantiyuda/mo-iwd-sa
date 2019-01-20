package coordinate_test

import (
	"testing"

	"github.com/m-tree/coordinate"
	"github.com/stretchr/testify/assert"
)

var coordTests = []struct {
	lat, lon float64
}{
	{1, 1},
	{-0.212, -04444},
	{0, 0},
}

func TestGetCoordinate(t *testing.T) {
	// Arrange
	var coordList []*coordinate.Coordinate
	for _, tc := range coordTests {
		coord := &coordinate.Coordinate{
			Latitude:  tc.lat,
			Longitude: tc.lon,
		}
		coordList = append(coordList, coord)
	}

	// Act
	var coordResultList []*coordinate.Coordinate
	for idx := range coordList {
		coord := coordList[idx].GetCoordinate()
		coordResultList = append(coordResultList, coord)
	}

	// Assert
	for idx := range coordTests {
		coordResult := coordResultList[idx]
		coordExpected := coordTests[idx]
		assert.Equal(t, coordResult.Latitude, coordExpected.lat, "Latitude must be equal")
		assert.Equal(t, coordResult.Longitude, coordExpected.lon, "Longitude must be equal")
	}
}
