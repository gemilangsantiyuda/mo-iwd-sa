package mtree

func getDistanceMatrix(objectList []Object, distCalc distanceCalculator) [][]float64 {
	distMatrix := make([][]float64, len(objectList))
	for idx := range objectList {
		distList := make([]float64, len(objectList))
		distMatrix[idx] = distList
	}

	for idx1 := 0; idx1 < len(objectList); idx1++ {
		object1 := objectList[idx1]
		for idx2 := idx1 + 1; idx2 < len(objectList); idx2++ {
			object2 := objectList[idx2]
			dist := distCalc.GetDistance(object1, object2)
			distMatrix[idx1][idx2] = dist
			distMatrix[idx2][idx1] = dist
		}
	}

	return distMatrix
}
