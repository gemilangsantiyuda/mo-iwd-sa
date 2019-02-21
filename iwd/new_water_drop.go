package iwd

import (
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

func newWaterDrop(soilMap SoilMap, ratingMap rating.Map, orderList []*order.Order, kitchenList []*kitchen.Kitchen, tree *mtree.Tree, distCalc distanceCalculator, config *config.Config) *WaterDrop {
	ksqMap := NewKitchenServedQtyMap(kitchenList)
	orderTree := tree.GetCopy()
	iwdParam := config.IwdParameter
	waterDrop := &WaterDrop{
		Velocity:            iwdParam.InitIWDVel,
		Soil:                iwdParam.InitIWDSoil,
		SoilMap:             soilMap,
		RatingMap:           ratingMap,
		KitchenServedQtyMap: ksqMap,
		OrderList:           orderList,
		KitchenList:         kitchenList,
		Tree:                orderTree,
		DistCalc:            distCalc,
		Config:              config,
	}
	return waterDrop
}
