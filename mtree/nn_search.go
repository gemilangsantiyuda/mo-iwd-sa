package mtree

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/vroup/mo-iwd-sa/object"
)

// An Item is something we manage in a priority queue.
type Item struct {
	entry    Entry
	distance float64
	index    int // The index of the item in the heap. essential for heap
}

// Neighbour entry and its distance to query object
type Neighbour struct {
	LeafEntry *LeafEntry
	Distance  float64
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push push the item to the priority queue,to be consumed by heap
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop function to be consumed by heap, to pop smallest element based on Less()
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// KnnSearch return k entry nearest to query object
func (tree *Tree) KnnSearch(root *Node, queryObject object.Object, k int) []*Neighbour {
	var pq PriorityQueue
	var neighbourList []*Neighbour
	heap.Init(&pq)

	item := &Item{
		entry: root,
	}
	heap.Push(&pq, item)

	// start best first search, if item is leafEntry (an actual object), then append it to neighbourList, else if it's node then push it to priority queue with distance equals to queryobject distance to the node's circle edge
	for pq.Len() > 0 && len(neighbourList) < k {
		item = heap.Pop(&pq).(*Item)
		entry := item.entry
		fmt.Println(entry)
		if _, ok := entry.(*LeafEntry); ok {
			leafEntry := entry.(*LeafEntry)
			distance := item.distance
			neighbour := &Neighbour{
				LeafEntry: leafEntry,
				Distance:  distance,
			}
			neighbourList = append(neighbourList, neighbour)
			continue
		}

		node := entry.(*Node)
		for idx := range node.EntryList {
			nextEntry := node.EntryList[idx]
			distMin := math.Max(tree.DistCalc.GetDistance(queryObject, nextEntry)-nextEntry.GetRadius(), 0.)
			item := &Item{
				entry:    nextEntry,
				distance: distMin,
			}
			heap.Push(&pq, item)
		}
	}
	return neighbourList
}
