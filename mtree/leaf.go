package mtree

import "math"

type leaf struct {
	parent             node
	radius             float64
	distanceFromParent float64
	centroidObject     Object
	entryList          []entry
}

func (le *leaf) isLeaf() bool {
	return true
}

func (le *leaf) getRadius() float64 {
	return le.radius
}

func (le *leaf) getDistanceFromParent() float64 {
	return le.distanceFromParent
}

func (le *leaf) setDistanceFromParent(distance float64) {
	le.distanceFromParent = distance
}

func (le *leaf) getCentroidObject() Object {
	return le.centroidObject
}

func (le *leaf) insertEntry(newEntry entry) {
	le.entryList = append(le.entryList, newEntry)
	newEntry.setParent(le)
	if le.centroidObject == nil {
		le.centroidObject = newEntry.getCentroidObject()
	}
}

func (le *leaf) updateRadius() {
	le.radius = math.Inf(-1)
	for idx := range le.entryList {
		radius := le.entryList[idx].getDistanceFromParent()
		le.radius = math.Max(le.radius, radius)
	}
}

func (le *leaf) isUnderFlown(minEntry int) bool {
	return len(le.entryList) < minEntry
}

func (le *leaf) getEntryList() []entry {
	return le.entryList
}

func (le *leaf) removeObject(object Object) {
	for idx := range le.entryList {
		obj := le.entryList[idx].getCentroidObject()
		if obj.GetID() == object.GetID() {
			le.entryList = append(le.entryList[:idx], le.entryList[idx+1:]...)
			break
		}
	}
}

func (le *leaf) getParent() node {
	return le.parent
}

func (le *leaf) setParent(parent node) {
	le.parent = parent
}

func (le *leaf) containsObject(searchedObject Object) bool {
	for idx := range le.entryList {
		obj := le.entryList[idx].getCentroidObject()
		if obj.GetID() == searchedObject.GetID() {
			return true
		}
	}
	return false
}
