package iwd

import "github.com/vroup/mo-iwd-sa/config"

// Score holds the 3 values of the objective functions
type Score struct {
	RiderCost         int
	KitchenOptimality int
	UserSatisfaction  float64
}

// GetWeightedScore for the weighted sum method
func (score *Score) GetWeightedScore(weight *config.Weight, KitchenOptimalityMax int) float64 {
	riderCost := float64(score.RiderCost) / 75000.
	kitchenOpt := float64(score.KitchenOptimality) / float64(KitchenOptimalityMax)
	userSat := score.UserSatisfaction / 5.
	weightedSum := riderCost*weight.RiderCost + kitchenOpt*weight.KitchenOptimality + userSat*weight.UserSatisfaction
	return weightedSum
}
