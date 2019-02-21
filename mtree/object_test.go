package mtree_test

import (
	"fmt"
	"math/rand"
	"strconv"
)

func generateObjectList(n int) []*object {

	// currentSeed := int64(time.Now().Nanosecond())
	currentSeed := 851359177
	fmt.Println("Current Seed ", currentSeed)
	rand.Seed(int64(currentSeed))
	objList := make([]*object, 0)
	for idx := 1; idx <= n; idx++ {
		x := rand.Float64() * 10
		y := rand.Float64() * 10
		newObj := &object{
			id: strconv.Itoa(idx),
			x:  x,
			y:  y,
		}
		objList = append(objList, newObj)
	}
	return objList
}
