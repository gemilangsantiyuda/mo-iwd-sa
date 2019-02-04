package iwd

func globalUpdate(soilMap SoilMap, wd *WaterDrop) {
	iwdParam := wd.Config.IwdParameter
	orderCount := float64(len(wd.OrderList))
	var currentNode, nextNode node
	for rIdx := range wd.RouteList {
		route := wd.RouteList[rIdx]
		currentNode = route.ServingKitchen
		for idx := range route.VisitedOrderList {
			nextNode = route.VisitedOrderList[idx]
			soil := soilMap.GetSoil(currentNode, nextNode)
			wdSoil := wd.Soil
			newSoil := (1+iwdParam.P)*soil - iwdParam.P*wdSoil/(orderCount*(orderCount-1))
			soilMap.UpdateSoil(currentNode, nextNode, newSoil)
			currentNode = nextNode
		}
	}
}
