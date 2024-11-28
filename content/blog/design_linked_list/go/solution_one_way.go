package design_linked_list

type OneWayLinkedListNode struct {
	Val  int
	Next *OneWayLinkedListNode
}

type OneWayLinkedList struct {
	head *OneWayLinkedListNode
	size int
}

func NewOneWayLinkedList() OneWayLinkedList {
	return OneWayLinkedList{
		size: 0,
	}
}

func (l *OneWayLinkedList) Get(index int) int {
	if index >= l.size || index < 0 {
		return -1
	}

	curr := l.head
	for i := 1; i <= index; i++ {
		curr = curr.Next
	}

	return curr.Val
}

func (l *OneWayLinkedList) findPrevNode(index int) *OneWayLinkedListNode {
	curr := l.head

	for i := 0; i < index-1; i++ {
		if curr != nil {
			curr = curr.Next
		}
	}

	return curr
}

func (l *OneWayLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *OneWayLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *OneWayLinkedList) AddAtIndex(index int, val int) {
	if index > l.size || index < 0 {
		return
	}

	l.size++

	// Если нужно добавить элемент в голову
	if index == 0 {
		l.head = &OneWayLinkedListNode{
			Val:  val,
			Next: l.head,
		}
		return
	}

	prev := l.findPrevNode(index)

	// Создаем новый узел и обновляем связи
	prev.Next = &OneWayLinkedListNode{
		Val:  val,
		Next: prev.Next,
	}
}

func (l *OneWayLinkedList) DeleteAtIndex(index int) {
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
