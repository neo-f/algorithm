package main

import "container/heap"

type MedianFinder struct {
	// 奇数的值存在max
	maxHeap *MaxHeap
	minHeap *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: &(MaxHeap{}),
		minHeap: &(MinHeap{}),
	}
}

func (f *MedianFinder) AddNum(num int) {
	// 1. 统统放到大顶堆里面
	heap.Push(f.maxHeap, num)
	// 2. 倒腾一遍是要保证小顶堆中的所有元素永远比大顶堆中的元素大
	heap.Push(f.minHeap, heap.Pop(f.maxHeap))
	// 3. 倒腾完之后，如果大顶堆的数少了，需要再倒腾回去
	// 要保证大堆大于等于小堆的个数
	if f.maxHeap.Len() < f.minHeap.Len() {
		heap.Push(f.maxHeap, heap.Pop(f.minHeap))
	}
	// 假设放 1, 3, 2, 5, 4
	// 1.       [1], []              -> [], [1]              -> [1], []
	// 3.       [1, 3], []           -> [1], [3]
	// 2.       [1, 2], [3]          -> [1], [2, 3]          -> [1, 2], [3]
	// 5.       [1, 2, 5], [3]       -> [1, 2], [3, 5]
	// 4.       [1, 2, 4], [3, 5]    -> [1, 2], [3, 4, 5]    -> [1, 2, 3], [4, 5]
}

func (f *MedianFinder) FindMedian() float64 {
	if f.maxHeap.Len() == f.minHeap.Len() {
		return float64(f.maxHeap.Top()+f.minHeap.Top()) / 2
	}
	return float64(f.maxHeap.Top())

}

type MinHeap []int

func (m MinHeap) Len() int            { return len(m) }
func (m MinHeap) Less(i, j int) bool  { return m[i] < m[j] }
func (m MinHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m MinHeap) Top() int            { return m[0] }
func (m *MinHeap) Push(x interface{}) { *m = append(*m, x.(int)); heap.Fix(m, m.Len()-1) }
func (m *MinHeap) Pop() interface{} {
	e := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return e
}

type MaxHeap []int

func (m MaxHeap) Len() int            { return len(m) }
func (m MaxHeap) Less(i, j int) bool  { return m[i] > m[j] }
func (m MaxHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Top() int            { return m[0] }
func (m *MaxHeap) Push(x interface{}) { *m = append(*m, x.(int)); heap.Fix(m, m.Len()-1) }
func (m *MaxHeap) Pop() interface{} {
	e := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return e
}
