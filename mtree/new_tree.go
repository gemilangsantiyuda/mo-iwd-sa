package mtree

// NewTree initiates mtree
func NewTree(minEntry, maxEntry int, distCalc distanceCalculator) *Tree {
	splitMecha := &splitMST{
		distCalc: distCalc,
	}
	root := &leaf{}
	return &Tree{
		root:       root,
		minEntry:   minEntry,
		maxEntry:   maxEntry,
		distCalc:   distCalc,
		splitMecha: splitMecha,
	}
}
