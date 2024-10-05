package design_linked_list

type TwoWayLinkedListNode struct {
	Val  int
	Next *TwoWayLinkedListNode
	Prev *TwoWayLinkedListNode
}

type TwoWayLinkedList struct {
	head *TwoWayLinkedListNode
	tail *TwoWayLinkedListNode
	size int
}

func NewTwoWayLinkedList() TwoWayLinkedList {
	head := &TwoWayLinkedListNode{
		Val: 0,
	}
	tail := &TwoWayLinkedListNode{
		Val: 0,
	}

	head.Next = tail
	tail.Prev = head

	return TwoWayLinkedList{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (l *TwoWayLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}

	curr := l.head

	if index+1 < l.size-index {
		for i := 0; i <= index; i++ {
			curr = curr.Next
		}

	} else {
		curr := l.tail
		for i := 0; i < l.size-index; i++ {
			curr = curr.Prev
		}
	}

	return curr.Val
}

func (l *TwoWayLinkedList) AddAtHead(val int) {
	prev := l.head
	next := l.head.Next

	newNode := &TwoWayLinkedListNode{
		Val: val,
	}

	newNode.Prev = prev
	newNode.Next = next
	prev.Next = newNode
	next.Prev = newNode

	l.size++
}

func (l *TwoWayLinkedList) AddAtTail(val int) {
	next := l.tail
	prev := l.tail.Prev

	newNode := &TwoWayLinkedListNode{
		Val: val,
	}

	newNode.Prev = prev
	newNode.Next = next
	prev.Next = newNode
	next.Prev = newNode

	l.size++
}

func (l *TwoWayLinkedList) AddAtIndex(index int, val int) {
	if index > l.size {
		return
	}

	if index < 0 {
		index = 0
	}

	var prev *TwoWayLinkedListNode
	var next *TwoWayLinkedListNode

	if index < l.size-index {
		prev = l.head
		for i := 0; i < index; i++ {
			prev = prev.Next
		}
		next = prev.Next
	} else {
		next = l.tail
		for i := 0; i < l.size-index; i++ {
			next = next.Prev
		}
		prev = next.Prev
	}

	newNode := &TwoWayLinkedListNode{
		Val: val,
	}

	newNode.Prev = prev
	newNode.Next = next
	prev.Next = newNode
	next.Prev = newNode

	l.size++
}

func (l *TwoWayLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > l.size-1 {
		return
	}

	var prev *TwoWayLinkedListNode
	var next *TwoWayLinkedListNode

	if index < l.size-index {
		prev = l.head
		for i := 0; i < index; i++ {
			prev = prev.Next
		}
		next = prev.Next.Next
	} else {
		next = l.tail
		for i := 0; i < l.size-index-1; i++ {
			next = next.Prev
		}
		prev = next.Prev.Prev
	}

	prev.Next = next
	next.Prev = prev
	l.size--
}
