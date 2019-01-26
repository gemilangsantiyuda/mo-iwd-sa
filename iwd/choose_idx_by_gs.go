package iwd

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type probGroup struct {
	Group []probItem
	Prob  float64
}

func chooseIdxByGS(valueList []float64, m int) int {

	const goldenRatio = 0.638
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

	groupCountList := make([]int, m)
	itemCount := len(valueList)
	for group := range groupCountList {
		if group == m-1 {
			groupCountList[m-group-1] = itemCount
		}
		count := int(math.Round(goldenRatio * float64(itemCount)))
		groupCountList[m-group-1] = count
		itemCount -= count
	}
	fmt.Println(groupCountList)
	groupList := make([]probGroup, m)

	itemIdx := 0
	prob := 1.
	for groupIdx := range groupCountList {
		newGroup := make([]probItem, groupCountList[groupIdx])
		for itemCount := range newGroup {
			newGroup[itemCount] = itemList[itemIdx]
			itemIdx++
		}
		groupList[groupIdx].Group = newGroup
		if groupIdx == len(groupCountList)-1 {
			groupList[groupIdx].Prob = prob
		} else {
			groupList[groupIdx].Prob = prob * goldenRatio
		}

		prob -= prob * goldenRatio
	}

	fmt.Println(groupList)
	var chosenGroup probGroup
	upper := 0.
	r := rand.Float64()
	for groupIdx := range groupList {
		group := groupList[groupIdx]
		upper += group.Prob
		fmt.Print(upper, " ")
		if r <= upper {
			chosenGroup = group
			break
		}
	}
	fmt.Println("upper")
	fmt.Println(chosenGroup)
	chosenItemIdx := rand.Intn(len(chosenGroup.Group))
	chosenItem := chosenGroup.Group[chosenItemIdx]
	chosenIdx := chosenItem.Idx
	return chosenIdx
}
