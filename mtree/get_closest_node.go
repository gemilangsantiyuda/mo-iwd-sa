package mtree

import (
	"math"
)

func getClosestNode(currentNode node, entryList []entry, distCalc distanceCalculator) node {
	var closestNode node
	closestDistance := math.Inf(1)

	for idx := range entryList {
		nextNode := entryList[idx].(node)
		if nextNode == currentNode {
			continue
		}
		dist := distCalc.GetDistance(currentNode.getCentroidObject(), nextNode.getCentroidObject())
		if dist < closestDistance {
			closestDistance = dist
			closestNode = nextNode
		}
	}
	return closestNode
}
