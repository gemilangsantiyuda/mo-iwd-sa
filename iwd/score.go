package iwd

import (
	"math"

	"github.com/vroup/mo-iwd-sa/config"
)

// Score holds the 3 values of the objective functions
type Score struct {
	RiderCost         int
	KitchenOptimality int
	UserSatisfaction  float64
}

// IsWorseThan if score is less than newscore lexicographically with tolerance
func (score *Score) IsWorseThan(newScore *Score) bool {
	if newScore.RiderCost < score.RiderCost {
		return true
	}
	if float64(newScore.RiderCost) == float64(score.RiderCost) {
		if float64(newScore.RiderCost) == float64(score.RiderCost) {
			if newScore.KitchenOptimality < score.KitchenOptimality {
				return true
			}
		}
		if newScore.KitchenOptimality < int(0.8*float64(score.KitchenOptimality)) {
			return true
		}
		if float64(newScore.KitchenOptimality) < 1.10*float64(score.KitchenOptimality) {
			return newScore.UserSatisfaction < score.UserSatisfaction
		}
	}
	return false
}

// IsDominate return wether the score dominate the newscore
func (score *Score) IsDominate(newScore *Score, tolerance config.Tolerance) bool {
	return (score.RiderCost <= tolerance.RiderCost+newScore.RiderCost) && (score.KitchenOptimality <= tolerance.KitchenOptimality+newScore.KitchenOptimality) && (score.UserSatisfaction <= tolerance.UserSatisfaction+newScore.UserSatisfaction)
}

// GetDifference return the difference between the 2 scores
func (score *Score) GetDifference(score2 *Score) float64 {
	return math.Abs(float64(score.RiderCost-score2.RiderCost)) + math.Abs(float64(score.KitchenOptimality-score2.KitchenOptimality)) + math.Abs(score.UserSatisfaction-score2.UserSatisfaction)
}
