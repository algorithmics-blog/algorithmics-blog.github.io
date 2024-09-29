package maximum_twin_sum_of_a_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pairSum(t *testing.T) {
	testCases := []struct {
		name string
		list []int
		out  int
	}{
		{
			name: "[5,4,2,1]",
			list: []int{5, 4, 2, 1},
			out:  6,
		},
		{
			name: "[4,2,2,3]",
			list: []int{4, 2, 2, 3},
			out:  7,
		},
		{
			name: "[1,100000]",
			list: []int{1, 100000},
			out:  100001,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cur := &ListNode{
				Val: testCase.list[0],
			}
			head := cur
			for i := 1; i < len(testCase.list); i++ {
				next := &ListNode{
					Val: testCase.list[i],
				}
				cur.Next = next
				cur = next
			}

			res := pairSum(head)
			assert.Equal(t, res, testCase.out)
		})
	}
}
