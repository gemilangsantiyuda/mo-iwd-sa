package iwd

import (
	"math/rand"

	"github.com/vroup/mo-iwd-sa/order"
)

// // first copy the
func (wd *WaterDrop) getMutation(maxSwap int) *WaterDrop {
	newWD := &WaterDrop{
		RouteList:           make([]*Route, 0),
		Velocity:            wd.Velocity,
		Soil:                wd.Soil,
		SoilMap:             wd.SoilMap,
		RatingMap:           wd.RatingMap,
		KitchenServedQtyMap: wd.KitchenServedQtyMap.GetCopy(),
		OrderList:           wd.OrderList,
		KitchenList:         wd.KitchenList,
		Config:              wd.Config,
		Tree:                wd.Tree,
	}

	// Copying route
	for rIdx := range wd.RouteList {
		route := wd.RouteList[rIdx]
		newRoute := &Route{
			ServingKitchen:   route.ServingKitchen,
			VisitedOrderList: make([]*order.Order, len(route.VisitedOrderList)),
			DistanceList:     make([]float64, len(route.DistanceList)),
			RatingList:       make([]float64, len(route.RatingList)),
			ServedQtyList:    make([]int, len(route.ServedQtyList)),
			RiderCost:        route.RiderCost,
		}
		copy(newRoute.VisitedOrderList, route.VisitedOrderList)
		copy(newRoute.DistanceList, route.DistanceList)
		copy(newRoute.RatingList, route.RatingList)
		copy(newRoute.ServedQtyList, route.ServedQtyList)
		newWD.RouteList = append(newWD.RouteList, newRoute)
	}

	for swap := 0; swap < maxSwap; swap++ {
		newWD.randomSwap()
	}
	return newWD
}

// randomswap in 2 different route, if wd has only 1 route then return
func (wd *WaterDrop) randomSwap() {
	if len(wd.RouteList) < 2 {
		return
	}
	routeIdx1 := rand.Intn(len(wd.RouteList))
	routeIdx2 := rand.Intn(len(wd.RouteList))
	for routeIdx1 == routeIdx2 {
		routeIdx2 = rand.Intn(len(wd.RouteList))
	}

	route1 := wd.RouteList[routeIdx1]
	route2 := wd.RouteList[routeIdx2]

	orderIdx1 := rand.Intn(len(route1.VisitedOrderList))
	orderIdx2 := rand.Intn(len(route2.VisitedOrderList))

	wd.swapRoutePath(route1, route2, orderIdx1, orderIdx2)
}

func (wd *WaterDrop) swapRoutePath(route1, route2 *Route, orderIdx1, orderIdx2 int) {
	// pisah masing-masing route pada index yang ditentukan
	ksqMap := wd.KitchenServedQtyMap
	orderList1, distanceList1, servedQtyList1, ratingList1 := route1.Split(ksqMap, orderIdx1)
	orderList2, distanceList2, servedQtyList2, ratingList2 := route2.Split(ksqMap, orderIdx2)

	distCalc := wd.Tree.DistCalc
	rtMap := wd.RatingMap

	// update list yang terpisah dengan nilai baru masing2 list yang dihasilkan jika mereka digabungkan ke rute lain
	var currentNode1, currentNode2 node
	kitchen1 := route1.ServingKitchen
	kitchen2 := route2.ServingKitchen
	if len(route1.VisitedOrderList) == 0 {
		currentNode1 = kitchen1
	} else {
		currentNode1 = route1.VisitedOrderList[len(route1.VisitedOrderList)-1]
	}
	if len(route2.VisitedOrderList) == 0 {
		currentNode2 = kitchen2
	} else {
		currentNode2 = route2.VisitedOrderList[len(route2.VisitedOrderList)-1]
	}

	nextNode1 := orderList2[0]
	nextNode2 := orderList1[0]

	newDistance1 := distCalc.GetDistance(currentNode1, nextNode1)
	newRating1 := rtMap.GetOrderToKitchenRating(nextNode1, kitchen1)
	newServedQty1 := nextNode1.Quantity

	newDistance2 := distCalc.GetDistance(currentNode2, nextNode2)
	newRating2 := rtMap.GetOrderToKitchenRating(nextNode2, kitchen2)
	newServedQty2 := nextNode2.Quantity

	// gabungkan
	route1.AddPath(ksqMap, orderList2, distanceList2, ratingList2, servedQtyList2, newDistance1, newRating1, newServedQty1)
	route2.AddPath(ksqMap, orderList1, distanceList1, ratingList1, servedQtyList1, newDistance2, newRating2, newServedQty2)

	// recalculate rider cost
	route1.CalcRiderCost(wd.Config)
	route2.CalcRiderCost(wd.Config)
}
