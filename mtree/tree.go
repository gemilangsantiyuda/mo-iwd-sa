package mtree

// Tree of the mtree
type Tree struct {
	root        node
	maxEntry    int
	minEntry    int
	distCalc    distanceCalculator
	ObjectCount int
	splitMecha  splitMecha
}
