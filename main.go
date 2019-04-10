package main

import (
	"math/rand"
	"time"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/distance"
	"github.com/vroup/mo-iwd-sa/iwd"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

func main() {

	currentSeed := time.Now().UTC().UnixNano()
	// fmt.Println(currentSeed)
	rand.Seed(currentSeed)
	startTime := time.Now()
	config := config.ReadConfig()
	config.MinTreeEntry = config.MaxTreeEntry / 2
	config.ReadMaxValue()
	orderList := order.GetOrderList(config)
	kitchenList := kitchen.GetKitchenList(config)
	ratingMap := rating.GetRatingMap(config)
	distCalc := &distance.HaversineDistance{}
	tree := mtree.NewTree(config.MinTreeEntry, config.MaxTreeEntry, distCalc)
	referencePoint := iwd.ReadRF(config)
	config.BestArchiveSize = len(referencePoint)
	// currentSeed := int64(1551319497440281112)

	for idx := range orderList {
		order := orderList[idx]
		tree.Insert(order)
	}
	iwd.Solve(0, orderList, kitchenList, ratingMap, tree, distCalc, referencePoint, startTime, config, nil)

}
