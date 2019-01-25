package iwd

import "github.com/vroup/mo-iwd-sa/config"

func calcMoDistance(distance, kitchenOpt, userRating float64, weight *config.Weight) float64 {
	return distance*weight.RiderCost + kitchenOpt*weight.KitchenOptimality + userRating*weight.UserSatisfaction
}
