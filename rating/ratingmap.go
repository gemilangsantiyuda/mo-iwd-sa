package rating

// Map to store rating of order (based on its userID) toward a kitchen
type Map map[string]float64

type order interface {
	GetUserID() string
}

type kitchen interface {
	GetID() string
}
