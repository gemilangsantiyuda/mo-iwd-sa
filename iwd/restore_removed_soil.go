package iwd

func (wd *WaterDrop) restoreRemovedSoil() {
	var currentNode, nextNode node
	soilMap := wd.SoilMap
	for rIdx := range wd.RouteList {
		route := wd.RouteList[rIdx]
		currentNode = route.ServingKitchen
		for idx := range route.VisitedOrderList {
			nextNode = route.VisitedOrderList[idx]
			soil := soilMap.GetSoil(currentNode, nextNode)
			soil += wd.Config.IwdParameter.InitSoil
			soilMap.UpdateSoil(currentNode, nextNode, soil)
			currentNode = nextNode
		}
	}
}
