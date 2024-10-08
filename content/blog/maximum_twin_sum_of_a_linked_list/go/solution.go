package maximum_twin_sum_of_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	firstHalfList, oddElementList := head, head

	for oddElementList != nil && oddElementList.Next != nil {
		firstHalfList = firstHalfList.Next
		oddElementList = oddElementList.Next.Next
	}

	revers := reverseList(firstHalfList)

	maxSum := 0
	for revers != nil {
		maxSum = max(maxSum, head.Val+revers.Val)
		head, revers = head.Next, revers.Next
	}
	return maxSum
}

func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next, newHead = newHead, head
		head = next
	}

	return newHead
}
