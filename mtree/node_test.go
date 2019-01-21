package mtree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vroup/mo-iwd-sa/mtree"
)

func TestInsertAndRemoveEntry(t *testing.T) {

	// Arrange
	mockLeafNode := &mtree.Node{
		EntryList: make([]mtree.Entry, 0),
	}
	mockInternalNode := &mtree.Node{
		EntryList: make([]mtree.Entry, 0),
	}
	const EntryNum = 10
	mockLeafEntryList := make([]*mtree.LeafEntry, 0)
	mockNodeEntryList := make([]*mtree.Node, 0)
	for idx := 0; idx < EntryNum; idx++ {
		mockLeafEntry := mtree.LeafEntry{
			ObjectID: idx,
		}
		mockNodeEntry := mtree.Node{
			Radius: float64(idx),
		}
		mockLeafEntryList = append(mockLeafEntryList, &mockLeafEntry)
		mockNodeEntryList = append(mockNodeEntryList, &mockNodeEntry)
	}
	leafEntryToRemove := mockLeafEntryList[1]
	nodeEntryToRemove := mockNodeEntryList[1]

	// Act
	for idx := 0; idx < EntryNum; idx++ {
		mockLeafNode.InsertEntry(mockLeafEntryList[idx])
		mockInternalNode.InsertEntry(mockNodeEntryList[idx])
	}
	mockLeafNode.RemoveEntry(leafEntryToRemove)
	mockInternalNode.RemoveEntry(nodeEntryToRemove)

	// Assert
	t.Run("Leaf Node Insert Entry Test", func(t *testing.T) {
		assert.Lenf(t, mockLeafNode.EntryList, EntryNum-1, "Length of mockLeafNode.EntryList must be equal to %d", EntryNum-1)
	})
	t.Run("Internal Node Insert Entry Test", func(t *testing.T) {
		assert.Lenf(t, mockInternalNode.EntryList, EntryNum-1, "Len of mockInternalNode.EntryList must be equal to %d", EntryNum-1)
	})
	t.Run("Check Leaf Node's entries' parent", func(t *testing.T) {
		for idx := range mockLeafEntryList {
			entry := mockLeafEntryList[idx]
			entryParent := entry.GetParent()
			assert.Equalf(t, mockLeafNode, entryParent, "entry's parent expected = %p, got = %p", mockLeafNode, entryParent)
		}
	})
	t.Run("Check Internal Node's entries' parent", func(t *testing.T) {
		for idx := range mockNodeEntryList {
			entry := mockNodeEntryList[idx]
			entryParent := entry.GetParent()
			assert.Equalf(t, mockInternalNode, entryParent, "entry's parent expected = %p, got = %p", mockInternalNode, entryParent)
		}
	})
	t.Run("Check Existence of removed leaf entry", func(t *testing.T) {
		assert.NotContains(t, mockLeafNode.EntryList, leafEntryToRemove, "Error! leaf node still contains removed leaf entry")
	})
	t.Run("Check existence of removed node entry", func(t *testing.T) {
		assert.NotContains(t, mockInternalNode.EntryList, nodeEntryToRemove, "Error! internal node still contains removed node entry")
	})

}

func TestParentDistance(t *testing.T) {

	// Arrange
	mockNode := &mtree.Node{}
	const expectedDistance = 999.99

	// Act
	mockNode.SetDistanceFromParent(expectedDistance)
	resultDistance := mockNode.GetDistanceFromParent()

	// Assert
	assert.Equalf(t, expectedDistance, resultDistance, "Error! Expected distance = %f, got %f", expectedDistance, resultDistance)
}

func TestRadius(t *testing.T) {

	// Arrange
	mockNode := &mtree.Node{}
	mockLeafEntry := &mtree.LeafEntry{}
	const expectedNodeRadius = 0.333

	// Act
	mockNode.SetRadius(expectedNodeRadius)
	resultNodeRadius := mockNode.GetRadius()
	resultLeafEntryRadius := mockLeafEntry.GetRadius()

	// Assert
	t.Run("Test Node Radius", func(t *testing.T) {
		assert.Equalf(t, expectedNodeRadius, resultNodeRadius, "Error! expected node radius = %f, got %f", expectedNodeRadius, resultNodeRadius)
	})
	t.Run("Test Leaf Entry Radius", func(t *testing.T) {
		assert.Equalf(t, 0., resultLeafEntryRadius, "Error! expected leaf entry radius = 0, got %f", resultLeafEntryRadius)
	})
}
