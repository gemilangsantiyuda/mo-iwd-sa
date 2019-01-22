package iwd

import "github.com/vroup/mo-iwd-sa/kitchen"

func (wd *WaterDrop) createNewRoute() *Route {
	servingKitchen := wd.chooseNewKitchen()
	route := &Route{
		ServingKitchen: servingKitchen,
	}
	return route
}

func (wd *WaterDrop) chooseNewKitchen() *kitchen.Kitchen {
	kitchenList := wd.KitchenList
	kitchenPreferenceList := wd.getKitchenPreferenceList()
	kitchenIdx := chooseIdxByRouletteWheel(kitchenPreferenceList)
	chosenKitchen := kitchenList[kitchenIdx]
	return chosenKitchen
}
