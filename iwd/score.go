package iwd

// Score holds the 3 values of the objective functions
type Score struct {
	RiderCost         int
	KitchenOptimality int
	UserSatisfaction  float64
}
