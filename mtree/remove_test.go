package mtree_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mtree"
)

func TestRemove(t *testing.T) {
	// t.Run("remove from root", testRemoveFromRoot)
	t.Run("remove until empty from complex tree", testRemoveComplete)
}

func testRemoveFromRoot(t *testing.T) {
	// Arrange
	tree := mtree.NewTree(2, 4, distCalc)
	objList := generateObjectList(4)
	for idx := range objList {
		obj := objList[idx]
		tree.Insert(obj)
	}
	removeIdxList := rand.Perm(4)

	// Act
	for _, idx := range removeIdxList {
		obj := objList[idx]
		fmt.Println("REMOVING ", obj)
		tree.Remove(obj)
	}

	// Assert
	assert.True(t, false, "EEE")
}

func testRemoveComplete(t *testing.T) {
	// Arrange
	tree := mtree.NewTree(2, 4, distCalc)
	objList := generateObjectList(100)
	for idx := range objList {
		obj := objList[idx]
		tree.Insert(obj)
	}
	removeIdxList := rand.Perm(100)

	// Act
	for _, idx := range removeIdxList {
		obj := objList[idx]
		fmt.Printf("\n\n\n\n REMOVING %+v\n", obj)
		tree.Remove(obj)
	}

	// Assert
	assert.True(t, false, "EEE")
}
