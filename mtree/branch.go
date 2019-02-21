package mtree

import (
	"fmt"
	"log"
	"math"
)

type branch struct {
	parent             node
	radius             float64
	distanceFromParent float64
	centroidObject     Object
	entryList          []entry
}

func (br *branch) isLeaf() bool {
	return false
}

func (br *branch) getRadius() float64 {
	return br.radius
}

func (br *branch) getCentroidObject() Object {
	return br.centroidObject
}

func (br *branch) getDistanceFromParent() float64 {
	return br.distanceFromParent
}

func (br *branch) setDistanceFromParent(distance float64) {
	br.distanceFromParent = distance
}

func (br *branch) insertEntry(newEntry entry) {
	if br.centroidObject == nil {
		br.centroidObject = newEntry.getCentroidObject()
	}
	newEntry.setParent(br)
	br.entryList = append(br.entryList, newEntry)
}

func (br *branch) removeEntry(entryToRemove entry) {
	for idx := range br.entryList {
		entry := br.entryList[idx]
		if entry == entryToRemove {
			br.entryList = append(br.entryList[:idx], br.entryList[idx+1:]...)
			return
		}
	}
	log.Fatal(fmt.Sprintf("Error Removing %v\n", entryToRemove))
}

func (br *branch) getParent() node {
	return br.parent
}

func (br *branch) setParent(parent node) {
	br.parent = parent
}

func (br *branch) updateRadius() {
	br.radius = math.Inf(-1)
	for idx := range br.entryList {
		radius := br.entryList[idx].getDistanceFromParent() + br.entryList[idx].getRadius()
		br.radius = math.Max(br.radius, radius)
	}
}

func (br *branch) isUnderFlown(minEntry int) bool {
	return len(br.entryList) < minEntry
}

func (br *branch) getEntryList() []entry {
	return br.entryList
}
