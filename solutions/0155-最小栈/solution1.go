package week11

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
