package iwd

import "github.com/vroup/mo-iwd-sa/mtree"

type distanceCalculator interface {
	GetDistance(mtree.Object, mtree.Object) float64
}
