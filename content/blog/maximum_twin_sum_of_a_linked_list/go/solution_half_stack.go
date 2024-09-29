package maximum_twin_sum_of_a_linked_list

func pairSumStack(head *ListNode) int {
	stack := make([]int, 0)
	currentListElement, oddElementList := head, head

	for oddElementList != nil && oddElementList.Next != nil {
		stack = append(stack, currentListElement.Val)
		currentListElement = currentListElement.Next
		oddElementList = oddElementList.Next.Next
	}

	res := 0
	for ; currentListElement != nil; currentListElement = currentListElement.Next {
		res = max(res, currentListElement.Val+stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return res
}
