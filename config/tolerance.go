package config

// Tolerance for the score comparison
type Tolerance struct {
	RiderCost         int     `json:"rider_cost"`
	KitchenOptimality int     `json:"kitchen_optimality"`
	UserSatisfaction  float64 `json:"user_satisfaction"`
}
