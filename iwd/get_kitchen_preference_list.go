package iwd

import (
	"github.com/vroup/mo-iwd-sa/kitchen"
)

func (wd *WaterDrop) getKitchenPreferenceList(kitchenList []*kitchen.Kitchen) []float64 {
	var preferenceList []float64
	for idx := range kitchenList {
		kitchen := kitchenList[idx]
		preference := wd.getKitchenPreference(kitchen)
		preferenceList = append(preferenceList, preference)
	}
	return preferenceList
}

func (wd *WaterDrop) getKitchenPreference(kitchen *kitchen.Kitchen) float64 {
	servedQty := wd.KitchenServedQtyMap.GetServedQty(kitchen)
	kitchenCap := kitchen.Capacity
	preference := 0.
	if servedQty <= kitchenCap.Minimum {
		preference = (float64(servedQty) / (float64(2*kitchenCap.Minimum) + 0.001)) + 0.5
	} else if servedQty < kitchenCap.Optimum {
		preference = 1.
	} else {
		preference = (float64(kitchenCap.Maximum-servedQty) / float64(kitchenCap.Maximum-kitchenCap.Optimum)) * (float64(kitchenCap.Maximum-servedQty) / float64(kitchenCap.Maximum-kitchenCap.Optimum))
	}
	return preference
}
