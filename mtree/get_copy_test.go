package mtree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vroup/mo-iwd-sa/coordinate"
	"github.com/vroup/mo-iwd-sa/distance"
	"github.com/vroup/mo-iwd-sa/mtree"
)

var getCopyCoordList = []*coordinate.Coordinate{
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
}

func TestGetCopy(t *testing.T) {
	// Arrange
	coordList := getCopyCoordList
	distCalc := &distance.ManhattanDistance{}
	splitMecha := &mtree.SplitMST{
		DistCalc: distCalc,
		MaxEntry: 4,
	}
	tree := mtree.NewTree(4, splitMecha, distCalc)
	for idx := range coordList {
		coord := coordList[idx]
		tree.Insert(coord, idx)
	}
	treeNumEntry := traverseAndTest(tree.Root, t, 0)

	// Act 1
	newTree := tree.GetCopy()

	// Assert
	t.Run("Comparing two tree node by node by traversal", func(t *testing.T) {
		compareNodeByNode(t, tree.Root, newTree.Root)
	})

	// Act 2
	coordToRemove := coordList[2]
	tree.Remove(tree.Root, coordToRemove, 2)
	expectedNumEntry := 14
	newTreeNumEntry := traverseAndTest(newTree.Root, t, 0)
	treeNumEntry = traverseAndTest(tree.Root, t, 0)
	fmt.Println(treeNumEntry, newTreeNumEntry)

	// Assert 2
	t.Run("Test num of entry in newTree after removal on tree", func(t *testing.T) {
		assert.Equalf(t, expectedNumEntry, newTreeNumEntry, "Error! expected num entry after romval %d, got %d", expectedNumEntry, newTreeNumEntry)
	})

}

func compareNodeByNode(t *testing.T, node *mtree.Node, newNode *mtree.Node) {
	// fmt.Println(node)
	// fmt.Println(newNode)
	// fmt.Println("---------------")
	t.Run("Compare Radius", func(t *testing.T) {
		assert.Equalf(t, node.Radius, newNode.Radius, "Error! different radius , expected %f, got %f", node.Radius, newNode.Radius)
	})

	t.Run("Compare number of entry in entrylist", func(t *testing.T) {
		assert.Equalf(t, len(node.EntryList), len(newNode.EntryList), "Error! expected num of entries %d, got %d", len(node.EntryList), len(newNode.EntryList))
	})

	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		newEntry := newNode.EntryList[idx]
		if _, isNode := entry.(*mtree.Node); isNode {
			nextNode := entry.(*mtree.Node)
			nextNewNode := newEntry.(*mtree.Node)
			compareNodeByNode(t, nextNode, nextNewNode)
		}
	}
}
