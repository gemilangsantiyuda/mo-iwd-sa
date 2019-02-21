package mtree

// Object which is stored in the m-tree must have unique ID and a way to get them
type Object interface {
	GetID() string
}
