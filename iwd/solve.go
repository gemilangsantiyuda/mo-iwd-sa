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

<<<<<<< HEAD
=======
	bestArchive := &Archive{
		ElementList: make([]*ArchiveElement, 0),
	}

>>>>>>> iwd-sa with SPEA 2
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

<<<<<<< HEAD
=======
		localArchive := &Archive{
			ElementList: make([]*ArchiveElement, 0),
		}
>>>>>>> iwd-sa with SPEA 2
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
<<<<<<< HEAD
					if wd.Score.IsWorseThan(mutationWD.Score) {
						wd = mutationWD
					} else {
						fmt.Print(iter)
=======
					if mutationWD.Score.IsDominate(wd.Score, config.Tolerance) {
						wd = mutationWD
					} else {
>>>>>>> iwd-sa with SPEA 2
						prob := getSAProb(mutationWD.Score, wd.Score, config)
						r := rand.Float64()
						if r <= prob {
							wd = mutationWD
						}
					}

				}

<<<<<<< HEAD
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

=======
				// if wd.Score.IsDominate(localBestScore, config.Tolerance) {
				// 	localBestScore = wd.Score
				// 	localBestWD = wd
				// }
				archiveE := &ArchiveElement{
					Wd: wd,
				}
				localArchive.ElementList = append(localArchive.ElementList, archiveE)
			} //else {
			// 	// the iwd create invalid routelist,, then restore the updated soil parameter
			// 	wd.restoreRemovedSoil()
			// }
		}

		localArchive.Update(config.IwdParameter.PopulationSize)
		fmt.Println(len(localArchive.ElementList))
		if len(localArchive.ElementList) > 0 {
			localBestWD = localArchive.ElementList[0].Wd
			localBestScore = localBestWD.Score
		} else {
			localBestWD = nil
		}
>>>>>>> iwd-sa with SPEA 2
		if localBestWD != nil {
			globalUpdate(soilMap, localBestWD)
		}

<<<<<<< HEAD
		if bestScore.IsWorseThan(localBestScore) {
=======
		bestArchive.ElementList = append(bestArchive.ElementList, localArchive.ElementList...)
		bestArchive.Update(config.ArchiveSize)

		if localBestScore.IsDominate(bestScore, config.Tolerance) {
>>>>>>> iwd-sa with SPEA 2
			bestScore = localBestScore
			bestWD = localBestWD
		}
		// if best score found then renew bestwd
<<<<<<< HEAD
		fmt.Printf("Local Best Score %+v \n", localBestScore)
		fmt.Printf("Best Score %+v \n", bestScore)

=======
		// fmt.Printf("Local Best Score %+v \n", localBestScore)
		// fmt.Printf("Best Score %+v \n", bestScore)
		fmt.Println("Best Archive Iteration ", iter)
		for arIdx := range bestArchive.ElementList {
			fmt.Printf("%+v\n", bestArchive.ElementList[arIdx].Wd.Score)
		}
		updateTemperature(config)
>>>>>>> iwd-sa with SPEA 2
	}

	return bestWD
}
