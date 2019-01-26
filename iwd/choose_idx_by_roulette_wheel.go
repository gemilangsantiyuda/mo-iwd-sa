package iwd

import (
	"math/rand"
	"sort"
)

type probItem struct {
	Idx  int
	Prob float64
}

func chooseIdxByRouletteWheel(valueList []float64) int {
	totalValue := 0.
	for idx := range valueList {
		val := valueList[idx]
		totalValue += val
	}

	var itemList []probItem
	for idx := range valueList {
		val := valueList[idx]
		item := probItem{
			Idx:  idx,
			Prob: val / totalValue,
		}
		itemList = append(itemList, item)
	}

	sort.SliceStable(itemList, func(i, j int) bool {
		return itemList[i].Prob > itemList[j].Prob
	})

	randomVal := rand.Float64()
	upper := 0.

	for idx := range itemList {
		item := itemList[idx]
		upper += item.Prob
		if randomVal <= upper {
			return item.Idx
		}
	}

	return -1
}
