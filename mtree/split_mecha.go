package mtree

type splitMecha interface {
	split([]entry) (node, node)
}
