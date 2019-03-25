package config

// MaxValue for normalizing score
type MaxValue struct {
	RiderCost         int     `json:"rider_cost"`
	KitchenOptimality int     `json:"kitchen_optimality"`
	UserSatisfaction  float64 `json:"user_satisfaction"`
}
