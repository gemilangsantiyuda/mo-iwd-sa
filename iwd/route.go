package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/order"
)

// Route to hold info about a route
type Route struct {
	ServingKitchen   *kitchen.Kitchen
	VisitedOrderList []*order.Order
	DistanceList     []float64
	RatingList       []float64
	RiderCost        int
	ServedQtyList    []int
}

// CalcRiderCost recalculate rider cost based
func (route *Route) CalcRiderCost(config *config.Config) {
	distanceTraveled := route.GetDistanceTraveled()
	duration := int(math.Ceil(distanceTraveled / config.DriverSpeed))
	route.RiderCost = duration * config.DriverRate
}

// GetCapacityLeft return the capacity left for this route based on the served quantity and the kitchen cap left
func (route *Route) GetCapacityLeft(ksqMap KitchenServedQtyMap, config *config.Config) int {
	servedQty := route.GetServedQty()
	capLeft := config.MaxDriverCapacity - servedQty
	kitchen := route.ServingKitchen
	kitchenCapLeft := kitchen.Capacity.Maximum - ksqMap.GetServedQty(kitchen)
	if capLeft > kitchenCapLeft {
		capLeft = kitchenCapLeft
	}
	return capLeft
}

// GetDistanceTraveled return the total traveled distance
func (route *Route) GetDistanceTraveled() float64 {
	if len(route.DistanceList) == 0 {
		return 0
	}
	return route.DistanceList[len(route.DistanceList)-1]
}

// GetTotalRating return the total rating
func (route *Route) GetTotalRating() float64 {
	if len(route.RatingList) == 0 {
		return 0
	}
	return route.RatingList[len(route.RatingList)-1]
}

// GetServedQty return the total served qty
func (route *Route) GetServedQty() int {
	if len(route.ServedQtyList) == 0 {
		return 0
	}
	return route.ServedQtyList[len(route.ServedQtyList)-1]
}

// Split split the route exactly at the orderIDx, some kind of cutting the path into two returning the information about that path. Essential to cut-paste path (e.g 2-opt)
func (route *Route) Split(ksqMap KitchenServedQtyMap, orderIdx int) ([]*order.Order, []float64, []int, []float64) {
	newVisOrderList := route.VisitedOrderList[orderIdx:]
	// copy and cutting
	route.VisitedOrderList = route.VisitedOrderList[:orderIdx]
	newDistanceList := route.DistanceList[orderIdx:]
	route.DistanceList = route.DistanceList[:orderIdx]
	newRatingList := route.RatingList[orderIdx:]
	route.RatingList = route.RatingList[:orderIdx]
	newServedQtyList := route.ServedQtyList[orderIdx:]
	route.ServedQtyList = route.ServedQtyList[:orderIdx]

	// because the lists are the sum of item 0 until idx .. F[idx] = sum(G[0]...G[idx]) then we have to remove the value of F[orderIdx] from the new list to remove the value from the other cut of the list
	removedDistance := newDistanceList[0]
	removedRating := newRatingList[0]
	removedQty := newServedQtyList[0]
	for idx := range newVisOrderList {
		newDistanceList[idx] -= removedDistance
		newRatingList[idx] -= removedRating
		newServedQtyList[idx] -= removedQty
	}

	// Adjust the route capacityLeft, and the kitchen served qty
	removedQty = newServedQtyList[len(newServedQtyList)-1]
	kitchen := route.ServingKitchen
	ksqMap.AddQty(kitchen, -removedQty)

	return newVisOrderList, newDistanceList, newServedQtyList, newRatingList
}

// AddPath join a path into route
func (route *Route) AddPath(ksqMap KitchenServedQtyMap, orderList []*order.Order, distanceList, ratingList []float64, servedQtyList []int, newDistance, newRating float64, newServedQty int) {
	// update kitchen served qty
	ksqMap.AddQty(route.ServingKitchen, newServedQty)

	// new value is added by  the old route last index value which is the sum from 0 -> last idx

	if len(route.VisitedOrderList) > 0 {
		newDistance += route.DistanceList[len(route.DistanceList)-1]
		newRating += route.RatingList[len(route.RatingList)-1]
		newServedQty += route.ServedQtyList[len(route.ServedQtyList)-1]
	}

	// update the list with their new added values
	for idx := range orderList {
		distanceList[idx] += newDistance
		ratingList[idx] += newRating
		servedQtyList[idx] += newServedQty
	}

	// gabung
	route.VisitedOrderList = append(route.VisitedOrderList, orderList...)
	route.DistanceList = append(route.DistanceList, distanceList...)
	route.RatingList = append(route.RatingList, ratingList...)
	route.ServedQtyList = append(route.ServedQtyList, servedQtyList...)

}
