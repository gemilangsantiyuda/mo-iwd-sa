package mtree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/distance"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/object"
)

type DistanceCalculator interface {
	GetDistance(object.Object, object.Object) float64
}

var testCase = struct {
	distCalc                                   DistanceCalculator
	maxEntry                                   int
	coordList                                  []*coordinate.Coordinate
	newEntryIdx                                int
	expectedCentroidIdx1, expectedCentroidIdx2 int
}{
	distCalc: &distance.ManhattanDistance{},
	maxEntry: 5,
	coordList: []*coordinate.Coordinate{
		&coordinate.Coordinate{2, 1},
		&coordinate.Coordinate{3, 3},
		&coordinate.Coordinate{3, 5},
		&coordinate.Coordinate{5, 1},
		&coordinate.Coordinate{1, 5},
		&coordinate.Coordinate{5, 4},
	},
	newEntryIdx:          5,
	expectedCentroidIdx1: 1,
	expectedCentroidIdx2: 5,
}

func TestSplit(t *testing.T) {

	// Arrange
	splitMecha := &mtree.SplitMST{
		MaxEntry: testCase.maxEntry,
		DistCalc: &distance.ManhattanDistance{},
	}
	var entryList []mtree.Entry
	coordList := testCase.coordList
	for idx := 0; idx < testCase.maxEntry; idx++ {
		coord := coordList[idx]
		entry := &mtree.LeafEntry{
			Object:   coord,
			ObjectID: idx,
		}
		entryList = append(entryList, entry)
	}

	newCoord := coordList[testCase.newEntryIdx]
	newEntry := &mtree.LeafEntry{
		Object:   newCoord,
		ObjectID: testCase.newEntryIdx,
	}

	// Arrange Continue + Act & Assert
	t.Run("Test Split on Root Node", func(t *testing.T) {
		testRootSplit(t, splitMecha, entryList, newEntry)
	})
	t.Run("Test Split on non Root Node", func(t *testing.T) {
		testSplit(t, splitMecha, entryList, newEntry)
	})

}

func testRootSplit(t *testing.T, splitMecha mtree.SplitMechanism, entryList []mtree.Entry, newEntry mtree.Entry) {
	// Arrange
	root := &mtree.Node{
		Parent:    nil,
		EntryList: entryList,
	}
	expectedCentroid1 := entryList[testCase.expectedCentroidIdx1]
	expectedCentroid2 := newEntry

	// Act
	newRoot := splitMecha.Split(root, newEntry)
	newNode1, newNode2 := newRoot.EntryList[0].(*mtree.Node), newRoot.EntryList[1].(*mtree.Node)
	resultCentroid1, resultCentroid2 := newNode1.CentroidEntry.(*mtree.LeafEntry), newNode2.CentroidEntry.(*mtree.LeafEntry)

	// Assert
	t.Run("Test root split cause new root", func(t *testing.T) {
		assert.NotNil(t, newRoot, "Error! got new root nil")
	})
	t.Run("Test new root has 2 entries", func(t *testing.T) {
		assert.Equalf(t, 2, len(newRoot.EntryList), "Error, expected %d entries, got %d", 2, len(newRoot.EntryList))
	})
	t.Run("Test newNodes' centroid", func(t *testing.T) {
		assert.Equalf(t, expectedCentroid1, resultCentroid1, "Error! expected centroid of newNode1 %v, got %v", expectedCentroid1, resultCentroid1)
		assert.Equalf(t, expectedCentroid2, resultCentroid2, "Error! expected centroid of newNode2 %v, got %v", expectedCentroid2, resultCentroid2)
	})
}

func testSplit(t *testing.T, splitMecha mtree.SplitMechanism, entryList []mtree.Entry, newEntry mtree.Entry) {
	// Arrange
	node := &mtree.Node{
		EntryList: entryList,
	}
	parent := &mtree.Node{}
	parent.InsertEntry(node)

	// Act
	newRoot := splitMecha.Split(node, newEntry)

	// Arrange
	t.Run("Test non root split cause no new root when parent entry is not full", func(t *testing.T) {
		assert.Nil(t, newRoot, "Error! Got new root not nil")
	})
	t.Run("Test parent has new entry", func(t *testing.T) {
		assert.Equalf(t, 2, len(parent.EntryList), "Error! expected %d entries, got %d", 2, len(parent.EntryList))
	})
}
