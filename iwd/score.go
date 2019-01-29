package iwd

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
