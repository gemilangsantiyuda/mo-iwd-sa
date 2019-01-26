package iwd

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
