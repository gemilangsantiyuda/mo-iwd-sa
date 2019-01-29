package iwd

// check after traversing wether this waterdrop generate a valid routelist or not
func (wd *WaterDrop) hasValidRouteList() bool {
	tree := wd.Tree
	// // there is still order unvisited
	if tree.ObjectCount > 0 {
		// fmt.Println("FALSE")
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

	// check each route's distance < maxdriverdistance and cap left of each route
	// and check all order has been really visited
	orderCount := len(wd.OrderList)
	totalOrderQty := 0

	for idx := range wd.OrderList {
		totalOrderQty += wd.OrderList[idx].Quantity
	}

	for idx := range wd.RouteList {
		route := wd.RouteList[idx]
		if route.GetDistanceTraveled() > wd.Config.MaxDriverDistance {
			return false
		}
		if route.GetCapacityLeft(ksqMap, wd.Config) < 0 {
			return false
		}
		orderCount -= len(route.VisitedOrderList)
		for odx := range route.VisitedOrderList {
			totalOrderQty -= route.VisitedOrderList[odx].Quantity
		}
	}

	// fmt.Println(orderCount, " ", totalOrderQty)

	if (orderCount != 0) || (totalOrderQty != 0) {
		return false
	}

	return true
}
