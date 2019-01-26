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
	servedQty := currentRoute.ServedQty
	servingKitchen := currentRoute.ServingKitchen

	var moDistanceList []float64
	var soilList []float64
	distance := 0.
	kitchenOpt := 0.
	userRating := 0.

	neighbourList := tree.KnnSearch(tree.Root, currentNode, conf.NeighbourCount, maxCap, maxDistance)
	if len(neighbourList) == 0 {
		// no neighbour found feasible
		return nil, 0, 0
	}
	// fmt.Printf("CurrentRoute---->%+v\n", currentRoute)
	for idx := range neighbourList {
		neighbour := neighbourList[idx]
		order := neighbour.Order
		distance = neighbour.Distance
		newQty := servedQty + order.Quantity
		if newQty <= servingKitchen.Capacity.Optimum {
			difQty := math.Max(float64(servingKitchen.Capacity.Optimum-newQty), float64(conf.MaxDriverCapacity))
			kitchenOpt = (difQty / float64(conf.MaxDriverCapacity)) * 0.8
		} else {
			difQty := math.Max(float64(newQty-servingKitchen.Capacity.Optimum), float64(conf.MaxDriverCapacity))
			kitchenOpt = difQty / float64(conf.MaxDriverCapacity)
		}
		userRating = wd.RatingMap.GetOrderToKitchenRating(order, servingKitchen) / 5.
		moDistance := calcMoDistance(distance/wd.Config.MaxDriverDistance, kitchenOpt, userRating, &conf.Weight)
		moDistanceList = append(moDistanceList, moDistance)

		soil := soilMap.GetSoil(currentNode, order)
		soilList = append(soilList, soil)
	}
	probList := wd.getProbList(soilList)
	chosenIdx := chooseIdxByRouletteWheel(probList)
	// fmt.Println("Chosen Idx", chosenIdx)
	nextOrder, distance, moDistance := neighbourList[chosenIdx].Order, neighbourList[chosenIdx].Distance, moDistanceList[chosenIdx]

	return nextOrder, distance, moDistance
}
