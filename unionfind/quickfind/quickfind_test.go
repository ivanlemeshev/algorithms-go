package unionfind

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFind(t *testing.T) {
	uf := New(10)

	t.Run("index is less than 0", func(t *testing.T) {
		err := uf.Union(-1, 0)
		assert.Equal(t, "index -1 is not between 0 and 9", err.Error())
	})

	t.Run("index is more or equal to n", func(t *testing.T) {
		err := uf.Union(0, 10)
		assert.Equal(t, "index 10 is not between 0 and 9", err.Error())
	})

	union := [][]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
	}

	for _, p := range union {
		name := fmt.Sprintf("union (%d,%d)", p[0], p[1])

		t.Run(name, func(t *testing.T) {
			assert.Nil(t, uf.Union(p[0], p[1]))
		})
	}

	connected := [][]int{
		{3, 4},
		{4, 3},
		{3, 8},
		{8, 3},
		{4, 8},
		{8, 4},
		{0, 5},
		{5, 6},
		{6, 1},
		{1, 2},
		{2, 7},
		{0, 7},
		{4, 9},
		{8, 9},
	}

	for _, p := range connected {
		name := fmt.Sprintf("connected (%d,%d)", p[0], p[1])

		t.Run(name, func(t *testing.T) {
			isConnected, err := uf.IsConnected(p[0], p[1])
			assert.True(t, isConnected)
			assert.Nil(t, err)
		})
	}

	notConnected := [][]int{
		{0, 3},
		{0, 4},
		{0, 8},
		{0, 9},
		{1, 3},
		{1, 4},
		{1, 8},
		{1, 9},
		{2, 3},
		{2, 4},
		{2, 8},
		{2, 9},
		{5, 3},
		{5, 4},
		{5, 8},
		{5, 9},
		{6, 3},
		{6, 4},
		{6, 8},
		{6, 9},
		{7, 3},
		{7, 4},
		{7, 8},
		{7, 9},
	}

	for _, p := range notConnected {
		name := fmt.Sprintf("not connected (%d,%d)", p[0], p[1])

		t.Run(name, func(t *testing.T) {
			isConnected, err := uf.IsConnected(p[0], p[1])
			assert.False(t, isConnected)
			assert.Nil(t, err)
		})
	}
}
