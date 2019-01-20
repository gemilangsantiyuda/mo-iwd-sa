package mtree

// NewTree initiate new tree
func NewTree(maxEntry int, splitMechanism SplitMechanism, distCalc DistanceCalculator) *Tree {

	newEntryList := make([]Entry, 0)
	root := &Node{
		Parent:             nil,
		Radius:             0.,
		DistanceFromParent: 0.,
		CentroidEntry:      nil,
		EntryList:          newEntryList,
	}

	newTree := &Tree{
		Root:           root,
		MaxEntry:       maxEntry,
		SplitMechanism: splitMechanism,
		DistCalc:       distCalc,
	}

	return newTree
}
