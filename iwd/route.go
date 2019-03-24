package iwd

import (
	"log"
	"math"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
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

	for idx := range route.DistanceList {
		if idx > 0 && route.DistanceList[idx] < route.DistanceList[idx-1] {
			log.Fatal("Kecils split")
		}
	}
	// copy and cutting
	cutLen := len(route.VisitedOrderList) - orderIdx
	newVisOrderList := make([]*order.Order, cutLen)
	copy(newVisOrderList, route.VisitedOrderList[orderIdx:])
	route.VisitedOrderList = route.VisitedOrderList[:orderIdx]

	newDistanceList := make([]float64, cutLen)
	copy(newDistanceList, route.DistanceList[orderIdx:])
	route.DistanceList = route.DistanceList[:orderIdx]

	newRatingList := make([]float64, cutLen)
	copy(newRatingList, route.RatingList[orderIdx:])
	route.RatingList = route.RatingList[:orderIdx]

	newServedQtyList := make([]int, cutLen)
	copy(newServedQtyList, route.ServedQtyList[orderIdx:])
	route.ServedQtyList = route.ServedQtyList[:orderIdx]

	// because the lists are the sum of item 0 until idx .. F[idx] = sum(G[0]...G[idx]) then we have to remove the value of F[orderIdx-1] from the new list to remove the value from the other cut of the list

	removedDistance := 0.
	removedRating := 0.
	removedQty := 0
	if orderIdx > 0 {
		removedDistance = route.DistanceList[orderIdx-1]
		removedRating = route.RatingList[orderIdx-1]
		removedQty = route.ServedQtyList[orderIdx-1]
	}

	for idx := range newVisOrderList {
		newDistanceList[idx] -= removedDistance
		newRatingList[idx] -= removedRating
		newServedQtyList[idx] -= removedQty
	}

	// Adjust the route capacityLeft, and the kitchen served qty
	removedQty = newServedQtyList[len(newServedQtyList)-1]
	kitchen := route.ServingKitchen
	ksqMap.AddQty(kitchen, -removedQty)

	for idx := range newDistanceList {
		if idx > 0 && newDistanceList[idx] < newDistanceList[idx-1] {
			log.Fatal("Kecils split")
		}
	}
	return newVisOrderList, newDistanceList, newServedQtyList, newRatingList
}

// AddPath join a path into route
func (route *Route) AddPath(ksqMap KitchenServedQtyMap, ratingMap rating.Map, orderList []*order.Order, distanceList, ratingList []float64, servedQtyList []int, newDistance float64) {
	// update kitchen served qty
	ksqMap.AddQty(route.ServingKitchen, servedQtyList[len(servedQtyList)-1])

	oldLen := len(route.VisitedOrderList)

	// gabung
	route.VisitedOrderList = append(route.VisitedOrderList, orderList...)
	route.DistanceList = append(route.DistanceList, distanceList...)
	route.RatingList = append(route.RatingList, ratingList...)
	route.ServedQtyList = append(route.ServedQtyList, servedQtyList...)

	newLen := len(route.VisitedOrderList)

	// new value is added by  the old route last index value which is the sum from 0 -> last idx

	// adjust distance list

	oldDist := route.DistanceList[oldLen]
	for idx := oldLen; idx < newLen; idx++ {
		route.DistanceList[idx] -= oldDist
		route.DistanceList[idx] += newDistance
		if oldLen > 0 {
			route.DistanceList[idx] += route.DistanceList[oldLen-1]
		}
	}

	// adjust servedqty
	if oldLen > 0 {
		for idx := oldLen; idx < newLen; idx++ {
			route.ServedQtyList[idx] += route.ServedQtyList[oldLen-1]
		}
	}

	// adjust rating
	for idx := oldLen; idx < newLen; idx++ {
		order := route.VisitedOrderList[idx]
		rate := ratingMap.GetOrderToKitchenRating(order, route.ServingKitchen)
		route.RatingList[idx] = rate
		if idx > 0 {
			route.RatingList[idx] += route.RatingList[idx-1]
		}
	}

}

func (route *Route) reverseOrders(startIdx, endIdx int, distCalc distanceCalculator, ksqMap KitchenServedQtyMap, ratingMap rating.Map, config *config.Config) {

	if startIdx == endIdx {
		return
	}

	// reverse the visit list
	for idx := 0; startIdx+idx <= (startIdx+endIdx)/2; idx++ {
		temp := route.VisitedOrderList[startIdx+idx]
		route.VisitedOrderList[startIdx+idx] = route.VisitedOrderList[endIdx-idx]
		route.VisitedOrderList[endIdx-idx] = temp
	}

	servingKitchen := route.ServingKitchen
	for idx := startIdx; idx < len(route.VisitedOrderList); idx++ {
		order := route.VisitedOrderList[idx]
		dist := 0.
		qty := order.Quantity
		rate := ratingMap.GetOrderToKitchenRating(order, servingKitchen)
		if idx > 0 {
			lastOrder := route.VisitedOrderList[idx-1]
			dist = distCalc.GetDistance(lastOrder, order)

			dist += route.DistanceList[idx-1]
			qty += route.ServedQtyList[idx-1]
			rate += route.RatingList[idx-1]
		} else {
			dist = distCalc.GetDistance(servingKitchen, order)
		}
		route.DistanceList[idx] = dist
		route.ServedQtyList[idx] = qty
		route.RatingList[idx] = rate
	}

}
