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

// GetID return kitchen's id
func (kitchen *Kitchen) GetID() string {
	return kitchen.ID
}

// IsKitchen return true
func (kitchen *Kitchen) IsKitchen() bool {
	return true
}

// IsOrder return false
func (kitchen *Kitchen) IsOrder() bool {
	return false
}
