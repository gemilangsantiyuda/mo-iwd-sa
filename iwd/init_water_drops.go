package iwd

import (
	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

func initWaterDrops(soilMap SoilMap, ratingMap rating.Map, orderList []*order.Order, kitchenList []*kitchen.Kitchen, tree *mtree.Tree, config *config.Config) []*WaterDrop {
	var waterDropList []*WaterDrop
	for idx := 0; idx < config.IwdParameter.PopulationSize; idx++ {
		wd := newWaterDrop(soilMap, ratingMap, orderList, kitchenList, tree, config)
		waterDropList = append(waterDropList, wd)
	}
	return waterDropList
}
