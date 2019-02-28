package iwd

import (
	"fmt"

	"github.com/vroup/mo-iwd-sa/config"
)

func optimizeInverse(iwd *WaterDrop, distCalc distanceCalculator, config *config.Config) {
	repetition := config.IwdParameter.MaxLSRepetition
	ksqMap := iwd.KitchenServedQtyMap
	ratingMap := iwd.RatingMap

	improvementFound := false
	for repetition > 0 {
		for idx := range iwd.RouteList {
			route := iwd.RouteList[idx]
			startIdx, endIdx := getImprovingReverseIndex(route, distCalc)
			if startIdx == -1 {
				continue
			} else {
				route.reverseOrders(startIdx, endIdx, distCalc, ksqMap, ratingMap, config)
				repetition--
				improvementFound = true
				fmt.Println(repetition)
			}
			if repetition == 0 {
				break
			}
		}
		if !improvementFound {
			break
		}
		improvementFound = false
	}
}

func getImprovingReverseIndex(route *Route, distCalc distanceCalculator) (int, int) {
	routeLen := len(route.VisitedOrderList)

	for idx1 := 0; idx1 < routeLen-1; idx1++ {
		for idx2 := idx1 + 1; idx2 < routeLen; idx2++ {
			oldDistance := route.DistanceList[idx1]
			if idx1 > 0 {
				oldDistance -= route.DistanceList[idx1-1]
			}
			if idx2 < routeLen-1 {
				oldDistance += route.DistanceList[idx2+1]
				oldDistance -= route.DistanceList[idx2]
			}

			var lastNode node = route.ServingKitchen
			if idx1 > 0 {
				lastNode = route.VisitedOrderList[idx1-1]
			}
			newDistance := distCalc.GetDistance(lastNode, route.VisitedOrderList[idx2])
			if idx2 < routeLen-1 {
				newDistance += distCalc.GetDistance(route.VisitedOrderList[idx1], route.VisitedOrderList[idx2+1])
			}

			if newDistance < oldDistance-1 {
				fmt.Print("new ", newDistance, " old ", oldDistance, " ")
				return idx1, idx2
			}
		}
	}

	return -1, -1
}
