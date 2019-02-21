package iwd

import (
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

type neighbour interface {
	GetOrder() *order.Order
	GetDistance() float64
}

// WaterDrop struct for the IWD
type WaterDrop struct {
	RouteList           []*Route
	Score               *Score
	WeightedScore       float64
	Velocity            float64
	Soil                float64
	SoilMap             SoilMap
	RatingMap           rating.Map
	KitchenServedQtyMap KitchenServedQtyMap
	OrderList           []*order.Order
	KitchenList         []*kitchen.Kitchen
	Tree                *mtree.Tree
	Config              *config.Config
	DistCalc            distanceCalculator
}

// calcScore after solving an iteration wd get the score of the routelist made
func (wd *WaterDrop) calcScore() {
	totalRiderCost := 0
	kitchenOptimality := 0
	totalUserRating := 0.
	ksqMap := wd.KitchenServedQtyMap
	for idx := range wd.RouteList {
		route := wd.RouteList[idx]
		totalRiderCost += route.RiderCost
		totalUserRating += route.GetTotalRating()
	}

	totalUserRating = float64(len(wd.OrderList)) / totalUserRating

	for idx := range wd.KitchenList {
		kitchen := wd.KitchenList[idx]
		servedQty := ksqMap.GetServedQty(kitchen)
		if servedQty > 0 {
			optimalityDist := kitchen.Capacity.Optimum - servedQty
			if optimalityDist < 0 {
				optimalityDist *= -1
			}
			kitchenOptimality += optimalityDist
		}
	}
	score := &Score{
		RiderCost:         totalRiderCost,
		KitchenOptimality: kitchenOptimality,
		UserSatisfaction:  totalUserRating,
	}
	wd.Score = score
}

func (wd *WaterDrop) getCurrentRoute() *Route {
	// check if wd has no route yet then make new route
	if len(wd.RouteList) == 0 {
		newRoute := wd.createNewRoute()
		wd.RouteList = append(wd.RouteList, newRoute)
		return newRoute
	}
	currentRoute := wd.RouteList[len(wd.RouteList)-1]
	return currentRoute
}

func (wd *WaterDrop) getCurrentNode(currentRoute *Route) node {
	// check if no node visited yet then return the serving kitchen
	if len(currentRoute.VisitedOrderList) == 0 {
		return currentRoute.ServingKitchen
	}
	currentNode := currentRoute.VisitedOrderList[len(currentRoute.VisitedOrderList)-1]
	return currentNode
}
