package _01910

import "testing"

func TestMinStack_GetMin(t *testing.T) {
	minStack:=Constructor()
	minStack.Push(-2)
	minStack.Push(1)
	minStack.Push(-3)
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())


}
