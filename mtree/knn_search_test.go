package mtree_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mtree"
)

func TestKnnSearch(t *testing.T) {
	// Arrange
	tree := mtree.NewTree(3, 6, distCalc)
	queryObj := &object{
		id: "181881",
		x:  4,
		y:  5,
	}
	sortedObjectDistance := make([]float64, 0)

	objList := generateObjectList(1000000)
	for idx := range objList {
		obj := objList[idx]
		distance := distCalc.GetDistance(queryObj, obj)
		sortedObjectDistance = append(sortedObjectDistance, distance)
		tree.Insert(obj)
	}

	sort.Float64s(sortedObjectDistance)
	fmt.Println(sortedObjectDistance)

	// Act
	neighbours := tree.KnnSearch(queryObj, 5)

	// Assert
	for idx := range neighbours {
		neighbour := neighbours[idx]
		distance := neighbour.Distance
		expectedDistance := sortedObjectDistance[idx]
		assert.Equalf(t, expectedDistance, distance, "Error, expected distance of the %d-th neighbour = %f, got %f", idx, expectedDistance, distance)
	}
}
