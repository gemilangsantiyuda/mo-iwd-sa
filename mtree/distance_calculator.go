package mtree

// calculator is the distance calculator method of the objects
type distanceCalculator interface {
	GetDistance(Origin Object, Destination Object) float64
}
