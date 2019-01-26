package main

import (
	"fmt"

	"github.com/vroup/mo-iwd-sa/config"
	"github.com/vroup/mo-iwd-sa/distance"
	"github.com/vroup/mo-iwd-sa/iwd"
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
	"github.com/vroup/mo-iwd-sa/rating"
)

func main() {
	config := config.ReadConfig()
	orderList := order.GetOrderList(config)
	kitchenList := kitchen.GetKitchenList(config)
	ratingMap := rating.GetRatingMap(config)
	distCalc := &distance.HaversineDistance{}
	tree := mtree.NewTree(config.MaxTreeEntry, nil, distCalc)

	for idx := range orderList {
		order := orderList[idx]
		id := order.ID
		tree.Insert(order, id)
	}

	bestWD := iwd.Solve(orderList, kitchenList, ratingMap, tree, config)
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
		fmt.Println("	Distance Traveled: ", route.DistanceTraveled)
		fmt.Println("	Served Qty: ", route.ServedQty)
	}
	fmt.Println("-----------------")

	fmt.Println("Kitchen and Served Qty")
	ksqMap := bestWD.KitchenServedQtyMap
	for idx := range kitchenList {
		kitchen := kitchenList[idx]
		fmt.Println(kitchen.ID, " cap ", kitchen.Capacity, " served qty ", ksqMap.GetServedQty(kitchen))
	}
}
