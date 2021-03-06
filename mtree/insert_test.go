package mtree_test

import (
	"fmt"
	"math"

	"github.com/mtree"
)

// your object must have unique id each, and must have a GetID() method
type object struct {
	id   string
	x, y float64
}

// GetID method to return the object id
func (obj *object) GetID() string {
	return obj.id
}

// define a distance calculator for the tree
type distanceCalculator struct {
}

func (distCalc *distanceCalculator) GetDistance(origin, destination mtree.Object) float64 {
	objOrigin := origin.(*object)
	objDest := destination.(*object)
	dist := math.Sqrt((objOrigin.x-objDest.x)*(objOrigin.x-objDest.x) + (objOrigin.y-objDest.y)*(objOrigin.y-objDest.y))
	return dist
}

var objectList = []*object{
	&object{"1", 0, 0},
	&object{"2", 2, 5},
	&object{"3", 5, 3},
	&object{"4", 0, 9},
	&object{"5", 2, 0},
	&object{"6", 5, 0},
}

var distCalc = &distanceCalculator{}

func main() {

	minEntry := 2
	maxEntry := 4
	tree := mtree.NewTree(minEntry, maxEntry, distCalc)
	for _, obj := range objectList {
		tree.Insert(obj)
	}

	queryObject := &object{
		id: "7",
		x:  3,
		y:  3,
	}
	nearestNeighbourList := tree.KnnSearch(queryObject, 3)
	for idx, neighbour := range nearestNeighbourList {
		fmt.Printf("%d-th nearest neighbour = %+v , with distance = %f\n", idx, neighbour.Object, neighbour.Distance)
	}

	tree.Remove(objectList[2])
	fmt.Println("After removing ", objectList[2])

	nearestNeighbourList = tree.KnnSearch(queryObject, 3)
	for idx, neighbour := range nearestNeighbourList {
		fmt.Printf("%d-th nearest neighbour = %+v , with distance = %f\n", idx, neighbour.Object, neighbour.Distance)
	}

}
