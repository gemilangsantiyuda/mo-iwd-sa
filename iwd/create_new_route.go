package iwd

import (
	"github.com/vroup/mo-iwd-sa/kitchen"
)

func (wd *WaterDrop) createNewRoute() *Route {
	servingKitchen := wd.chooseNewKitchen()
	if servingKitchen == nil {
		return nil
	}

	capacityLeft := wd.Config.MaxDriverCapacity
	ksqMap := wd.KitchenServedQtyMap
	kitchenCapLeft := servingKitchen.Capacity.Maximum - ksqMap.GetServedQty(servingKitchen)
	if capacityLeft > kitchenCapLeft {
		capacityLeft = kitchenCapLeft
	}
	route := &Route{
		ServingKitchen: servingKitchen,
		CapacityLeft:   capacityLeft,
	}
	return route
}

func (wd *WaterDrop) chooseNewKitchen() *kitchen.Kitchen {
	kitchenList := wd.getPossibleKitchenList()
	if len(kitchenList) == 0 {
		return nil
	}
	kitchenPreferenceList := wd.getKitchenPreferenceList(kitchenList)
	kitchenIdx := chooseIdxByRouletteWheel(kitchenPreferenceList)
	if kitchenIdx == -1 {
		return nil
	}
	chosenKitchen := kitchenList[kitchenIdx]
	return chosenKitchen
}
