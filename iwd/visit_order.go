package iwd

import "github.com/vroup/mo-iwd-sa/order"

func (wd *WaterDrop) visitOrder(currentRoute *Route, currentNode node, nextOrder *order.Order, distance, moDistance float64) {
	tree := wd.Tree
	rtMap := wd.RatingMap
	ksqMap := wd.KitchenServedQtyMap
	servingKitchen := currentRoute.ServingKitchen

	currentRoute.VisitedOrderList = append(currentRoute.VisitedOrderList, nextOrder)
	tree.Remove(tree.Root, nextOrder, nextOrder.ID)
	currentRoute.CapacityLeft -= nextOrder.Quantity
	currentRoute.DistanceTraveled += distance
	currentRoute.TotalRating += rtMap.GetOrderToKitchenRating(nextOrder, servingKitchen)
	currentRoute.ServedQty += nextOrder.Quantity
	ksqMap.AddQty(servingKitchen, nextOrder.Quantity)
	wd.updateDynamicParameter(currentNode, nextOrder, moDistance)
}
