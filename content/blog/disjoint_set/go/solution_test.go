package disjoint_set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DisjointSet(t *testing.T) {
	set := NewDisjointSet(8)

	set.union(0, 2)
	set.union(0, 1)
	set.union(2, 3)
	set.union(1, 3)

	set.union(4, 5)
	set.union(4, 6)
	set.union(5, 6)

	t.Run("Create disjoint set", func(t *testing.T) {
		assert.Equal(t, []int{0, 0, 0, 0, 4, 4, 4, 7}, set.nodes)
	})

	t.Run("Find direct connection", func(t *testing.T) {
		assert.Equal(t, set.find(3), 0)
	})

	t.Run("Find transitive connection", func(t *testing.T) {
		assert.Equal(t, set.find(5), 4)
	})

	t.Run("Find in set with one element", func(t *testing.T) {
		assert.Equal(t, set.find(7), 7)
	})

	t.Run("Check direct connection", func(t *testing.T) {
		assert.Equal(t, set.connected(4, 5), true)
	})

	t.Run("Check transitive connection", func(t *testing.T) {
		assert.Equal(t, set.connected(0, 3), true)
	})

	t.Run("Check root connection", func(t *testing.T) {
		assert.Equal(t, set.connected(7, 7), true)
	})

	t.Run("Check not connected elements", func(t *testing.T) {
		assert.Equal(t, set.connected(0, 7), false)
	})

	t.Run("Check union", func(t *testing.T) {
		set.union(3, 4)
		assert.Equal(t, set.nodes, []int{0, 0, 0, 0, 0, 0, 0, 7})
		assert.Equal(t, set.connected(0, 6), true)
	})
}
