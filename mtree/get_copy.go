package mtree

// GetCopy return deep copy by value of the entire tree
func (tree *Tree) GetCopy() *Tree {
	newRoot := deepCopy(tree.Root)
	newTree := &Tree{
		Root:           newRoot,
		MaxEntry:       tree.MaxEntry,
		SplitMechanism: tree.SplitMechanism,
		DistCalc:       tree.DistCalc,
	}
	return newTree
}

func deepCopy(node *Node) *Node {

	nodeCentroidIdx := node.GetCentroidIdx()
	newEntryList := make([]Entry, 0)

	newNode := Node{
		Radius:             node.Radius,
		DistanceFromParent: node.DistanceFromParent,
	}

	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		if _, isLeafEntry := entry.(*LeafEntry); isLeafEntry {
			leafEntry := entry.(*LeafEntry)
			newLeafEntry := LeafEntry{
				Object:             leafEntry.Object,
				ObjectID:           leafEntry.ObjectID,
				Parent:             &newNode,
				DistanceFromParent: leafEntry.DistanceFromParent,
			}
			newEntryList = append(newEntryList, &newLeafEntry)
		} else {
			nodeEntry := entry.(*Node)
			newNodeEntry := deepCopy(nodeEntry)
			newNodeEntry.SetParent(&newNode)
			newEntryList = append(newEntryList, newNodeEntry)
		}
	}
	newNode.EntryList = newEntryList
	newNode.CentroidEntry = newEntryList[nodeCentroidIdx]

	return &newNode
}
