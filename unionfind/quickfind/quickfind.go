package unionfind

import (
	"fmt"

	"github.com/ivanlemeshev/algorithms-go/unionfind"
)

// QuickFind is an implementation of union–find data type. This implementation
// uses quick find. The constructor takes O(n) time, where n is the number of
// sites. The find, connected and count operations take O(1) time; the union
// operation takes O(n) time.
type QuickFind struct {
	id    []int
	count int
}

// New initializes an empty union-find data structure with n elements from 0
// through n-1. Initially, each elements is in its own set.
func New(n int) unionfind.UnionFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return &QuickFind{
		id:    id,
		count: n,
	}
}

// Union merges the set containing element p with the the set containing element q.
func (qf *QuickFind) Union(p, q int) error {
	if err := qf.validate(p); err != nil {
		return err
	}

	if err := qf.validate(q); err != nil {
		return err
	}

	pID := qf.id[p]
	qID := qf.id[q]

	// p and q are already in the same component
	if pID == qID {
		return nil
	}

	for i := range qf.id {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}

	qf.count--

	return nil
}

// Find returns the canonical element of the set containing element p.
func (qf *QuickFind) Find(p int) (int, error) {
	if err := qf.validate(p); err != nil {
		return 0, err
	}

	return qf.id[p], nil
}

// IsConnected returns true if the two elements are in the same set.
func (qf *QuickFind) IsConnected(p, q int) (bool, error) {
	if err := qf.validate(p); err != nil {
		return false, err
	}

	if err := qf.validate(q); err != nil {
		return false, err
	}

	return qf.id[p] == qf.id[q], nil
}

// Count returns the number of sets.
func (qf *QuickFind) Count() int {
	return qf.count
}

// validate that p is a valid index.
func (qf *QuickFind) validate(p int) error {
	n := len(qf.id)
	if p < 0 || p >= n {
		return fmt.Errorf("index %d is not between 0 and %d", p, n-1)
	}

	return nil
}
