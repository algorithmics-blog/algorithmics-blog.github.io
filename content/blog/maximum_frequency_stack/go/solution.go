package maximum_frequency_stack

type FreqStack struct {
	h        heap
	freqElem map[int]uint16
	curSeq   uint16
}

func Constructor() FreqStack {
	return FreqStack{
		h:        make(heap, 0),
		freqElem: make(map[int]uint16),
	}
}

func (fs *FreqStack) Push(val int) {
	fs.curSeq++
	fs.freqElem[val]++
	fs.h.add(&Item{
		freq: fs.freqElem[val],
		seq:  fs.curSeq,
		val:  val,
	})
}

func (fs *FreqStack) Pop() int {
	item := fs.h.pop()
	fs.freqElem[item.val]--
	if fs.freqElem[item.val] == 0 {
		delete(fs.freqElem, item.val)
	}

	return item.val
}

type Item struct {
	freq uint16
	seq  uint16
	val  int
}

func (i *Item) more(second *Item) bool {
	if i.freq != second.freq {
		return i.freq > second.freq
	}

	return i.seq > second.seq
}

type heap []*Item

func (h *heap) add(new *Item) {
	*h = append(*h, new)
	h.upElement(len(*h) - 1)
}

func (h *heap) pop() *Item {
	if len(*h) == 0 {
		return nil
	}

	res := (*h)[0]
	if len(*h) > 0 {
		(*h)[0] = (*h)[len(*h)-1]
		*h = (*h)[:len(*h)-1]
		h.heapify(0, len(*h)-1)
	}

	return res
}

func (h *heap) heapify(elemIdx, limit int) {
	left := elemIdx*2 + 1
	if left > limit {
		return
	}
	larger := elemIdx

	if (*h)[left].more((*h)[larger]) {
		larger = left
	}
	right := left + 1
	if right <= limit && (*h)[right].more((*h)[larger]) {
		larger = right
	}

	if larger != elemIdx {
		(*h)[larger], (*h)[elemIdx] = (*h)[elemIdx], (*h)[larger]
		h.heapify(larger, limit)
	}

}

func (h *heap) upElement(elemIdx int) {
	rootIdx := (elemIdx - 1) / 2

	if rootIdx >= 0 && (*h)[elemIdx].more((*h)[rootIdx]) {
		(*h)[rootIdx], (*h)[elemIdx] = (*h)[elemIdx], (*h)[rootIdx]
		h.upElement(rootIdx)
	}
}
