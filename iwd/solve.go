package iwd

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

// Solve the mdovrp returning the best waterdrop
func Solve(rep int, orderList []*order.Order, kitchenList []*kitchen.Kitchen, ratingMap rating.Map, tree *mtree.Tree, distCalc distanceCalculator, referencePoint []*Score, startTime time.Time, config *config.Config, wg *sync.WaitGroup) []*WaterDrop {

	// defer wg.Done()
	bestArchive := &Archive{
		ElementList: make([]*ArchiveElement, 0),
	}
	// temp0 := config.SaParam.Temp
	// report := ""
	soilMap := NewSoilMap(kitchenList, orderList, config)
	oldIGD := math.Inf(1)
	igdStaticCount := 0
	iter := 0
	meantime := 0.
	var relapsedTime time.Duration
	for iter = 0; iter < 99999999; iter++ {
		relapsedTime = time.Since(startTime)
		if relapsedTime.Seconds() > 12000. || igdStaticCount == 50 {
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
		var wg2 sync.WaitGroup
		for arIdx := range localArchive.ElementList {
			wg2.Add(1)
			go optimizeInverse(localArchive.ElementList[arIdx].Wd, distCalc, config, &wg2)
			localArchive.ElementList[arIdx].Wd.calcScore()
			if localArchive.ElementList[arIdx].Wd.hasValidRouteList() {
				newElement := localArchive.ElementList[arIdx]
				bestArchive.ElementList = append(bestArchive.ElementList, newElement)
			}
		}

		wg2.Wait()
		// logSolution(config, localArchive)
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

		// fmt.Println("Iteration ", iter, "Time ", relapsedTime)
		for arIdx := range bestArchive.ElementList {
			element := bestArchive.ElementList[arIdx]
			fmt.Printf("%+v fit: %f\n", element.Wd.Score, element.Fitness)
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
		if math.Abs(igdValue-oldIGD) < 0.00001 {
			igdStaticCount++
		} else {
			igdStaticCount = 0
		}
		oldIGD = igdValue
		meantime += relapsedTime.Seconds()
		fmt.Println("IGD VALUE :", igdValue)
		// fmt.Printf("%s,%d,%f,%f\n", config.DataSize, rep, relapsedTime.Seconds(), igdValue)
	}
	// meantime /= 5
	// fmt.Println(config.DataSize, rep, temp0, config.SaParam.CoolingRate, meantime, oldIGD)
	// fmt.Print(report)
	// fmt.Println(config.LocalArchiveSize, relapsedTime.Seconds(), igdValue)
	solutionList := make([]*WaterDrop, 0)
	for idx := range bestArchive.ElementList {
		solutionList = append(solutionList, bestArchive.ElementList[idx].Wd)
	}

	return solutionList
}
