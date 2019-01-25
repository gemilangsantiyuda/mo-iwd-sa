package mtree

import "github.com/vroup/mo-iwd-sa/distance"

// NewTree initiate new tree
func NewTree(maxEntry int, splitMechanism SplitMechanism, distCalc DistanceCalculator) *Tree {

	if distCalc == nil {
		distCalc = &distance.ManhattanDistance{}
	}

	if splitMechanism == nil {
		splitMechanism = &SplitMST{
			MaxEntry: maxEntry,
			DistCalc: distCalc,
		}
	}

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
