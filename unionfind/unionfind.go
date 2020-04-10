package unionfind

// UnionFind interface represents a union–find data type (also known as the
// disjoint-sets data type. It supports the classic union and find operations,
// along with a count operation that returns the total number of sets.
//
// The union-find data type models a collection of sets containing n elements,
// with each element in exactly one set. The elements are named from 0 through
// n–1. Initially, there are n sets, with each element in its own set. The
// canonical element of a set (also known as the root, identifier, leader, or
// set representative) is one distinguished element in the set.
//
// The canonical element of a set can change only when the set itself changes
// during a call to union; it cannot change during a call to either find or
// count.
type UnionFind interface {
	// Union merges the set containing element p with the the set containing
	// element q.
	Union(p, q int) error

	// Find returns the canonical element of the set containing element p.
	Find(p int) (int, error)

	// IsConnected returns true if the two elements are in the same set.
	IsConnected(p, q int) (bool, error)

	// Count returns the number of sets.
	Count() int
}
