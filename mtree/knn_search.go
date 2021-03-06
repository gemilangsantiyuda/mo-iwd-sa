package mtree

import (
	"container/heap"

	"github.com/vroup/mo-iwd-sa/order"
)

// An Item is something we manage in a priority queue.
type Item struct {
	entry    entry
	distance float64
	index    int // The index of the item in the heap. essential for heap
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

type neighbour struct {
	Object   Object
	Distance float64
}

// KnnSearch return k entry nearest to query object
func (tree *Tree) KnnSearch(queryObject Object, k int, maxDistance float64, maxQty int) []*neighbour {
	var pq PriorityQueue
	var neighbourList []*neighbour
	heap.Init(&pq)

	item := &Item{
		entry: tree.root,
	}
	heap.Push(&pq, item)

	// start best first search, if item is leafEntry (an actual object), then append it to neighbourList, else if it's node then push it to priority queue with distance equals to queryobject distance to the node's circle edge
	for pq.Len() > 0 && len(neighbourList) < k {
		item = heap.Pop(&pq).(*Item)
		entry := item.entry
		distance := item.distance

		if _, ok := entry.(*leafEntry); ok {
			object := entry.getCentroidObject().(*order.Order)
			if object.Quantity > maxQty {
				continue
			}
			newNeighbour := &neighbour{
				Object:   object,
				Distance: distance,
			}
			neighbourList = append(neighbourList, newNeighbour)
			continue
		}

		distCalc := tree.distCalc
		node := entry.(node)
		entryList := node.getEntryList()
		for idx := range entryList {
			nextEntry := entryList[idx]
			distanceToNextEntry := distCalc.GetDistance(queryObject, nextEntry.getCentroidObject()) - nextEntry.getRadius()
			if distanceToNextEntry < 0 {
				distanceToNextEntry = 0
			}
			if distanceToNextEntry > maxDistance {
				continue
			}
			item := &Item{
				entry:    nextEntry,
				distance: distanceToNextEntry,
			}
			heap.Push(&pq, item)
		}
	}
	return neighbourList
}
