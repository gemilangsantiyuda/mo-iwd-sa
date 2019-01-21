package mtree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/vroup/mo-iwd-sa/mtree"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/distance"
)

var insertTestCase = struct {
	coordList []*coordinate.Coordinate
	maxEntry  int
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
	maxEntry: 4,
}

func TestInsertAndRemove(t *testing.T) {
	// Arrange
	coordList := insertTestCase.coordList
	distCalc := &distance.ManhattanDistance{}
	splitMecha := &mtree.SplitMST{
		MaxEntry: insertTestCase.maxEntry,
		DistCalc: distCalc,
	}
	tree := mtree.NewTree(insertTestCase.maxEntry, splitMecha, distCalc)

	// Act
	t.Run("Test Insertion", func(t *testing.T) {
		for idx := range coordList {
			coord := coordList[idx]
			fmt.Println("--------------------")
			fmt.Println("inserting", coord)
			tree.Insert(coord, idx)
			// Act Continued + Assert
			numEntries := traverseAndTest(tree.Root, t, idx+1)
			t.Run("Test number of entries inserted consistent", func(t *testing.T) {
				require.Equalf(t, idx+1, numEntries, "Error! Expected num of entries %d, got %d", idx+1, numEntries)
			})
		}
	})

	// Act
	t.Run("Test Removal", func(t *testing.T) {
		for idx := range coordList {
			coord := coordList[idx]
			removeSuccess := tree.Remove(tree.Root, coord, idx)
			fmt.Println(idx+1, removeSuccess)
			expectedNumEntries := len(coordList) - idx - 1
			// Act Continue + Assert
			numEntries := traverseAndTest(tree.Root, t, idx+1)
			t.Run("Test number of entries consistent after removal", func(t *testing.T) {
				require.Equalf(t, expectedNumEntries, numEntries, "Error! Expected num of entries %d, got %d", expectedNumEntries, numEntries)
			})
		}
	})

}

func traverseAndTest(node *mtree.Node, t *testing.T, subTestIdx int) int {

	// Act
	nodeRadius := node.GetRadius()

	fmt.Println(node.IsLeaf(), node)
	// Assert
	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		fmt.Println("----->", entry)
		radius := entry.GetDistanceFromParent() + entry.GetRadius()
		t.Run("Checking recalculated radius smaller than node's radius", func(t *testing.T) {
			assert.Truef(t, nodeRadius >= radius, "Error! node's radius (%f) smaller than recalculated radius (%f), currentNode = %v,, current subtest : %d", nodeRadius, radius, node, subTestIdx)
		})
	}

	numEntries := 0
	for idx := range node.EntryList {
		if _, isNode := node.EntryList[idx].(*mtree.Node); isNode {
			nextNode := node.EntryList[idx].(*mtree.Node)
			numEntries += traverseAndTest(nextNode, t, subTestIdx)
		} else {
			numEntries += 1
		}

	}
	return numEntries
}
