package main

import (
	"fmt"
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
	start := time.Now()
	config := config.ReadConfig()
	orderList := order.GetOrderList(config)
	kitchenList := kitchen.GetKitchenList(config)
	ratingMap := rating.GetRatingMap(config)
	distCalc := &distance.HaversineDistance{}
	tree := mtree.NewTree(config.MinTreeEntry, config.MaxTreeEntry, distCalc)

	currentSeed := int64(1551319497440281112)
	// currentSeed := time.Now().UTC().UnixNano()
	fmt.Println(currentSeed)
	rand.Seed(currentSeed)

	for idx := range orderList {
		order := orderList[idx]
		tree.Insert(order)
	}

	bestWDs := iwd.Solve(orderList, kitchenList, ratingMap, tree, distCalc, config)
	if bestWDs == nil {
		return
	}
	bestWD := bestWDs[0]
	fmt.Println(bestWD.Soil)
	fmt.Println("---------------")
	for idx := range bestWD.RouteList {
		route := bestWD.RouteList[idx]
		fmt.Println("Route", idx+1)
		fmt.Println("	Serving Kitchen : ", route.ServingKitchen.ID)
		fmt.Println("	Order visited :")
		fmt.Print("	-->")
		for odx := range route.VisitedOrderList {
			order := route.VisitedOrderList[odx]
			fmt.Print(order.ID, ",")
		}
		fmt.Println()
		fmt.Println("	Distance Traveled: ", route.GetDistanceTraveled())
		fmt.Println("	Served Qty: ", route.GetServedQty())
	}
	fmt.Println("-----------------")

	fmt.Println("Kitchen and Served Qty")
	ksqMap := bestWD.KitchenServedQtyMap
	for idx := range kitchenList {
		kitchen := kitchenList[idx]
		fmt.Println(kitchen.ID, " cap ", kitchen.Capacity, " served qty ", ksqMap.GetServedQty(kitchen))
	}
	fmt.Println("relapsed Time :", time.Since(start))

}
