package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/config"
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
	CapacityLeft     int
	TotalRating      float64
}

// CalcRiderCost recalculate rider cost based
func (route *Route) CalcRiderCost(config *config.Config) {
	duration := int(math.Ceil(route.DistanceTraveled / config.DriverSpeed))
	route.RiderCost = duration * config.DriverRate
}
