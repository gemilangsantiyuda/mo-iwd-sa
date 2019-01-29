package iwd

import (
	"fmt"
	"math"
	"math/rand"

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

		// if iter < 50 {
		// 	config.NeighbourCount = 1
		// } else {
		// 	config.NeighbourCount = rand.Intn(4) + 2
		// }

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
				// mutate water drop and replace based on SA
				swapCount := 1
				mutationWD := wd.getMutation(swapCount)
				if mutationWD.hasValidRouteList() {
					mutationWD.calcScore()
					if wd.Score.IsWorseThan(mutationWD.Score) {
						wd = mutationWD
					} else {
						fmt.Print(iter)
						prob := getSAProb(mutationWD.Score, wd.Score, config)
						r := rand.Float64()
						if r <= prob {
							wd = mutationWD
						}
					}

				}

				if localBestScore.IsWorseThan(wd.Score) {
					localBestScore = wd.Score
					localBestWD = wd
				}
			} else {
				// the iwd create invalid routelist,, then restore the updated soil parameter
				wd.restoreRemovedSoil()
			}
			updateTemperature(config)
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
