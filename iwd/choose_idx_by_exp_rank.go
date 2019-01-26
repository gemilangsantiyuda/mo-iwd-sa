package iwd

import (
	"math/rand"
	"sort"
)

func chooseIdxByExpRank(valueList []float64, sp float64) int {
	var itemList []probItem
	for idx := range valueList {
		val := valueList[idx]
		item := probItem{
			Idx:  idx,
			Prob: val,
		}
		itemList = append(itemList, item)
	}
	sort.SliceStable(itemList, func(i, j int) bool {
		return itemList[i].Prob > itemList[j].Prob
	})

	spList := make([]float64, len(itemList)+1)
	spX := 1.
	for idx := range spList {
		spList[idx] = spX
		spX *= sp
	}

	r := rand.Float64()
	upper := 0.
	chosenItemIdx := -1
	// renewing the prob in each item
	for idx := range itemList {
		item := itemList[idx]
		item.Prob = spList[idx] * (1. - sp) / (1. - spList[len(itemList)])
		// fmt.Println(item.Prob, r)
		upper += item.Prob
		if r <= upper {
			chosenItemIdx = idx
			break
		}
	}

	chosenIdx := itemList[chosenItemIdx].Idx
	return chosenIdx
}
