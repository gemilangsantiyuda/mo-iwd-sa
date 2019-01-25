package iwd

import (
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/order"
)

// SoilMap hold the soil of edge
type SoilMap map[string]float64

// GetCopy return a copy of the map
func (sm SoilMap) GetCopy() SoilMap {
	newSM := make(SoilMap)
	for key, val := range sm {
		sm[key] = val
	}
	return newSM
}

// GetSoil return soil value of the edge
func (sm SoilMap) GetSoil(origin, destination node) float64 {
	key := origin.GetID() + "+" + destination.GetID()
	if origin.IsKitchen() {
		key = "K" + key
	}
	return sm[key]
}

// UpdateSoil renew soil value
func (sm SoilMap) UpdateSoil(origin, destination node, newSoil float64) {
	key := origin.GetID() + "+" + destination.GetID()
	if origin.IsKitchen() {
		key = "K" + key
	}
	sm[key] = newSoil
}

// NewSoilMap initialize soil map
func NewSoilMap(kitchenList []*kitchen.Kitchen, orderList []*order.Order, conf *config.Config) SoilMap {
	sm := make(SoilMap)
	iwdParam := conf.IwdParameter

	// kitchen -> orders
	for kitchenIdx := range kitchenList {
		kitchenID := "K" + kitchenList[kitchenIdx].GetID()
		for orderIdx := range orderList {
			orderID := orderList[orderIdx].GetID()
			key := kitchenID + "+" + orderID
			sm[key] = iwdParam.InitSoil
		}
	}

	// orders -> orders
	for orderIdx1 := range orderList {
		for orderIdx2 := range orderList {
			if orderIdx1 != orderIdx2 {
				orderID1 := orderList[orderIdx1].GetID()
				orderID2 := orderList[orderIdx2].GetID()
				key := orderID1 + "+" + orderID2
				sm[key] = iwdParam.InitSoil
			}
		}
	}
	return sm
}
