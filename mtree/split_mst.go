package mtree

import (
	"sort"
)

// splitMST split MST implement splitting with MST from slim tree
type splitMST struct {
	distCalc distanceCalculator
}

func (sm *splitMST) split(entryList []entry) (node, node) {
	var newNode1, newNode2 node
	entryList1, entryList2 := sm.partitionEntryList(entryList)

	newNode1 = createNodeWithEntries(entryList1, sm.distCalc)
	newNode2 = createNodeWithEntries(entryList2, sm.distCalc)

	return newNode1, newNode2
}

type edge struct {
	EntryIdx1, EntryIdx2 int
	Length               float64
}

type adjacency []int

func (sm *splitMST) partitionEntryList(entryList []entry) ([]entry, []entry) {
	edgeList := sm.getMSTEdgeList(entryList)
	// Start partitioning, make 2 subgraph out of the edgelist minus the longest edge
	var newEntryList1, newEntryList2 []entry
	parentList := make([]int, len(entryList))
	for idx := range parentList {
		parentList[idx] = idx
	}

	for idx := 0; idx < len(edgeList)-1; idx++ {
		edge := edgeList[idx]
		entryIdx1, entryIdx2 := edge.EntryIdx1, edge.EntryIdx2
		parent1 := parentList[entryIdx1]
		parent2 := parentList[entryIdx2]
		if parent1 != parent2 {
			parentList[parent2] = parent1
		}
	}

	subgraphIdx1 := sm.getParent(0, parentList)
	for idx := range entryList {
		entry := entryList[idx]
		if sm.getParent(idx, parentList) == subgraphIdx1 {
			newEntryList1 = append(newEntryList1, entry)
		} else {
			newEntryList2 = append(newEntryList2, entry)
		}
	}

	return newEntryList1, newEntryList2
}

func (sm *splitMST) getMSTEdgeList(entryList []entry) []*edge {

	objList := make([]Object, 0)
	for idx := range entryList {
		obj := entryList[idx].getCentroidObject()
		objList = append(objList, obj)
	}
	distanceMatrix := getDistanceMatrix(objList, sm.distCalc)

	edgeList := sm.getSortedEdgeList(distanceMatrix)
	var mstEdgeList []*edge
	// MST Start
	parentList := make([]int, len(entryList))
	for idx := range parentList {
		parentList[idx] = idx
	}
	for idx := range edgeList {
		edge := edgeList[idx]
		entryIdx1, entryIdx2 := edge.EntryIdx1, edge.EntryIdx2
		parent1 := sm.getParent(entryIdx1, parentList)
		parent2 := sm.getParent(entryIdx2, parentList)
		if parent1 != parent2 {
			parentList[parent2] = parent1
			mstEdgeList = append(mstEdgeList, edge)
		}
	}

	return mstEdgeList
}

func (sm *splitMST) getSortedEdgeList(distanceMatrix [][]float64) []*edge {
	var sortedEdgeList []*edge
	for idx1 := range distanceMatrix {
		for idx2 := range distanceMatrix[idx1] {
			if idx1 != idx2 {
				edge := &edge{
					EntryIdx1: idx1,
					EntryIdx2: idx2,
					Length:    distanceMatrix[idx1][idx2],
				}
				sortedEdgeList = append(sortedEdgeList, edge)
			}
		}
	}
	sort.SliceStable(sortedEdgeList, func(i, j int) bool {
		return sortedEdgeList[i].Length < sortedEdgeList[j].Length
	})
	return sortedEdgeList
}

func (sm *splitMST) getParent(idx int, parentList []int) int {
	parentIdx := parentList[idx]
	for parentIdx != parentList[parentIdx] {
		parentIdx = parentList[parentIdx]
	}
	return parentIdx
}
