package mtree

import (
	"math"

	"github.com/vroup/mo-iwd-sa/object"
)

// Insert make a new leaf entry of the new inserted object then insert it into the fittest leaf node
func (tree *Tree) Insert(object object.Object, objectID string) {

	newLeafEntry := &LeafEntry{
		Object:             object,
		ObjectID:           objectID,
		Parent:             nil,
		DistanceFromParent: 0,
	}

	tree.insertLeafEntry(tree.Root, newLeafEntry)
	distanceFromNewEntryToRoot := tree.DistCalc.GetDistance(tree.Root, newLeafEntry)
	if distanceFromNewEntryToRoot > tree.Root.GetRadius() {
		tree.Root.SetRadius(distanceFromNewEntryToRoot)
	}
	tree.ObjectCount++
}

func (tree *Tree) insertLeafEntry(currentNode *Node, newLeafEntry *LeafEntry) {
	// If leaf node is met, and it is not full then insert directly
	// else split the leaf node
	if currentNode.IsLeaf() {
		if len(currentNode.EntryList) < tree.MaxEntry {
			currentNode.InsertEntry(newLeafEntry)
		} else {
			// check if perhaps the split create a new root
			newRoot := tree.SplitMechanism.Split(currentNode, newLeafEntry)
			if newRoot != nil {
				tree.Root = newRoot
			}
		}
		return
	}

	// else if it's internal node then find the most suitable node's entry to traverse
	// the first options are the nodes which already cover the newLeafEntry
	nextNode := tree.findNearestCoveringNode(currentNode.EntryList, newLeafEntry)
	if nextNode != nil {
		tree.insertLeafEntry(nextNode, newLeafEntry)
	} else {
		// if none found then find the node that need less to expand
		nextNode, distanceToNextNode := tree.findLeastRadiusExpansionNode(currentNode.EntryList, newLeafEntry)

		if distanceToNextNode > nextNode.GetRadius() {
			nextNode.SetRadius(distanceToNextNode)
		}
		tree.insertLeafEntry(nextNode, newLeafEntry)
	}
}

func (tree *Tree) findNearestCoveringNode(entryList []Entry, newLeafEntry *LeafEntry) *Node {

	nearesetDistance := math.Inf(1)
	nearestNodeIdx := -1
	for idx := range entryList {
		entry := entryList[idx]
		distanceToEntry := tree.DistCalc.GetDistance(newLeafEntry, entry)
		if distanceToEntry > entry.GetRadius() {
			continue
		}
		if distanceToEntry < nearesetDistance {
			nearesetDistance = distanceToEntry
			nearestNodeIdx = idx
		}
	}

	if nearestNodeIdx == -1 {
		return nil
	}
	nearestNode := entryList[nearestNodeIdx].(*Node)
	return nearestNode
}

func (tree *Tree) findLeastRadiusExpansionNode(entryList []Entry, newLeafEntry *LeafEntry) (*Node, float64) {
	leastRadiusExpansion := math.Inf(1)
	distanceToNextNode := math.Inf(1)
	nextNode := &Node{}

	for idx := range entryList {
		nextNodeCandidate := entryList[idx].(*Node)
		distanceToNextNodeCandidate := tree.DistCalc.GetDistance(nextNodeCandidate, newLeafEntry)
		radiusExpansion := distanceToNextNodeCandidate - nextNodeCandidate.GetRadius()
		if radiusExpansion < leastRadiusExpansion {
			leastRadiusExpansion = radiusExpansion
			distanceToNextNode = distanceToNextNodeCandidate
			nextNode = nextNodeCandidate
		}
	}
	return nextNode, distanceToNextNode
}
