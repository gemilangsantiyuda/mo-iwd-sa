package iwd

func (wd *WaterDrop) updateKitchenPreference() {
	totalServedQty := 0
	ksqMap := wd.KitchenServedQtyMap
	for idx := range wd.KitchenList {
		kitchen := wd.KitchenList[idx]
		servedQty := ksqMap.GetServedQty(kitchen)
		totalServedQty += servedQty
	}

	for idx := range wd.KitchenList {
		kitchen := wd.KitchenList[idx]
		servedQty := ksqMap.GetServedQty(kitchen)
		preference := float64(servedQty) / float64(totalServedQty) / 10000.
		kitchen.Preference += preference
	}
}
