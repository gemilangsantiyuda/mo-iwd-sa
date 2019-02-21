package mtree

// GetCopy of a already constructed tree, because reinsertion of all objects is way too costly in runtime
func (tree *Tree) GetCopy() *Tree {
	newTree := &Tree{
		maxEntry:    tree.maxEntry,
		minEntry:    tree.minEntry,
		distCalc:    tree.distCalc,
		ObjectCount: tree.ObjectCount,
		splitMecha:  tree.splitMecha,
	}

	var newRoot node
	root := tree.root
	newEntryList := copyEntryList(tree.root)
	if root.isLeaf() {
		newRoot = &leaf{
			radius:         root.getRadius(),
			centroidObject: root.getCentroidObject(),
			entryList:      newEntryList,
		}
	} else {
		newRoot = &branch{
			radius:         root.getRadius(),
			centroidObject: root.getCentroidObject(),
			entryList:      newEntryList,
		}
	}
	for idx := range newEntryList {
		newEntryList[idx].setParent(newRoot)
	}
	newTree.root = newRoot
	return newTree
}

func copyEntryList(currentNode node) []entry {

	entryList := currentNode.getEntryList()
	newEntryList := make([]entry, 0)

	if currentNode.isLeaf() {
		for idx := range entryList {
			ent := entryList[idx]
			newLeafEntry := &leafEntry{
				distanceFromParent: ent.getDistanceFromParent(),
				object:             ent.getCentroidObject(),
			}
			newEntryList = append(newEntryList, newLeafEntry)
		}
		return newEntryList
	}
	for idx := range entryList {
		ent := entryList[idx]
		var newEntry entry
		newEntryEntryList := copyEntryList(ent.(node))
		if _, ok := ent.(*leaf); ok {
			newEntry = &leaf{
				radius:             ent.getRadius(),
				distanceFromParent: ent.getDistanceFromParent(),
				centroidObject:     ent.getCentroidObject(),
				entryList:          newEntryEntryList,
			}
		} else {
			newEntry = &branch{
				radius:             ent.getRadius(),
				distanceFromParent: ent.getDistanceFromParent(),
				centroidObject:     ent.getCentroidObject(),
				entryList:          newEntryEntryList,
			}
		}
		for idx2 := range newEntryEntryList {
			newEntryEntryList[idx2].setParent(newEntry.(node))
		}
		newEntryList = append(newEntryList, newEntry)
	}

	return newEntryList
}
