package iwd

import (
	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/mtree"
	"github.com/vroup/mo-iwd-sa/order"
)

type neighbour interface {
	GetOrder() *order.Order
	GetDistance() float64
}

// WaterDrop struct for the IWD
type WaterDrop struct {
	RouteList           []*Route
	Score               *Score
	WeightedScore       float64
	Velocity            float64
	Soil                float64
	SoilMap             SoilMap
	KitchenServedQtyMap KitchenServedQtyMap
	OrderList           []*order.Order
	KitchenList         []*kitchen.Kitchen
	Tree                *mtree.Tree
}
