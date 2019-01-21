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
