package rating

// GetOrderToKitchenRating return rating of order (based on its userID) toward a kitchen, combining their ID with a '+' as the key to the map
func (m Map) GetOrderToKitchenRating(od order, kc kitchen) float64 {
	key := od.GetUserID() + "+" + kc.GetKitchenID()
	return m[key]
}
