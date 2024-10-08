package maximum_twin_sum_of_a_linked_list

func valuesSliceToList(values []int) *ListNode {
	cur := &ListNode{
		Val: values[0],
	}

	head := cur

	for i := 1; i < len(values); i++ {
		next := &ListNode{
			Val: values[i],
		}
		cur.Next = next
		cur = next
	}

	return head
}
