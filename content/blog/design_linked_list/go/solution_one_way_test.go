package design_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func buildOneWayValues(list OneWayLinkedList) []int {
	current := list.head
	values := make([]int, 0)
	for current != nil {
		values = append(values, current.Val)
		current = current.Next
	}

	return values
}

func Test_OneWayLinkedList(t *testing.T) {
	t.Run("scenario 1", func(t *testing.T) {
		list := NewOneWayLinkedList()

		list.AddAtHead(7)
		list.AddAtHead(2)
		list.AddAtHead(1)
		list.AddAtIndex(3, 0)
		list.DeleteAtIndex(2)
		list.AddAtHead(6)
		list.AddAtTail(4)
		list.AddAtHead(4)
		list.AddAtIndex(5, 0)
		list.AddAtHead(6)

		values := buildOneWayValues(list)

		assert.Equal(t, []int{6, 4, 6, 1, 2, 0, 0, 4}, values)
		assert.Equal(t, 8, list.size)
	})

	t.Run("scenario 2", func(t *testing.T) {
		list := NewOneWayLinkedList()

		list.AddAtHead(1)
		list.AddAtTail(3)
		list.AddAtIndex(1, 2)
		list.DeleteAtIndex(0)

		values := buildOneWayValues(list)

		assert.Equal(t, []int{2, 3}, values)
		assert.Equal(t, 2, list.size)
	})

	t.Run("scenario 3", func(t *testing.T) {
		list := NewOneWayLinkedList()

		list.AddAtHead(1)
		list.AddAtTail(3)
		list.AddAtIndex(1, 2)

		assert.Equal(t, 2, list.Get(1))

		list.DeleteAtIndex(1)

		assert.Equal(t, 3, list.Get(1))
		assert.Equal(t, -1, list.Get(3))

		list.DeleteAtIndex(3)
		list.DeleteAtIndex(0)

		assert.Equal(t, 3, list.Get(0))

		list.DeleteAtIndex(0)

		assert.Equal(t, -1, list.Get(0))

		values := buildOneWayValues(list)

		assert.Equal(t, []int{}, values)
		assert.Equal(t, 0, list.size)
	})
}
