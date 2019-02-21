package mtree

import "log"

// Remove an object from the tree, will panic if the removed object does not exist
func (tree *Tree) Remove(object Object) {
	removed := tree.removeFromTree(tree.root, object)
	if !removed {
		log.Fatal("Error! Removed Fail!")
	}
	tree.ObjectCount--

	// while the root is a branch and it only has 1 entry, let that entry becomes the new root to reduce the tree height and so fasten the queries (insert, remove or knn_search)
	treeEntryList := tree.root.getEntryList()
	for !tree.root.isLeaf() && len(treeEntryList) == 1 {
		newRoot := treeEntryList[0].(node)
		newRoot.setParent(nil)
		tree.root = newRoot
	}
}

func (tree *Tree) removeFromTree(currentNode node, object Object) bool {

	if currentNode.isLeaf() {
		currentLeaf := currentNode.(*leaf)
		if currentLeaf.containsObject(object) {
			currentLeaf.removeObject(object)
			currentLeaf.updateRadius()
			return true
		}
		return false
	}

	currentBranch := currentNode.(*branch)
	for idx := range currentBranch.entryList {
		nextNode := currentBranch.entryList[idx].(node)

		distanceToNextNode := tree.distCalc.GetDistance(object, nextNode.getCentroidObject())
		if distanceToNextNode > nextNode.getRadius() {
			continue
		}

		removed := tree.removeFromTree(nextNode, object)
		if !removed {
			continue
		}

		// if remove occures in nextNode, and resulting its entries to become underFlown (has less than minimum entries) or  as a result of its entries' merging, then we have to merge it with the closest different node if exist other node
		if !nextNode.isUnderFlown(tree.minEntry) {
			currentBranch.updateRadius()
			return true
		}

		// if no other node exist just let it be
		closestNode := getClosestNode(nextNode, currentBranch.entryList, tree.distCalc)
		if closestNode == nil {
			currentBranch.updateRadius()
			return true
		}

		// if merging with the closestNode make the closestNode entries exceed maxEntry, then just split the combined entries, and replace both closestNode and nextNode with the split result
		if len(nextNode.getEntryList())+len(closestNode.getEntryList()) > tree.maxEntry {
			nextNodeEntryList := nextNode.getEntryList()
			closestNodeEntryList := closestNode.getEntryList()
			for idx := range nextNodeEntryList {
				nextNodeEntry := nextNodeEntryList[idx]
				closestNodeEntryList = append(closestNodeEntryList, nextNodeEntry)
			}
			currentBranch.removeEntry(nextNode)
			currentBranch.removeEntry(closestNode)

			newNode1, newNode2 := tree.splitMecha.split(closestNodeEntryList)
			currentBranch.insertEntry(newNode1)
			currentBranch.insertEntry(newNode2)

			dist1 := tree.distCalc.GetDistance(currentBranch.getCentroidObject(), newNode1.getCentroidObject())
			newNode1.setDistanceFromParent(dist1)

			dist2 := tree.distCalc.GetDistance(currentBranch.getCentroidObject(), newNode2.getCentroidObject())
			newNode2.setDistanceFromParent(dist2)
			currentBranch.updateRadius()
			return true
		}

		// remove nextNode and insert all its entries into the closestNode
		nextNodeEntryList := nextNode.getEntryList()
		for idx := range nextNodeEntryList {
			nextNodeEntry := nextNodeEntryList[idx]
			closestNode.insertEntry(nextNodeEntry)
			distFromParent := tree.distCalc.GetDistance(nextNodeEntry.getCentroidObject(), closestNode.getCentroidObject())
			nextNodeEntry.setDistanceFromParent(distFromParent)
		}
		closestNode.updateRadius()
		currentBranch.removeEntry(nextNode)
		currentBranch.updateRadius()
		return true
	}
	return false
}
