package kitchen

import "github.com/vroup/mo-iwd-sa/coordinate"

// Kitchen holds kitchen info
type Kitchen struct {
	ID         string
	Coordinate *coordinate.Coordinate
	Capacity   *Capacity
}

// Capacity holds kitchen caps info
type Capacity struct {
	Minimum, Optimum, Maximum int
}

// GetCoordinate return this coord
func (kitchen *Kitchen) GetCoordinate() *coordinate.Coordinate {
	return kitchen.Coordinate
}

// GetKitchenID return kitchen's id
func (kitchen *Kitchen) GetKitchenID() string {
	return kitchen.ID
}
