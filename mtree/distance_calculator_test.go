package mtree_test

import (
	"math"

	"github.com/mtree"
)

type distCalcTest struct {
}

func (dc *distCalcTest) GetDistance(origin mtree.Object, dest mtree.Object) float64 {
	rOrigin := origin.(*object)
	rDest := dest.(*object)
	dist := math.Sqrt((rOrigin.x-rDest.x)*(rOrigin.x-rDest.x) + (rOrigin.y-rDest.y)*(rOrigin.y-rDest.y))
	return dist
}
