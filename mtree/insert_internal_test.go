package mtree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vroup/mo-iwd-sa/coordinate"

	"github.com/vroup/mo-iwd-sa/distance"
)

var testCase = struct {
	coordList  []*coordinate.Coordinate
	radiusList []float64
	newCoord1  *coordinate.Coordinate
	newCoord2  *coordinate.Coordinate
}{
	coordList: []*coordinate.Coordinate{
		&coordinate.Coordinate{0, 0},
		&coordinate.Coordinate{1, 1},
	},
	radiusList: []float64{
		4.,
		3.,
	},
	newCoord1: &coordinate.Coordinate{1, 2},
	newCoord2: &coordinate.Coordinate{-0, 5},
}

func TestChooseEntry(t *testing.T) {
	// Arrange
	coordList := testCase.coordList
	radiusList := testCase.radiusList
	var entryList []Entry
	for idx := range coordList {
		coord := coordList[idx]
		leafEntry := &LeafEntry{
			Object:   coord,
			ObjectID: idx,
		}
		node := &Node{}
		node.InsertEntry(leafEntry)
		node.SetRadius(radiusList[idx])
		entryList = append(entryList, node)
	}
	newEntry1 := &LeafEntry{
		Object: testCase.newCoord1,
	}
	newEntry2 := &LeafEntry{
		Object: testCase.newCoord2,
	}
	distCalc := &distance.ManhattanDistance{}
	splitMecha := &SplitMST{
		DistCalc: distCalc,
	}
	tree := NewTree(3, splitMecha, distCalc)
	expectedNearestCoverNode := entryList[1]
	expectedCoverNodeDistance := 1.
	expectedLeastExpansion := entryList[0]
	expectedLeastExpansionDistance := 5.

	// Act
	nearestCoverNode1 := tree.findNearestCoveringNode(entryList, newEntry1)
	nearestCoverNodeDistance := distCalc.GetDistance(newEntry1, nearestCoverNode1)
	nearestCoverNode2 := tree.findNearestCoveringNode(entryList, newEntry2)
	leastExpansionNode, distance := tree.findLeastRadiusExpansionNode(entryList, newEntry2)

	// Assert
	t.Run("(1)Test Finding Nearest Cover Node (non nil)", func(t *testing.T) {
		assert.Equalf(t, expectedNearestCoverNode, nearestCoverNode1, "Error! expected nearest cover node %v, got %v", expectedNearestCoverNode, nearestCoverNode1)
	})
	t.Run("(1)Test nearest cover node distance", func(t *testing.T) {
		assert.Equalf(t, expectedCoverNodeDistance, nearestCoverNodeDistance, "Error! expected cover distance %f, got %f", expectedCoverNodeDistance, nearestCoverNodeDistance)
	})
	t.Run("(2)Test Finding nearest cover node (nil)", func(t *testing.T) {
		assert.Nil(t, nearestCoverNode2, "Error! expected nil, got %v", nearestCoverNode2)
	})
	t.Run("Test Finding Least Expansion", func(t *testing.T) {
		assert.Equalf(t, expectedLeastExpansion, leastExpansionNode, "Error! expected least expansion node %v, got %v", expectedLeastExpansion, leastExpansionNode)
		assert.Equalf(t, expectedLeastExpansionDistance, distance, "Error! expected least expansion distance %f, got %f", expectedLeastExpansionDistance, distance)
	})
}
