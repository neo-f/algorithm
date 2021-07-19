package main

import "fmt"

// 优先级队列解法
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewPriorityQueue()
	for i := 0; i < k-1; i++ {
		queue.Push(entry{Idx: i, Val: nums[i]})
	}
	var ans []int
	// [[1, 2, 3], 4, 5, 6]
	for i := k - 1; i < len(nums); i++ {
		queue.Push(entry{Idx: i, Val: nums[i]})
		//i:4   idx:2
		for queue.Top().Idx < i-k+1 {
			queue.Pop()
		}
		fmt.Println(queue)
		ans = append(ans, queue.Top().Val)
	}
	return ans
}

type entry struct {
	Idx int
	Val int
}
type PriorityQueue struct {
	data []entry
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{data: make([]entry, 0)}
}
func (q *PriorityQueue) IsEmpty() bool {
	return len(q.data) == 0
}
func (q *PriorityQueue) Top() entry {
	return q.data[0]
}
func (q *PriorityQueue) Push(e entry) {
	q.data = append(q.data, e)
	// HeapifyUp
	curr := len(q.data) - 1
	for curr > 0 {
		father := (curr - 1) / 2
		if q.data[father].Val < q.data[curr].Val {
			q.data[father], q.data[curr] = q.data[curr], q.data[father]
			curr = father
		} else {
			break
		}
	}
}
func (q *PriorityQueue) Pop() entry {
	max := q.data[0]
	q.data[0] = q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]
	// HeapDown
	idx := 0
	child := idx*2 + 1
	for child < len(q.data) {
		other := child + 1
		if other < len(q.data) && q.data[child].Val < q.data[other].Val {
			child = other
		}
		if q.data[idx].Val < q.data[child].Val {
			q.data[idx], q.data[child] = q.data[child], q.data[idx]
			idx = child
			child = idx*2 + 1
		} else {
			break
		}
	}
	return max
}

func main() {
	maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3)
}
