package mtree

import "math"

func chooseBestCentroidIdx(entryList []entry, distanceMatrix [][]float64) (int, float64) {

	bestCentroidIdx := -1
	bestRadius := math.Inf(1)
	for idx1 := range entryList {
		maxRadius := math.Inf(-1)
		for idx2 := range entryList {
			entry2 := entryList[idx2]
			radius := distanceMatrix[idx1][idx2] + entry2.getRadius()
			maxRadius = math.Max(maxRadius, radius)
		}
		if maxRadius < bestRadius {
			bestCentroidIdx = idx1
			bestRadius = maxRadius
		}
	}
	return bestCentroidIdx, bestRadius
}
