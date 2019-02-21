package mtree_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mtree"
)

func TestGetCopy(t *testing.T) {
	// Arrange
	tree := mtree.NewTree(2, 4, distCalc)
	objList := generateObjectList(20)
	for idx := range objList {
		tree.Insert(objList[idx])
	}

	// Act
	newTree := tree.GetCopy()
	tree.Remove(objList[3])
	tree.Remove(objList[5])
	newTree.Remove(objList[4])

	// Assert
	fmt.Println(tree.ObjectCount, newTree.ObjectCount)
	assert.True(t, newTree.ObjectCount == tree.ObjectCount+1, "Error! tree and newTree ObjectCount should have differed!")
}
