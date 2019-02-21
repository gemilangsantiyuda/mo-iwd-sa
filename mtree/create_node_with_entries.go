package mtree

func createNodeWithEntries(entryList []entry, distCalc distanceCalculator) node {
	objList := make([]Object, 0)
	for idx := range entryList {
		obj := entryList[idx].getCentroidObject()
		objList = append(objList, obj)
	}
	distanceMatrix := getDistanceMatrix(objList, distCalc)

	centroidIdx, radius := chooseBestCentroidIdx(entryList, distanceMatrix)

	var newNode node

	// checking if the entrylist is of leaf entries, then we have to return leaf nodes, else we return branch nodes
	if _, ok := entryList[0].(*leafEntry); ok {
		newNode = &leaf{
			centroidObject: objList[centroidIdx],
			radius:         radius,
			entryList:      entryList,
		}
	} else {
		newNode = &branch{
			centroidObject: objList[centroidIdx],
			radius:         radius,
			entryList:      entryList,
		}
	}

	// update the entries distance to parent and the parent
	for idx := range entryList {
		entryList[idx].setParent(newNode)
		distanceFromParent := distanceMatrix[centroidIdx][idx]
		entryList[idx].setDistanceFromParent(distanceFromParent)
	}

	return newNode
}
