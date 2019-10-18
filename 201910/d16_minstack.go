package _01910

import "errors"

var errEmptyStack error = errors.New("min stack is empty")

type MinStack struct {
	orderList       []int
	realJiaLocation map[int]int
	err             error
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{orderList: make([]int, 0), realJiaLocation: make(map[int]int, 0)}
}

func (this *MinStack) Push(x int) {
	currentL := locationMyself(this.orderList, x)
	newList := make([]int, 0)
	newList = append(newList, this.orderList[:currentL]...)
	newList = append(newList, x)
	newList = append(newList, this.orderList[currentL:]...)
	this.orderList = newList
	this.realJiaLocation[len(this.orderList)] = currentL
}

//divide and find--find index
func locationMyself(orderDescList []int, value int) int {
	for i, v := range orderDescList {
		if value > v {
			return i
		}
	}
	return len(orderDescList)
}

func (this *MinStack) Pop() {
	realLocation := this.realJiaLocation[len(this.orderList)]
	this.orderList = append(this.orderList[:realLocation], this.orderList[realLocation+1:]...)
	delete(this.realJiaLocation, len(this.orderList))
}

func (this *MinStack) Top() int {
	if len(this.orderList) >= 0 {
		realLocation := this.realJiaLocation[len(this.orderList)]
		return this.orderList[realLocation]
	}
	this.err = errEmptyStack
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.orderList) > 0 {
		return this.orderList[len(this.orderList)-1]
	}
	this.err = errEmptyStack
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
