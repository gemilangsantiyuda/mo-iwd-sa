package mtree

import (
	"github.com/vroup/mo-iwd-sa/coordinate"
)

// Node is the struct for the m-tree node
type Node struct {
	Parent             *Node
	Radius             float64
	DistanceFromParent float64
	CentroidEntry      Entry
	EntryList          []Entry
}

// GetCoordinate return this node centroid coordinate
func (node *Node) GetCoordinate() *coordinate.Coordinate {
	return node.CentroidEntry.GetCoordinate()
}

// GetParent return this node's parent
func (node *Node) GetParent() *Node {
	return node.Parent
}

// SetParent set new Parent to this node and update its distance from node to its new parent
func (node *Node) SetParent(newParent *Node) {
	node.Parent = newParent
}

// GetRadius return this node's radius
func (node *Node) GetRadius() float64 {
	return node.Radius
}

// SetRadius set node's radius to new radius
func (node *Node) SetRadius(newRadius float64) {
	node.Radius = newRadius
}

// GetDistanceFromParent return node's distance from its parent
func (node *Node) GetDistanceFromParent() float64 {
	return node.DistanceFromParent
}

// SetDistanceFromParent set this node's distance to its parent node
func (node *Node) SetDistanceFromParent(dist float64) {
	node.DistanceFromParent = dist
}

// InsertEntry insert a new entry to node entry list
func (node *Node) InsertEntry(entry Entry) {
	node.EntryList = append(node.EntryList, entry)
	if node.CentroidEntry == nil {
		node.CentroidEntry = entry
	}
	entry.SetParent(node)
}

// RemoveEntry remove an entry from node's entry list
func (node *Node) RemoveEntry(entryToRemove Entry) {
	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		if entry == entryToRemove {
			node.EntryList = append(node.EntryList[:idx], node.EntryList[idx+1:]...)
			if entry == node.CentroidEntry {
				node.SetCentroidEntry(nil)
			}
			return
		}
	}
	entryToRemove.SetParent(nil)
}

// SetCentroidEntry set an entry of node as its new centroid entry
func (node *Node) SetCentroidEntry(entry Entry) {
	node.CentroidEntry = entry
}

// IsLeaf check wether a node is leaf by checking if it has any leaf (no leaf = leafnode) or its entries are leaf entries (which means it is a leaf node)
func (node *Node) IsLeaf() bool {
	entryList := node.EntryList
	// check if node has node entry
	for idx := range entryList {
		_, isNodeEntry := entryList[idx].(*Node)
		if isNodeEntry {
			return false
		}
	}
	return true
}

// ContainsObjectID may only be called by leafNode
func (node *Node) ContainsObjectID(objectID int) bool {
	for idx := range node.EntryList {
		leafEntry := node.EntryList[idx].(*LeafEntry)
		if leafEntry.ObjectID == objectID {
			return true
		}
	}
	return false
}

// RemoveEntryWithObjectID may only be called by leafNode
func (node *Node) RemoveEntryWithObjectID(objectID int) {
	for idx := range node.EntryList {
		entry := node.EntryList[idx].(*LeafEntry)
		if entry.ObjectID == objectID {
			node.EntryList = append(node.EntryList[:idx], node.EntryList[idx+1:]...)
			if entry == node.CentroidEntry {
				node.SetCentroidEntry(nil)
			}
			entry.SetParent(nil)
			return
		}
	}
}

// GetCentroidIdx on the entrylist of node
func (node *Node) GetCentroidIdx() int {
	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		if entry == node.CentroidEntry {
			return idx
		}
	}
	return -1
}
