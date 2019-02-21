# go-mtree
M-Tree implementation for Golang

This implementation is based on a technical report on improvisation of M-Tree to support deletion called Symmetric M-Tree (Sexton and Swinbank, 2003).

The split mechanism used is the minimum spanning tree (MST) split from Slim Tree (Traina et al., 2002)

and finally the KNN-Search method used is best first search method which is firstly implemented for the R-Tree (Hjaltason and Samet, 1999)


Sexton, Alan P. and Swinbank, Richard (2003) <i> Symmetric M-Tree </i>, University of Birmingham. Available at : http://www.cs.bham.ac.uk/~aps/research/papers/pdf/SeSw-TRCSR-04-2-SymmetricMTree.pdf (Accessed : 19 February 2019)

Traina, C., Traina, A., Faloutsos, C. and Seeger, B. (2002). Fast indexing and visualization of metric data sets using slim-trees. IEEE Transactions on Knowledge and Data Engineering, 14(2), pp.244-260.

Hjaltason, G. and Samet, H. (1999). Distance browsing in spatial databases. ACM Transactions on Database Systems, 24(2), pp.265-318.

## Example Usage
```go
package main

import (
	"fmt"
	"math"

	mtree "github.com/gemilangsantiyuda/go-mtree"
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

```
