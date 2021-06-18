package week1

import "strconv"

// 20. 有效的括号 https://leetcode-cn.com/problems/valid-parentheses/
type runeStack []rune

func (s *runeStack) Push(r rune) {
	*s = append(*s, r)
}
func (s *runeStack) Pop() (r rune, ok bool) {
	length := len(*s)
	if length == 0 {
		return 0, false
	}
	r = (*s)[length-1]
	*s = (*s)[:length-1]
	return r, true
}

func isValid(s string) bool {
	stack := runeStack{}
	for _, token := range s {
		switch token {
		case '(', '[', '{':
			stack.Push(token)
		case ')':
			if r, ok := stack.Pop(); !ok || r != '(' {
				return false
			}
		case ']':
			if r, ok := stack.Pop(); !ok || r != '[' {
				return false
			}
		case '}':
			if r, ok := stack.Pop(); !ok || r != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}

// 155. 最小栈 https://leetcode-cn.com/problems/min-stack/

type MinStack struct {
	Origin []int
	Min    []int
}

func NewMinStack() MinStack {
	return MinStack{
		Origin: nil,
		Min:    nil,
	}
}

func (s *MinStack) Push(val int) {
	s.Origin = append(s.Origin, val)
	var min int
	if len(s.Min) == 0 {
		min = val
	} else {
		min = s.GetMin()
		if val < min {
			min = val
		}
	}
	s.Min = append(s.Min, min)
}

func (s *MinStack) Pop() {
	s.Origin = s.Origin[:len(s.Origin)-1]
	s.Min = s.Min[:len(s.Min)-1]
}

func (s *MinStack) Top() int {
	return s.Origin[len(s.Origin)-1]
}

func (s *MinStack) GetMin() int {
	return s.Min[len(s.Min)-1]
}

//150. 逆波兰表达式求值 https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
//输入：tokens = ["2","1","+","3","*"]
//输出：9
type stringStack []string

func (s *stringStack) Push(r string) {
	*s = append(*s, r)
}
func (s *stringStack) Pop() string {
	idx := len(*s) - 1
	r := (*s)[idx]
	*s = (*s)[:idx]
	return r
}

func calc(b, a, op string) int {
	ib, _ := strconv.Atoi(b)
	ia, _ := strconv.Atoi(a)
	switch op {
	case "+":
		return ia + ib
	case "-":
		return ia - ib
	case "*":
		return ia * ib
	case "/":
		return ia / ib
	}
	return 0
}

func evalRPN(tokens []string) int {
	stack := stringStack{}
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			stack.Push(strconv.Itoa(calc(stack.Pop(), stack.Pop(), token)))
		default:
			stack.Push(token)
		}
	}
	result, _ := strconv.Atoi(stack.Pop())
	return result
}

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


// 141. 环形链表 https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		head = head.Next
		if fast == head{
			return true
		}
	}
	return false
}