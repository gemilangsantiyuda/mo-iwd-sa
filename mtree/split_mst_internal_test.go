package mtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/m-tree/coordinate"
	"github.com/m-tree/distance"
)

// var distanceMatrixTestCase = struct {
// 	distCalc               DistanceCalculator
// 	coordList              []*coordinate.Coordinate
// 	expectedDistanceMatrix [][]float64
// }{
// 	distCalc: &distance.ManhattanDistance{},
// 	coordList: []*coordinate.Coordinate{
// 		&coordinate.Coordinate{1, 4},
// 		&coordinate.Coordinate{1, 6},
// 		&coordinate.Coordinate{2, 2},
// 	},
// 	expectedDistanceMatrix: [][]float64{
// 		{0, 2, 3},
// 		{2, 0, 5},
// 		{3, 5, 0},
// 	},
// }

// func TestDistanceMatrix(t *testing.T) {
// 	// Arrange
// 	var entryList []Entry
// 	for idx := range distanceMatrixTestCase.coordList {
// 		coord := distanceMatrixTestCase.coordList[idx]
// 		entry := &LeafEntry{
// 			Object:   coord,
// 			ObjectID: idx,
// 		}
// 		entryList = append(entryList, entry)
// 	}
// 	splitMecha := &SplitMST{
// 		DistCalc: &distance.ManhattanDistance{},
// 	}
// 	expectedDistanceMatrix := distanceMatrixTestCase.expectedDistanceMatrix

// 	// Act
// 	distanceMatrixResult := splitMecha.getDistanceMatrix(entryList)

// 	// Assert
// 	assert.Equal(t, expectedDistanceMatrix, distanceMatrixResult, "Error! expected distance matrix = %+v\n got %+v", expectedDistanceMatrix, distanceMatrixResult)
// }

var mstTestCase = struct {
	distCalc       DistanceCalculator
	coordList      []*coordinate.Coordinate
	expectedLength float64
}{
	distCalc: &distance.ManhattanDistance{},
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
	expectedLength: 31.,
}

func TestInternalMST(t *testing.T) {
	// Arrange
	splitMecha := &SplitMST{
		DistCalc: &distance.ManhattanDistance{},
	}
	var entryList []Entry
	for idx := range mstTestCase.coordList {
		coord := mstTestCase.coordList[idx]
		entry := &LeafEntry{
			Object:   coord,
			ObjectID: idx,
		}
		entryList = append(entryList, entry)
	}
	expectedLength := mstTestCase.expectedLength

	// Act
	// Act Testing MST
	mstEdgeList := splitMecha.getMSTEdgeList(entryList)
	resultLength := 0.
	for idx := range mstEdgeList {
		edge := mstEdgeList[idx]
		resultLength += edge.Length
	}
	// Act Testing Partition
	newEntryList1, newEntryList2 := splitMecha.partitionEntryList(entryList)
	resultEntryNumber := len(newEntryList1) + len(newEntryList2)

	// Assert
	t.Run("Test MST", func(t *testing.T) {
		assert.Equalf(t, expectedLength, resultLength, "Error! expected MST Length = %f, got %f", expectedLength, resultLength)
	})
	t.Run("Test Partition", func(t *testing.T) {
		assert.Equalf(t, len(entryList), resultEntryNumber, "Error! Entries number expected = %d, got %d", len(entryList), resultEntryNumber)
	})

}

var createNewNodeTestCase = struct {
	distCalc       DistanceCalculator
	coordList      []*coordinate.Coordinate
	expectedRadius float64
}{
	distCalc: &distance.ManhattanDistance{},
	coordList: []*coordinate.Coordinate{
		&coordinate.Coordinate{1, 4},
		&coordinate.Coordinate{1, 6},
		&coordinate.Coordinate{2, 2},
	},
	expectedRadius: 3.,
}

func TestCreateNewNodeWithExistingEntries(t *testing.T) {
	// Arrange
	splitMecha := &SplitMST{
		DistCalc: &distance.ManhattanDistance{},
	}
	var entryList []Entry
	for idx := range createNewNodeTestCase.coordList {
		coord := createNewNodeTestCase.coordList[idx]
		entry := &LeafEntry{
			Object:   coord,
			ObjectID: idx,
		}
		entryList = append(entryList, entry)
	}
	expectedRadius := createNewNodeTestCase.expectedRadius
	expectedCentroid := entryList[0]

	// Act
	node := splitMecha.createNewNodeWithExistingEntries(entryList)
	resultRadius := node.GetRadius()

	// Assert
	t.Run("test centroid choice", func(t *testing.T) {
		assert.Equalf(t, expectedCentroid, node.CentroidEntry, "Error! expected centroid entry %+v, got %+v", expectedCentroid, node.CentroidEntry)
	})
	t.Run("test New Node radius", func(t *testing.T) {
		assert.Equalf(t, expectedRadius, resultRadius, "Error! radius expected = %f, got %f", expectedRadius, resultRadius)
	})
}

var replaceEntryTestCase = struct {
	distCalc                       DistanceCalculator
	coordList                      []*coordinate.Coordinate
	replacementCoord               *coordinate.Coordinate
	expectedDistanceFromParentList [][]float64
}{
	distCalc: &distance.ManhattanDistance{},
	coordList: []*coordinate.Coordinate{
		&coordinate.Coordinate{1, 4},
		&coordinate.Coordinate{1, 6},
		&coordinate.Coordinate{2, 2},
	},
	replacementCoord: &coordinate.Coordinate{2, 3},
	expectedDistanceFromParentList: [][]float64{
		{0, 4, 1},
		{0, 2, 3},
		{0, 2, 2},
	},
}

func TestReplaceEntry(t *testing.T) {

	coordList := replaceEntryTestCase.coordList
	splitMecha := &SplitMST{
		DistCalc: &distance.ManhattanDistance{},
	}
	for replaceIdx := range coordList {
		// Arrange
		var entryList []Entry
		for idx := range coordList {
			coord := coordList[idx]
			entry := &LeafEntry{
				Object:   coord,
				ObjectID: idx,
			}
			entryList = append(entryList, entry)
		}
		node := splitMecha.createNewNodeWithExistingEntries(entryList)
		replacementEntry := &LeafEntry{
			Object:   replaceEntryTestCase.replacementCoord,
			ObjectID: len(coordList),
		}

		// Act
		entryToReplace := entryList[replaceIdx]
		splitMecha.replaceNodeEntry(node, entryToReplace, replacementEntry)
		expectedDistanceFromParentList := replaceEntryTestCase.expectedDistanceFromParentList[replaceIdx]

		// Assert
		t.Run(fmt.Sprintf("Replace Node %d", replaceIdx), func(t *testing.T) {
			for idx := range node.EntryList {
				entry := node.EntryList[idx]
				resultDistanceFromParent := entry.GetDistanceFromParent()
				expectedDistanceFromParent := expectedDistanceFromParentList[idx]
				assert.Equalf(t, expectedDistanceFromParent, resultDistanceFromParent, "Error! Expected distance from parent of entry %d = %f, got %f", idx, expectedDistanceFromParent, resultDistanceFromParent)
			}
		})

	}

}
