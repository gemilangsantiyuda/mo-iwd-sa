package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/order"
)

func (wd *WaterDrop) getNextOrder(currentRoute *Route, currentNode node) (*order.Order, float64, float64) {
	conf := wd.Config
	tree := wd.Tree
	soilMap := wd.SoilMap
	maxDistance := conf.MaxDriverDistance - currentRoute.DistanceTraveled
	maxCap := currentRoute.CapacityLeft
	neighbourList := tree.KnnSearch(tree.Root, currentNode, 5, maxCap, maxDistance)
	servedQty := currentRoute.ServedQty
	servingKitchen := currentRoute.ServingKitchen

	var moDistanceList []float64
	var soilList []float64
	distance := 0.
	kitchenOpt := 0.
	userRating := 0.

	for idx := range neighbourList {
		neighbour := neighbourList[idx]
		order := neighbour.Order
		distance = neighbour.Distance / 90000
		newQty := servedQty + order.Quantity
		if newQty <= servingKitchen.Capacity.Optimum {
			difQty := math.Max(float64(servingKitchen.Capacity.Optimum-newQty), float64(conf.MaxDriverCapacity))
			kitchenOpt = (difQty / float64(conf.MaxDriverCapacity)) * 0.8
		} else {
			difQty := math.Max(float64(newQty-servingKitchen.Capacity.Optimum), float64(conf.MaxDriverCapacity))
			kitchenOpt = difQty / float64(conf.MaxDriverCapacity)
		}
		userRating = wd.RatingMap.GetOrderToKitchenRating(order, servingKitchen) / 5.
		moDistance := calcMoDistance(distance, kitchenOpt, userRating, &conf.Weight)
		moDistanceList = append(moDistanceList, moDistance)

		soil := soilMap.GetSoil(currentNode, order)
		soilList = append(soilList, soil)
	}
	probList := wd.getProbList(soilList)
	chosenIdx := chooseIdxByRouletteWheel(probList)
	nextOrder, distance, moDistance := neighbourList[chosenIdx].Order, neighbourList[chosenIdx].Distance, moDistanceList[chosenIdx]

	return nextOrder, distance, moDistance
}

func (wd *WaterDrop) getProbList(soilList []float64) []float64 {

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
		gSoilList = append(gSoilList)
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

	return problist
}
