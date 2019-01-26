package iwd

import (
	"math"
)

func (wd *WaterDrop) getProbList(soilList []float64) []float64 {

	// fmt.Println("Soil List\n", soilList)

	minSoil := math.Inf(1)
	for idx := range soilList {
		minSoil = math.Min(soilList[idx], minSoil)
	}

	var gSoilList []float64
	for idx := range soilList {
		gSoil := soilList[idx]
		if minSoil < 0 {
			gSoil -= minSoil
		}
		gSoilList = append(gSoilList, gSoil)
	}

	var fSoilList []float64
	totalfSoil := 0.
	for idx := range gSoilList {
		fSoil := 1 / (0.0001 + gSoilList[idx])
		fSoilList = append(fSoilList, fSoil)
		totalfSoil += fSoil
	}

	var problist []float64
	for idx := range fSoilList {
		prob := fSoilList[idx] / totalfSoil
		problist = append(problist, prob)
	}

	// fmt.Println("prob list\n", problist)
	return problist
}
