package reverse_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	testCases := []struct {
		name string
		list *ListNode
		out  *ListNode
	}{
		{
			name: "[1, 2, 3, 4, 5]",
			list: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			out: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			name: "[1]",
			list: &ListNode{
				Val: 1,
			},
			out: &ListNode{
				Val: 1,
			},
		},
		{
			name: "[]",
			list: &ListNode{},
			out:  &ListNode{},
		},
		{
			name: "nil",
			list: nil,
			out:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			reversed := reverseList(testCase.list)
			assert.Equal(t, reversed, testCase.out)
		})
	}
}
