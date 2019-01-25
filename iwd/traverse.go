package iwd

func (wd *WaterDrop) traverse() {
	tree := wd.Tree
	rtMap := wd.RatingMap
	ksqMap := wd.KitchenServedQtyMap
	for tree.ObjectCount > 0 {
		route := wd.createNewRoute()
		// no route can be made anymore because no kitchen can be chosen
		if route == nil {
			break
		}
		servingKitchen := route.ServingKitchen
		currentNode := servingKitchen
		nextOrder, distance, moDistance := wd.getNextOrder(route, currentNode)
		for nextOrder != nil {
			route.VisitedOrderList = append(route.VisitedOrderList, nextOrder)
			tree.Remove(tree.Root, nextOrder, nextOrder.ID)
			route.CapacityLeft -= nextOrder.Quantity
			route.DistanceTraveled += distance
			route.TotalRating += rtMap.GetOrderToKitchenRating(nextOrder, servingKitchen)
			ksqMap.AddQty(servingKitchen, nextOrder.Quantity)
			wd.updateDynamicParameter(currentNode, nextOrder, moDistance)

			nextOrder, distance, moDistance = wd.getNextOrder(route, currentNode)
		}
		wd.RouteList = append(wd.RouteList, route)
	}

	if wd.hasValidRouteList() {
		wd.calcScore()
	}
}

// check after traversing wether this waterdrop generate a valid routelist or not
func (wd *WaterDrop) hasValidRouteList() bool {
	tree := wd.Tree
	// there is still order unvisited
	if tree.ObjectCount > 0 {
		return false
	}
	// a kitchen serving but served less then its minimum capacity
	kitchenList := wd.KitchenList
	ksqMap := wd.KitchenServedQtyMap
	for idx := range kitchenList {
		kitchen := kitchenList[idx]
		servedQty := ksqMap.GetServedQty(kitchen)
		if servedQty > 0 {
			if servedQty < kitchen.Capacity.Minimum {
				return false
			}
		}
	}
	return true
}
