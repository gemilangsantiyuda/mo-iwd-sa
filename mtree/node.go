package mtree

type node interface {
	isLeaf() bool
	isUnderFlown(int) bool
	getEntryList() []entry
	insertEntry(entry)
	updateRadius()
	entry
}
