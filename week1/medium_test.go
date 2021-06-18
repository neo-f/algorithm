package week1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, isValid("{}"), true)
	assert.Equal(t, isValid("{[]}"), true)
	assert.Equal(t, isValid("{[]})"), false)
	assert.Equal(t, isValid("{[{]}"), false)
}

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(1)
	s.Push(2)
	s.Push(-10)
	s.Push(2)
	assert.Equal(t, s.Top(), 2)
	assert.Equal(t, s.GetMin(), -10)
	s.Pop()
	assert.Equal(t, s.Top(), -10)
	assert.Equal(t, s.GetMin(), -10)
	s.Pop()
	assert.Equal(t, s.Top(), 2)
	assert.Equal(t, s.GetMin(), 1)
}

func TestEvalRPN(t *testing.T) {
	assert.Equal(t, evalRPN([]string{"2", "1", "+", "3", "*"}), 9)
	assert.Equal(t, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}), 22)
}
func TestCircularDeque(t *testing.T) {
	dq := NewMyCircularDeque(3)
	assert.Equal(t, dq.InsertLast(1), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertLast(2), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertFront(3), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertFront(4), false)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.GetRear(), 2)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.IsFull(), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.DeleteLast(), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.InsertFront(4), true)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
	assert.Equal(t, dq.GetFront(), 4)
	t.Log(dq.data, dq.GetFront(), dq.GetRear())
}
