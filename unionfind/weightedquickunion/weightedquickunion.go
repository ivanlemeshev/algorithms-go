package quickunion

import (
	"fmt"

	"github.com/ivanlemeshev/algorithms-go/unionfind"
)

// WeightedQuickUnion is an implementation of union–find data type. This
// implementation uses weighted quick union by size (without path compression).
// The constructor takes O(n), where n is the number of elements. The union and
// find operations  take O(log n) time in the worst case. The count operation
// takes O(1) time.
type WeightedQuickUnion struct {
	parent []int
	size   []int
	count  int
}

// New initializes an empty union-find data structure with n elements from 0
// through n-1. Initially, each elements is in its own set.
func New(n int) unionfind.UnionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = i
	}

	return &WeightedQuickUnion{
		parent: parent,
		size:   size,
		count:  n,
	}
}

// Union merges the set containing element p with the the set containing element q.
func (wqu *WeightedQuickUnion) Union(p, q int) error {
	if err := wqu.validate(p); err != nil {
		return err
	}

	if err := wqu.validate(q); err != nil {
		return err
	}

	rootP := wqu.find(p)
	rootQ := wqu.find(q)

	// p and q are already in the same component
	if rootP == rootQ {
		return nil
	}

	// make smaller root point to larger one
	if wqu.size[rootP] < wqu.size[rootQ] {
		wqu.parent[rootP] = rootQ
		wqu.size[rootQ] += wqu.size[rootP]
	} else {
		wqu.parent[rootQ] = rootP
		wqu.size[rootP] += wqu.size[rootQ]
	}

	wqu.count--

	return nil
}

// Find returns the canonical element of the set containing element p.
func (wqu *WeightedQuickUnion) Find(p int) (int, error) {
	if err := wqu.validate(p); err != nil {
		return 0, err
	}

	return wqu.find(p), nil
}

// IsConnected returns true if the two elements are in the same set.
func (wqu *WeightedQuickUnion) IsConnected(p, q int) (bool, error) {
	if err := wqu.validate(p); err != nil {
		return false, err
	}

	if err := wqu.validate(q); err != nil {
		return false, err
	}

	rootP := wqu.find(p)
	rootQ := wqu.find(q)

	return rootP == rootQ, nil
}

// Count returns the number of sets.
func (wqu *WeightedQuickUnion) Count() int {
	return wqu.count
}

// validate that p is a valid index.
func (wqu *WeightedQuickUnion) validate(p int) error {
	n := len(wqu.parent)
	if p < 0 || p >= n {
		return fmt.Errorf("index %d is not between 0 and %d", p, n-1)
	}

	return nil
}

// find returns the canonical element of the set containing element p.
func (wqu *WeightedQuickUnion) find(p int) int {
	for p != wqu.parent[p] {
		p = wqu.parent[p]
	}

	return p
}
