package mtree

import "math"

func chooseBestNextNode(entryList []entry, newLeafEntry *leafEntry, distCalc distanceCalculator) node {
	var closestCoveringNode, closestNode entry
	closestNodeDistance := math.Inf(1)
	closestCoveringNodeDistance := math.Inf(1)

	for idx := range entryList {
		radius := entryList[idx].getRadius()
		distance := distCalc.GetDistance(newLeafEntry.getCentroidObject(), entryList[idx].getCentroidObject())
		if radius >= distance {
			if distance < closestCoveringNodeDistance {
				closestCoveringNodeDistance = distance
				closestCoveringNode = entryList[idx]
			}
		}
		if distance < closestNodeDistance {
			closestNodeDistance = distance
			closestNode = entryList[idx]
		}
	}

	if closestCoveringNode != nil {
		return closestCoveringNode.(node)
	}
	return closestNode.(node)
}
