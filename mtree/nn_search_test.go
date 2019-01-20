package mtree_test

import (
	"testing"

	"github.com/m-tree/mtree"
	"github.com/stretchr/testify/assert"

	"github.com/m-tree/coordinate"
	"github.com/m-tree/distance"
)

var nnTestCase = struct {
	coordList      []*coordinate.Coordinate
	queryCoord     *coordinate.Coordinate
	maxEntry       int
	nnDistanceList []float64
}{
	coordList: []*coordinate.Coordinate{
		&coordinate.Coordinate{1, 4},
		&coordinate.Coordinate{1, 6},
		&coordinate.Coordinate{2, 2},
		&coordinate.Coordinate{4, 1},
		&coordinate.Coordinate{4, 4},
		&coordinate.Coordinate{5, 3},
		&coordinate.Coordinate{6, 2},
		&coordinate.Coordinate{6, 5},
		&coordinate.Coordinate{5, 6},
		&coordinate.Coordinate{7, 4},
		&coordinate.Coordinate{8, 1},
		&coordinate.Coordinate{8, 6},
		&coordinate.Coordinate{9, 2},
		&coordinate.Coordinate{9, 3},
	},
	maxEntry:       5,
	queryCoord:     &coordinate.Coordinate{3, 3},
	nnDistanceList: []float64{2, 2, 2, 3, 3, 4, 5, 5, 5, 5, 6, 7, 7, 8},
}

func TestNN(t *testing.T) {
	// Arrange
	distCalc := &distance.ManhattanDistance{}
	splitMecha := &mtree.SplitMST{
		DistCalc: distCalc,
		MaxEntry: nnTestCase.maxEntry,
	}
	tree := mtree.NewTree(testCase.maxEntry, splitMecha, distCalc)
	coordList := nnTestCase.coordList
	for idx := range coordList {
		coord := coordList[idx]
		tree.Insert(coord, idx)
	}

	// Act
	queryCoord := nnTestCase.queryCoord
	nnList := tree.KnnSearch(tree.Root, queryCoord, 14)
	expectedDistanceList := nnTestCase.nnDistanceList

	// Assert
	t.Run("Test the distance order in nearest neighbour search", func(t *testing.T) {
		for idx := range nnList {
			distance := nnList[idx].Distance
			expectedDistance := expectedDistanceList[idx]
			assert.Equalf(t, expectedDistance, distance, "Error! expected distance for the %d-th neighbour = %f, got %f", idx+1, expectedDistance, distance)
		}
	})
}
