package quickunion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnion(t *testing.T) {
	uf := New(10)

	t.Run("index is less than 0", func(t *testing.T) {
		err := uf.Union(-1, 0)
		assert.Equal(t, "index -1 is not between 0 and 9", err.Error())
	})

	t.Run("index is more or equal to n", func(t *testing.T) {
		err := uf.Union(0, 10)
		assert.Equal(t, "index 10 is not between 0 and 9", err.Error())
	})

	t.Run("index between 0 and n", func(t *testing.T) {
		assert.Nil(t, uf.Union(4, 3))
		assert.Nil(t, uf.Union(3, 8))
		assert.Nil(t, uf.Union(6, 5))
		assert.Nil(t, uf.Union(9, 4))
	})

	t.Run("union twice", func(t *testing.T) {
		assert.Nil(t, uf.Union(4, 3))
	})
}

func TestCount(t *testing.T) {
	uf := New(5)

	// 1 2 3 4 0
	assert.Equal(t, 5, uf.Count())

	// 1 2 3 4-0
	assert.Nil(t, uf.Union(0, 4))
	assert.Equal(t, 4, uf.Count())

	// 1 2-3 4-0
	assert.Nil(t, uf.Union(2, 3))
	assert.Equal(t, 3, uf.Count())

	// 1-2-3 4-0
	assert.Nil(t, uf.Union(1, 2))
	assert.Equal(t, 2, uf.Count())

	// 1-2-3-4-0
	assert.Nil(t, uf.Union(4, 3))
	assert.Equal(t, 1, uf.Count())
}

func TestIsConnected(t *testing.T) {
	uf := New(5)

	assert.Nil(t, uf.Union(0, 4))
	assert.Nil(t, uf.Union(2, 3))
	assert.Nil(t, uf.Union(1, 2))

	// 1-2-3 4-0

	t.Run("index is less than 0", func(t *testing.T) {
		isConnected, err := uf.IsConnected(-1, 4)
		assert.False(t, isConnected)
		assert.Equal(t, "index -1 is not between 0 and 4", err.Error())
	})

	t.Run("index is more or equal to n", func(t *testing.T) {
		isConnected, err := uf.IsConnected(0, 5)
		assert.False(t, isConnected)
		assert.Equal(t, "index 5 is not between 0 and 4", err.Error())
	})

	t.Run("index between 0 and n", func(t *testing.T) {
		isConnected, err := uf.IsConnected(4, 0)
		assert.True(t, isConnected)
		assert.Nil(t, err)

		isConnected, err = uf.IsConnected(1, 3)
		assert.True(t, isConnected)
		assert.Nil(t, err)

		isConnected, err = uf.IsConnected(1, 2)
		assert.True(t, isConnected)
		assert.Nil(t, err)

		isConnected, err = uf.IsConnected(2, 3)
		assert.True(t, isConnected)
		assert.Nil(t, err)

		isConnected, err = uf.IsConnected(3, 4)
		assert.False(t, isConnected)
		assert.Nil(t, err)

		isConnected, err = uf.IsConnected(1, 0)
		assert.False(t, isConnected)
		assert.Nil(t, err)
	})
}

func TestFind(t *testing.T) {
	uf := New(5)

	t.Run("index is less than 0", func(t *testing.T) {
		v, err := uf.Find(-1)
		assert.Equal(t, 0, v)
		assert.Equal(t, "index -1 is not between 0 and 4", err.Error())
	})

	t.Run("index is more or equal to n", func(t *testing.T) {
		v, err := uf.Find(5)
		assert.Equal(t, 0, v)
		assert.Equal(t, "index 5 is not between 0 and 4", err.Error())
	})

	t.Run("index between 0 and n", func(t *testing.T) {
		assert.Nil(t, uf.Union(0, 4))

		v, err := uf.Find(0)
		assert.Equal(t, 4, v)
		assert.Nil(t, err)

		v, err = uf.Find(4)
		assert.Equal(t, 4, v)
		assert.Nil(t, err)

		assert.Nil(t, uf.Union(2, 3))
		assert.Nil(t, uf.Union(1, 2))

		v, err = uf.Find(1)
		assert.Equal(t, 3, v)
		assert.Nil(t, err)

		v, err = uf.Find(2)
		assert.Equal(t, 3, v)
		assert.Nil(t, err)

		v, err = uf.Find(3)
		assert.Equal(t, 3, v)
		assert.Nil(t, err)
	})
}
