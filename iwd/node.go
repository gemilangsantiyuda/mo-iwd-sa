package iwd

import (
	"github.com/vroup/mo-iwd-sa/coordinate"
)

type node interface {
	GetID() string
	IsKitchen() bool
	IsOrder() bool
	GetCoordinate() *coordinate.Coordinate
}
