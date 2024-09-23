package design_linked_list

type MyLinkedListNode struct {
	Val  int
	Next *MyLinkedListNode
}

type MyLinkedList struct {
	head *MyLinkedListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		size: 0,
	}
}

func (l *MyLinkedList) Get(index int) int {
	if index >= l.size || index < 0 {
		return -1
	}

	curr := l.head
	for i := 1; i <= index; i++ {
		curr = curr.Next
	}

	return curr.Val
}

func (l *MyLinkedList) findPrevNode(index int) *MyLinkedListNode {
	curr := l.head

	for i := 0; i < index-1; i++ {
		if curr != nil {
			curr = curr.Next
		}
	}

	return curr
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index > l.size || index < 0 {
		return
	}

	l.size++

	// Если нужно добавить элемент в голову
	if index == 0 {
		l.head = &MyLinkedListNode{
			Val:  val,
			Next: l.head,
		}
		return
	}

	prev := l.findPrevNode(index)

	// Создаем новый узел и обновляем связи
	prev.Next = &MyLinkedListNode{
		Val:  val,
		Next: prev.Next,
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index > l.size-1 {
		return
	}

	l.size--

	if index == 0 {
		l.head = l.head.Next
		return
	}

	prev := l.findPrevNode(index)

	prev.Next = prev.Next.Next
}
