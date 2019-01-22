package order

import (
	"github.com/vroup/mo-iwd-sa/coordinate"
)

// Order to holds order information
type Order struct {
	ID         string
	UserID     string
	Quantity   int
	Coordinate *coordinate.Coordinate
}

// GetCoordinate return this order's coord
func (order *Order) GetCoordinate() *coordinate.Coordinate {
	return order.Coordinate
}

// GetUserID return user id of this order
func (order *Order) GetUserID() string {
	return order.UserID
}

// GetID return order's id
func (order *Order) GetID() string {
	return order.ID
}

// IsKitchen false
func (order *Order) IsKitchen() bool {
	return false
}

// IsOrder true
func (order *Order) IsOrder() bool {
	return true
}
