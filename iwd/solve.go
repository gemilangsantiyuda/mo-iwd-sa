package iwd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

// Solve the mdovrp returning the best waterdrop
func Solve(orderList []*order.Order, kitchenList []*kitchen.Kitchen, ratingMap rating.Map, tree *mtree.Tree, distCalc distanceCalculator, referencePoint []*Score, startTime time.Time, config *config.Config) []*WaterDrop {

	bestArchive := &Archive{
		ElementList: make([]*ArchiveElement, 0),
	}

	soilMap := NewSoilMap(kitchenList, orderList, config)
	// oldIGD := math.Inf(1)
	// igdStaticCount := 0
	for iter := 0; iter < config.IwdParameter.MaximumIteration; iter++ {
		relapsedTime := time.Since(startTime)
		if relapsedTime.Seconds() > 2.*3600. /*|| igdStaticCount == 50 */ {
			break
		}
		finishedWaterDrops := make([]*WaterDrop, 0)
		workingWaterDrops := initWaterDrops(soilMap, ratingMap, orderList, kitchenList, tree, distCalc, config)

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
				}

				archiveE := &ArchiveElement{
					Wd: wd,
				}
				localArchive.ElementList = append(localArchive.ElementList, archiveE)
			}
		}

		// optimize with local search all in the local archive

		for arIdx := range localArchive.ElementList {
			optimizeInverse(localArchive.ElementList[arIdx].Wd, distCalc, config)
			if localArchive.ElementList[arIdx].Wd.hasValidRouteList() {
				newElement := localArchive.ElementList[arIdx]
				bestArchive.ElementList = append(bestArchive.ElementList, newElement)
			}
		}
		logSolution(config, localArchive)
		localArchive.Update(config.LocalArchiveSize)
		bestArchive.Update(config.BestArchiveSize)
		// fmt.Println(len(localArchive.ElementList))
		for arIdx := range localArchive.ElementList {
			element := localArchive.ElementList[arIdx]
			wd := element.Wd
			fitness := element.Fitness
			if fitness >= 1 {
				prob := getSAProb(config, fitness)
				r := rand.Float64()
				if r <= prob {
					globalUpdate(soilMap, wd)
				}
			}
		}

		fmt.Println("Iteration ", iter, "Time ", relapsedTime)
		for arIdx := range bestArchive.ElementList {
			element := bestArchive.ElementList[arIdx]
			// fmt.Printf("%+v fit: %f\n", element.Wd.Score, element.Fitness)
			ksqMap := element.Wd.KitchenServedQtyMap
			totalServedQty := 0
			for kdx := range kitchenList {
				svKitchen := kitchenList[kdx]
				totalServedQty += ksqMap.GetServedQty(svKitchen)
			}
			// fmt.Println("total Served : ", totalServedQty)
		}
		updateTemperature(config)
		igdValue := bestArchive.getIGD(referencePoint, config)
		// if math.Abs(igdValue-oldIGD) < 0.00001 {
		// 	igdStaticCount++
		// } else {
		// 	igdStaticCount = 0
		// }
		// oldIGD = igdValue
		fmt.Println("IGD VALUE :", igdValue)
	}

	solutionList := make([]*WaterDrop, 0)
	for idx := range bestArchive.ElementList {
		solutionList = append(solutionList, bestArchive.ElementList[idx].Wd)
	}

	return solutionList
}
