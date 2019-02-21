package mtree

import "math"

// Insert an object to the mtree
func (tree *Tree) Insert(object Object) {
	newLeafEntry := &leafEntry{
		object: object,
	}
	// fmt.Println("new Leaf ", newLeafEntry)
	newNode1, newNode2 := tree.insertNewLeafEntry(tree.root, newLeafEntry)
	if newNode1 != nil {
		newRoot := &branch{}
		newRoot.insertEntry(newNode1)
		newRoot.insertEntry(newNode2)
		dist2 := tree.distCalc.GetDistance(newNode2.getCentroidObject(), newRoot.getCentroidObject())
		newNode2.setDistanceFromParent(dist2)
		radius := math.Max(newNode1.getRadius(), dist2+newNode2.getRadius())
		newRoot.radius = radius
		tree.root = newRoot
	}

	tree.ObjectCount++
	// traverse(tree.root)
}

func (tree *Tree) insertNewLeafEntry(currentNode node, newLeafEntry *leafEntry) (node, node) {
	if currentNode.isLeaf() {
		currentLeaf := currentNode.(*leaf)
		currentLeaf.insertEntry(newLeafEntry)

		// if the leaf's entries does not exceed maxEntry, then just return the radius,, else we split and return the promoted entry (2 new nodes)
		if len(currentLeaf.entryList) <= tree.maxEntry {
			distToParent := tree.distCalc.GetDistance(newLeafEntry.object, currentLeaf.getCentroidObject())
			newLeafEntry.setDistanceFromParent(distToParent)
			if currentLeaf.radius < distToParent {
				currentLeaf.radius = distToParent
			}
			return nil, nil
		}
		// split and make 2 new nodes
		newNode1, newNode2 := tree.splitMecha.split(currentLeaf.entryList)
		return newNode1, newNode2
	}

	currentBranch := currentNode.(*branch)
	nextNode := chooseBestNextNode(currentBranch.entryList, newLeafEntry, tree.distCalc)
	newNode1, newNode2 := tree.insertNewLeafEntry(nextNode, newLeafEntry)

	if newNode1 != nil {
		currentBranch.removeEntry(nextNode)
		currentBranch.insertEntry(newNode1)
		currentBranch.insertEntry(newNode2)
		if len(currentBranch.entryList) <= tree.maxEntry {
			dist1 := tree.distCalc.GetDistance(newNode1.getCentroidObject(), currentBranch.getCentroidObject())
			newNode1.setDistanceFromParent(dist1)
			dist2 := tree.distCalc.GetDistance(newNode2.getCentroidObject(), currentBranch.getCentroidObject())
			newNode2.setDistanceFromParent(dist2)
			if currentBranch.radius < dist1+newNode1.getRadius() {
				currentBranch.radius = dist1 + newNode1.getRadius()
			}
			if currentBranch.radius < dist2+newNode2.getRadius() {
				currentBranch.radius = dist2 + newNode2.getRadius()
			}
			return nil, nil
		}
		newNode1, newNode2 = tree.splitMecha.split(currentBranch.entryList)
		return newNode1, newNode2
	}
	newRadius := nextNode.getDistanceFromParent() + nextNode.getRadius()
	if newRadius > currentBranch.radius {
		currentBranch.radius = newRadius
	}

	return nil, nil
}
