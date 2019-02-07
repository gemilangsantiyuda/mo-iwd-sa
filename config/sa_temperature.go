package config

// SaParameter temperature for each objective function in SA and rate
type SaParameter struct {
	Temp        float64 `json:"temp"`
	CoolingRate float64 `json:"cooling_rate"`
}

// Temp temperature for each objective
// type Temp struct {
// 	RiderCost         float64 `json:"rider_cost"`
// 	KitchenOptimality float64 `json:"kitchen_optimality"`
// 	UserSatisfaction  float64 `json:"user_satisfaction"`
// }
