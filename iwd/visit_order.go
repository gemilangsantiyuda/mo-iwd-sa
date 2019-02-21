package iwd

import "github.com/vroup/mo-iwd-sa/order"

func (wd *WaterDrop) visitOrder(currentRoute *Route, currentNode node, nextOrder *order.Order, distance, moDistance float64) {
	tree := wd.Tree
	rtMap := wd.RatingMap
	ksqMap := wd.KitchenServedQtyMap
	servingKitchen := currentRoute.ServingKitchen

	currentRoute.VisitedOrderList = append(currentRoute.VisitedOrderList, nextOrder)
	tree.Remove(nextOrder)
	distance += currentRoute.GetDistanceTraveled()
	currentRoute.DistanceList = append(currentRoute.DistanceList, distance)
	rating := rtMap.GetOrderToKitchenRating(nextOrder, servingKitchen)
	rating += currentRoute.GetTotalRating()
	currentRoute.RatingList = append(currentRoute.RatingList, rating)
	servedQty := nextOrder.Quantity + currentRoute.GetServedQty()
	currentRoute.ServedQtyList = append(currentRoute.ServedQtyList, servedQty)
	ksqMap.AddQty(servingKitchen, nextOrder.Quantity)
	wd.updateDynamicParameter(currentNode, nextOrder, moDistance)
}
