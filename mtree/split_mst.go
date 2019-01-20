package mtree

import (
	"math"
	"sort"
)

// SplitMST based on the slim-tree MST Split mechanism
type SplitMST struct {
	MaxEntry int
	DistCalc DistanceCalculator
}

// Edge and Adjacency FOR MST making
type Edge struct {
	EntryIdx1, EntryIdx2 int
	Length               float64
}

// Adjacency is adjacency, God linting so bothersome sometimes
type Adjacency []int

// Split the full node into 2 new nodes with MST mechanism based on slim-tree
func (sm *SplitMST) Split(nodeToSplit *Node, newEntry Entry) *Node {
	entryList := append(nodeToSplit.EntryList, newEntry)
	newEntryList1, newEntryList2 := sm.partitionEntryList(entryList)

	// creating new nodes from the partitioned previous entrylist
	newNode1 := sm.createNewNodeWithExistingEntries(newEntryList1)
	newNode2 := sm.createNewNodeWithExistingEntries(newEntryList2)

	// check if nodeToSplit is not the root ,if it isn't then replace the nodeToSplit with newNode1 and insert newNode2. if parentNode has full entries, then split the parentNode
	if nodeToSplit.Parent != nil {
		parentNode := nodeToSplit.Parent
		sm.replaceNodeEntry(parentNode, nodeToSplit, newNode1)
		if len(parentNode.EntryList) == sm.MaxEntry {
			newRoot := sm.Split(parentNode, newNode2)
			return newRoot
		}
		// else
		parentNode.InsertEntry(newNode2)
		// Update parent radius and each entries distance to parent
		sm.updateNodeRadius(parentNode)
		return nil
	}

	// create new root
	var rootEntryList = []Entry{newNode1, newNode2}
	newRoot := sm.createNewNodeWithExistingEntries(rootEntryList)

	return newRoot
}

// ReplaceEntry to replace an entry of node. In case the replaced entry was the node's centroid, then the node entries need to update their distanceFromParent
func (sm *SplitMST) replaceNodeEntry(node *Node, entryToReplace, replacementEntry Entry) {
	for idx := range node.EntryList {
		if node.EntryList[idx] == entryToReplace {
			node.EntryList[idx] = replacementEntry
			break
		}
	}

	if node.CentroidEntry == entryToReplace {
		node.CentroidEntry = replacementEntry
		for idx := range node.EntryList {
			distance := sm.DistCalc.GetDistance(node, node.EntryList[idx])
			node.EntryList[idx].SetDistanceFromParent(distance)
		}
	} else {
		distance := sm.DistCalc.GetDistance(node, replacementEntry)
		replacementEntry.SetDistanceFromParent(distance)
	}
	replacementEntry.SetParent(node)
}

func (sm *SplitMST) updateNodeRadius(node *Node) {
	newRadius := 0.
	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		distanceNodeToEntry := sm.DistCalc.GetDistance(node, entry)
		entry.SetDistanceFromParent(distanceNodeToEntry)
		tempRadius := distanceNodeToEntry + entry.GetRadius()
		newRadius = math.Max(tempRadius, newRadius)
	}
	node.Radius = newRadius
}

func (sm *SplitMST) createNewNodeWithExistingEntries(entryList []Entry) *Node {
	el := EntryList(entryList)
	distanceMatrix := sm.DistCalc.GetDistanceMatrix(el)
	bestRadius := math.Inf(1)
	var centroidEntryIdx int
	for idx1 := range entryList {
		maxRadius := math.Inf(-1)
		for idx2 := range entryList {
			entry2 := entryList[idx2]
			radius := distanceMatrix[idx1][idx2] + entry2.GetRadius()
			maxRadius = math.Max(maxRadius, radius)
		}
		if maxRadius < bestRadius {
			bestRadius = maxRadius
			centroidEntryIdx = idx1
		}
	}

	// centroid entry is decided, then make new node with the centroid entry and set all entries to this node accordingly
	centroidEntry := entryList[centroidEntryIdx]
	node := &Node{
		CentroidEntry: centroidEntry,
		Radius:        bestRadius,
		EntryList:     entryList,
	}
	for idx := range entryList {
		entryList[idx].SetParent(node)
		distance := distanceMatrix[centroidEntryIdx][idx]
		entryList[idx].SetDistanceFromParent(distance)
	}

	return node
}

func (sm *SplitMST) partitionEntryList(entryList []Entry) ([]Entry, []Entry) {
	edgeList := sm.getMSTEdgeList(entryList)
	// Start partitioning, make 2 subgraph out of the edgelist minus the longest edge
	var newEntryList1, newEntryList2 []Entry
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

func (sm *SplitMST) getMSTEdgeList(entryList []Entry) []*Edge {
	el := EntryList(entryList)
	distanceMatrix := sm.DistCalc.GetDistanceMatrix(el)
	edgeList := sm.getSortedEdgeList(distanceMatrix)
	var mstEdgeList []*Edge
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

func (sm *SplitMST) getSortedEdgeList(distanceMatrix [][]float64) []*Edge {
	var sortedEdgeList []*Edge
	for idx1 := range distanceMatrix {
		for idx2 := range distanceMatrix[idx1] {
			if idx1 != idx2 {
				edge := &Edge{
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

func (sm *SplitMST) getParent(idx int, parentList []int) int {
	parentIdx := parentList[idx]
	for parentIdx != parentList[parentIdx] {
		parentIdx = parentList[parentIdx]
	}
	return parentIdx
}
