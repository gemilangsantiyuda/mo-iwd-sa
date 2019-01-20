package coordinate

// Coordinate struct to store Latitude and Longitude of Kulina orders
// Lat lon are using google maps geodatum
type Coordinate struct {
	Latitude, Longitude float64
}

// GetCoordinate return the lat lon coordinate of this object (current is also the coordinate struct)
func (coord *Coordinate) GetCoordinate() *Coordinate {
	return coord
}
