package quickunion

import (
	"fmt"

	"github.com/ivanlemeshev/algorithms-go/unionfind"
)

// QuickUnion is an implementation of union–find data type. This implementation
// uses quick union. The constructor takes O(n) time, where n is the number of
// sites. The union and find operations take O(n) time in the worst case. The
// count operation takes O(1) time.
type QuickUnion struct {
	parent []int
	count  int
}

// New initializes an empty union-find data structure with n elements from 0
// through n-1. Initially, each elements is in its own set.
func New(n int) unionfind.UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	return &QuickUnion{
		parent: parent,
		count:  n,
	}
}

// Union merges the set containing element p with the the set containing element q.
func (qu *QuickUnion) Union(p, q int) error {
	if err := qu.validate(p); err != nil {
		return err
	}

	if err := qu.validate(q); err != nil {
		return err
	}

	rootP := qu.find(p)
	rootQ := qu.find(q)

	// p and q are already in the same component
	if rootP == rootQ {
		return nil
	}

	qu.parent[rootP] = rootQ
	qu.count--

	return nil
}

// Find returns the canonical element of the set containing element p.
func (qu *QuickUnion) Find(p int) (int, error) {
	if err := qu.validate(p); err != nil {
		return 0, err
	}

	return qu.find(p), nil
}

// IsConnected returns true if the two elements are in the same set.
func (qu *QuickUnion) IsConnected(p, q int) (bool, error) {
	if err := qu.validate(p); err != nil {
		return false, err
	}

	if err := qu.validate(q); err != nil {
		return false, err
	}

	rootP := qu.find(p)
	rootQ := qu.find(q)

	return rootP == rootQ, nil
}

// Count returns the number of sets.
func (qu *QuickUnion) Count() int {
	return qu.count
}

// validate that p is a valid index.
func (qu *QuickUnion) validate(p int) error {
	n := len(qu.parent)
	if p < 0 || p >= n {
		return fmt.Errorf("index %d is not between 0 and %d", p, n-1)
	}

	return nil
}

// find returns the canonical element of the set containing element p.
func (qu *QuickUnion) find(p int) int {
	for p != qu.parent[p] {
		p = qu.parent[p]
	}

	return p
}
