package iwd

import (
	"fmt"
	"math"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

// Solve the mdovrp returning the best waterdrop
func Solve(orderList []*order.Order, kitchenList []*kitchen.Kitchen, ratingMap rating.Map, tree *mtree.Tree, config *config.Config) *WaterDrop {

	soilMap := NewSoilMap(kitchenList, orderList, config)
	var bestWD *WaterDrop
	var bestScore = &Score{
		RiderCost:         99999999,
		KitchenOptimality: 99999999,
		UserSatisfaction:  math.Inf(1),
	}

	for iter := 0; iter < config.IwdParameter.MaximumIteration; iter++ {
		finishedWaterDrops := make([]*WaterDrop, 0)
		workingWaterDrops := initWaterDrops(soilMap, ratingMap, orderList, kitchenList, tree, config)

		var localBestWD *WaterDrop
		var localBestScore = &Score{
			RiderCost:         99999999,
			KitchenOptimality: 99999999,
			UserSatisfaction:  math.Inf(1),
		}

		// As long as there is a water drop not finished traversing, then traverse the unfinisheds 1 by 1 , 1 step at a time
		for len(workingWaterDrops) > 0 {
			workingWaterDrops, finishedWaterDrops = nextStepWD(workingWaterDrops, finishedWaterDrops)
		}

		// calculate score for all finishedWaterDrops
		for wdIdx := range finishedWaterDrops {
			wd := finishedWaterDrops[wdIdx]
			wd.updateKitchenPreference()
			if wd.hasValidRouteList() {
				wd.calcScore()
				if localBestScore.IsWorseThan(wd.Score) {
					localBestScore = wd.Score
					localBestWD = wd
				}
			}
		}

		if localBestWD != nil {
			globalUpdate(soilMap, localBestWD)
		}

		if bestScore.IsWorseThan(localBestScore) {
			bestScore = localBestScore
			bestWD = localBestWD
		}
		// if best score found then renew bestwd
		fmt.Printf("Local Best Score %+v \n", localBestScore)
		fmt.Printf("Best Score %+v \n", bestScore)

	}

	return bestWD
}

func globalUpdate(soilMap SoilMap, wd *WaterDrop) {
	iwdParam := wd.Config.IwdParameter
	orderCount := float64(len(wd.OrderList))
	var currentNode, nextNode node
	for rIdx := range wd.RouteList {
		route := wd.RouteList[rIdx]
		currentNode = route.ServingKitchen
		for idx := range route.VisitedOrderList {
			nextNode = route.VisitedOrderList[idx]
			soil := soilMap.GetSoil(currentNode, nextNode)
			wdSoil := wd.Soil
			newSoil := (1-iwdParam.P)*soil - 2*iwdParam.P*wdSoil/(orderCount*(orderCount-1))
			soilMap.UpdateSoil(currentNode, nextNode, newSoil)
			// fmt.Println("edge ", currentNode.GetID(), " ", nextNode.GetID(), " soil ", soil, " new soil ", newSoil)
			currentNode = nextNode
		}
	}
}
