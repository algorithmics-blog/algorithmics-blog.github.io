package unique_value_structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func containsValue(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Test_Store(t *testing.T) {
	t.Run("Insert duplicates", func(t *testing.T) {
		store := New()

		store.Insert(1)
		store.Insert(2)
		store.Insert(1)

		assert.Equal(t, []int{1, 2}, store.getValues())
	})

	t.Run("Remove value", func(t *testing.T) {
		store := New()

		store.Insert(1)
		store.Insert(2)
		store.Insert(3)
		store.Remove(2)

		assert.Equal(t, []int{1, 3}, store.getValues())
	})

	t.Run("Remove value", func(t *testing.T) {
		store := New()

		store.Insert(1)
		store.Insert(2)
		store.Insert(3)

		value := store.GetRandom()

		if !containsValue(store.getValues(), value) {
			t.Errorf("GetRandom returns value which is not in store")
		}
	})
}
