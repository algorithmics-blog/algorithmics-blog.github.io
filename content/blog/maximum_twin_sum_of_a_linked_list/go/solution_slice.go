package maximum_twin_sum_of_a_linked_list

func pairSumSlice(head *ListNode) int {
	list := []int{head.Val}
	for cur := head.Next; cur != nil; cur = cur.Next {
		list = append(list, cur.Val)
	}

	res := 0
	for i := 0; i < len(list)/2; i++ {
		j := len(list) - i - 1
		res = max(res, list[i]+list[j])
	}

	return res
}
