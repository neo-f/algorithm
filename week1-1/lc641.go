package week1_1

// 641. 设计循环双端队列 https://leetcode-cn.com/problems/design-circular-deque/

type MyCircularDeque struct {
	data []int
	// 指向队列头部第1个有效数据的位置；
	head int
	// 指向队列尾部的**下一个位置**，即下一个从队尾入队元素的位置。
	tail int
	// 当前大小
	size int
	// 容量
	cap int
}

// NewMyCircularDeque Initialize your data structure here. Set the size of the deque to be k.
func NewMyCircularDeque(k int) MyCircularDeque {
	return MyCircularDeque{
		data: make([]int, k),
		cap:  k,
		size: 0,
		head: 0,
	}
}

// InsertFront /** Adds an item at the front of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.head = (d.cap + d.head - 1) % d.cap
	d.data[d.head] = value
	d.size++
	return true
}

// InsertLast /** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.data[d.tail] = value
	d.tail = (d.cap + d.tail + 1) % d.cap
	d.size++
	return true
}

// DeleteFront /** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	// 从头删除元素，head往右走
	d.head = (d.head + 1) % d.cap
	d.size--
	return true
}

// DeleteLast /** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.tail = (d.cap + d.tail - 1) % d.cap
	d.size--
	return true
}

// GetFront /** Get the front item from the deque. */
func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[d.head]
}

// GetRear /** Get the last item from the deque. */
func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.data[(d.cap+d.tail-1)%d.cap]
}

// IsEmpty /** Checks whether the circular deque is empty or not. */
func (d *MyCircularDeque) IsEmpty() bool {
	return d.size == 0
}

// IsFull /** Checks whether the circular deque is full or not. */
func (d *MyCircularDeque) IsFull() bool {
	return d.size == d.cap
}
