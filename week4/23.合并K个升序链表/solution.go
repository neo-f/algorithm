package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := NewBinaryHeap()
	for _, node := range lists {
		if node != nil {
			heap.Push(node)
		}
	}
	dummy := ListNode{}
	tail := &dummy

	// 每次取一个、放一个
	for !heap.IsEmpty() {
		// 堆中的最小值
		val := heap.Pop()
		tail.Next = val
		tail = tail.Next
		if val.Next != nil {
			heap.Push(val.Next)
		}
	}

	return dummy.Next
}

// 实现一个小根二叉堆
type binaryHeap struct {
	data []*ListNode
}

func NewBinaryHeap() *binaryHeap {
	return &binaryHeap{data: make([]*ListNode, 0)}
}

func (h *binaryHeap) Push(node *ListNode) {
	h.data = append(h.data, node)
	h.HeapifyUp(len(h.data) - 1)
}

func (h *binaryHeap) Pop() *ListNode {
	ret := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.HeapifyDown(0)
	return ret
}
func (h *binaryHeap) IsEmpty() bool {
	return len(h.data) == 0
}

func (h *binaryHeap) HeapifyUp(idx int) {
	for idx > 0 {
		p := (idx - 1) / 2
		if h.data[p].Val > h.data[idx].Val {
			h.data[p], h.data[idx] = h.data[idx], h.data[p]
			idx = p
		} else {
			break
		}
	}
}
func (h *binaryHeap) HeapifyDown(idx int) {
	child := idx*2 + 1
	for child < len(h.data) {
		other := child + 1
		if other < len(h.data) && h.data[other].Val < h.data[child].Val {
			child = other
		}
		if h.data[idx].Val > h.data[child].Val {
			h.data[idx], h.data[child] = h.data[child], h.data[idx]
			idx = child
			child = idx*2 + 1
		} else {
			break
		}
	}
}
