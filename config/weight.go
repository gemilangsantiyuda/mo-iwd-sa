package config

// Weight for the weighted sum objective function
type Weight struct {
	RiderCost         float64 `json:"rider_cost"`
	KitchenOptimality float64 `json:"kitchen_optimality"`
	UserSatisfaction  float64 `json:"user_satisfaction"`
}
