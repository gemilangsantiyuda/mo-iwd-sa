package mtree

import (
	"math"

	"github.com/vroup/mo-iwd-sa/object"
)

// Remove an object having objectID from the tree
func (tree *Tree) Remove(node *Node, object object.Object, objectID string) (removeSuccess bool) {

	// Check all covered node and traverse them until removeSuccess
	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		if _, isLeafEntry := entry.(*LeafEntry); isLeafEntry {
			lEntry := entry.(*LeafEntry)
			if lEntry.ObjectID == objectID {
				node.EntryList = append(node.EntryList[:idx], node.EntryList[idx+1:]...)
				if entry == node.CentroidEntry {
					node.SetCentroidEntry(nil)
				}
				entry.SetParent(nil)
				tree.repairNode(node)
				return true
			}
			continue
		}

		nextNode := node.EntryList[idx].(*Node)
		distanceToNextNode := tree.DistCalc.GetDistance(object, nextNode)
		if distanceToNextNode <= node.GetRadius() {
			removed := tree.Remove(nextNode, object, objectID)
			if removed {
				if node == tree.Root {
					tree.ObjectCount--
				}
				return true
			}
		}
	}
	return false
}

// repairNode starts to repair the current node, about radius changes which also affects its parent radius and especially about changes on its centroid entries which affects all other entries' distance to parent. then after that it repairs up to the root
func (tree *Tree) repairNode(node *Node) {
	centroidIsChanged := false
	if node.CentroidEntry == nil {
		// if node is empty and it is root then let it be, meaning tree is empty,  just set its radius to 0. else remove it from its parent and repair parent.
		if len(node.EntryList) == 0 {
			parent := node.GetParent()
			if parent == nil {
				node.SetRadius(0)
			} else {
				parent.RemoveEntry(node)
				tree.repairNode(parent)
			}
			return
		}
		// else if it isn't empty then centrodIsChanged =true, and remake the radius also update the entries distanceToParent
		tree.resetCentroid(node)
		centroidIsChanged = true
	} else {
		// if the centroid is not removed, then we just have to update the radius of this node
		tree.updateNodeRadius(node)
	}

	// Now we repair the parent node upto the root
	parent := node.GetParent()
	for parent != nil {
		// If the node is the parent's centroid entry then , if the node's centroid previously changed with chance of change of coordinate of its centroid also (centroidIsChanged) then we have to reset all parents' entries distance to parent and update parent's radius.  else we simply have to compare the node's radius and parent's radius set the larger to become parent's radius
		if node == parent.CentroidEntry {
			if centroidIsChanged {
				tree.resetEntriesDistanceToParent(parent)
				tree.updateNodeRadius(parent)
			} else {
				maxRad := math.Max(node.Radius, parent.Radius)
				parent.SetRadius(maxRad)
			}
		} else {
			// if it is not the parent's centroid entry, then if the centroidIsChanged (node's centroid is previously changed) then we have to reset node's distanceFromParent and compare its radius to parent's radius, also switch the centroidIsChanged to false (because parent's centroid is not changed by this repairment).. else just compare the radius
			if centroidIsChanged {
				distance := tree.DistCalc.GetDistance(node, parent)
				node.SetDistanceFromParent(distance)
				maxRad := math.Max(node.Radius+node.GetDistanceFromParent(), parent.Radius)
				parent.SetRadius(maxRad)
				centroidIsChanged = false
			} else {
				maxRad := math.Max(node.GetRadius()+node.GetDistanceFromParent(), parent.Radius)
				parent.SetRadius(maxRad)
			}
		}
		node = parent
		parent = node.GetParent()
	}
}

func (tree *Tree) resetEntriesDistanceToParent(node *Node) {
	entryList := node.EntryList
	for idx := range entryList {
		entry := entryList[idx]
		distanceFromParent := tree.DistCalc.GetDistance(node, entry)
		entry.SetDistanceFromParent(distanceFromParent)
	}
}

func (tree *Tree) resetCentroid(node *Node) {
	entryList := node.EntryList
	el := EntryList(entryList)
	distanceMatrix := tree.DistCalc.GetDistanceMatrix(el)
	bestRadius := math.Inf(1)
	var centroidEntryIdx int
	for idx1 := range entryList {
		maxRadius := math.Inf(-1)
		for idx2 := range entryList {
			entry2 := entryList[idx2]
			distance := distanceMatrix[idx1][idx2]
			radius := distance + entry2.GetRadius()
			maxRadius = math.Max(maxRadius, radius)
		}
		if maxRadius < bestRadius {
			bestRadius = maxRadius
			centroidEntryIdx = idx1
		}
	}

	// centroid entry is decided, then make new node with the centroid entry and set all entries to this node accordingly
	centroidEntry := entryList[centroidEntryIdx]
	node.SetCentroidEntry(centroidEntry)
	node.SetRadius(bestRadius)
	for idx := range entryList {
		entryList[idx].SetParent(node)
		distance := distanceMatrix[centroidEntryIdx][idx]
		entryList[idx].SetDistanceFromParent(distance)
	}
}

func (tree *Tree) updateNodeRadius(node *Node) {
	newRadius := math.Inf(-1)
	for idx := range node.EntryList {
		entry := node.EntryList[idx]
		radius := entry.GetDistanceFromParent() + entry.GetRadius()
		newRadius = math.Max(newRadius, radius)
	}
	node.SetRadius(newRadius)
}
