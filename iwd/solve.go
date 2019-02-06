package iwd

import (
	"fmt"
	"math/rand"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

// Solve the mdovrp returning the best waterdrop
func Solve(orderList []*order.Order, kitchenList []*kitchen.Kitchen, ratingMap rating.Map, tree *mtree.Tree, config *config.Config) *WaterDrop {

	bestArchive := &Archive{
		ElementList: make([]*ArchiveElement, 0),
	}

	soilMap := NewSoilMap(kitchenList, orderList, config)
	for iter := 0; iter < config.IwdParameter.MaximumIteration; iter++ {
		finishedWaterDrops := make([]*WaterDrop, 0)
		workingWaterDrops := initWaterDrops(soilMap, ratingMap, orderList, kitchenList, tree, config)

		localArchive := &Archive{
			ElementList: make([]*ArchiveElement, 0),
		}

		// As long as there is a water drop not finished traversing, then traverse the unfinisheds 1 by 1 , 1 step at a time
		for len(workingWaterDrops) > 0 {
			workingWaterDrops, finishedWaterDrops = nextStepWD(workingWaterDrops, finishedWaterDrops)
		}

		// calculate score for all finishedWaterDrops
		for wdIdx := range finishedWaterDrops {
			wd := finishedWaterDrops[wdIdx]
			// wd.updateKitchenPreference()
			if wd.hasValidRouteList() {
				wd.calcScore()
				// mutate water drop and replace based on SA
				swapCount := 1
				mutationWD := wd.getMutation(swapCount)
				if mutationWD.hasValidRouteList() {
					mutationWD.calcScore()
					archiveE := &ArchiveElement{
						Wd: mutationWD,
					}
					localArchive.ElementList = append(localArchive.ElementList, archiveE)
					// if mutationWD.Score.IsDominate(wd.Score, config.Tolerance) {
					// 	wd = mutationWD
					// } else {
					// 	prob := getSAProb(mutationWD.Score, wd.Score, config)
					// 	r := rand.Float64()
					// 	if r <= prob {
					// 		wd = mutationWD
					// 	}
					// }

				}

				archiveE := &ArchiveElement{
					Wd: wd,
				}
				localArchive.ElementList = append(localArchive.ElementList, archiveE)
			}
		}

		bestArchive.ElementList = append(bestArchive.ElementList, localArchive.ElementList...)
		bestArchive.Update(config.ArchiveSize)

		localArchive.Update(2)
		fmt.Println(len(localArchive.ElementList))
		for arIdx := range localArchive.ElementList {
			element := localArchive.ElementList[arIdx]
			wd := element.Wd
			fitness := element.Fitness
			if fitness > 1 {
				prob := getSAProb(config, fitness)
				r := rand.Float64()
				if r <= prob {
					globalUpdate(soilMap, wd)
				}
			}
		}

		fmt.Println("Best Archive Iteration ", iter)
		for arIdx := range bestArchive.ElementList {
			element := bestArchive.ElementList[arIdx]
			fmt.Printf("%+v fit: %f\n", element.Wd.Score, element.Fitness)
		}
		updateTemperature(config)
	}

	return bestArchive.ElementList[0].Wd
}
