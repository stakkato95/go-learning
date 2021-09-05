package stack

import "errors"

type IntStack struct {
	nextIndex int
	capacity  int
	data      []int
}

func NewStack(cap int) IntStack {
	return IntStack{capacity: cap, data: make([]int, cap)}
}

func (is *IntStack) Push(n int) error {
	if is.capacity == is.nextIndex {
		return errors.New("stack is full")
	}

	is.data[is.nextIndex] = n
	is.nextIndex++
	return nil
}

func (is *IntStack) Pop() (int, bool) {
	if is.nextIndex == 0 {
		return 0, false
	}
	is.nextIndex--
	return is.data[is.nextIndex], true
}

func (is *IntStack) Peek() (int, bool) {
	if is.nextIndex == 0 {
		return 0, false
	}
	return is.data[is.nextIndex-1], true
}

func (is IntStack) Size() int {
	return is.nextIndex
}
