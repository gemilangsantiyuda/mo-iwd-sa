package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/kitchen"
	"github.com/vroup/mo-iwd-sa/order"
)

// Route to hold info about a route
type Route struct {
	ServingKitchen   *kitchen.Kitchen
	VisitedOrderList []*order.Order
	DistanceTraveled float64
	RiderCost        int
	ServedQty        int
}

// CalcRiderCost recalculate rider cost based
func (route *Route) CalcRiderCost() {
	duration := int(math.Ceil(route.DistanceTraveled / 30000.))
	route.RiderCost = duration * 25000
}
