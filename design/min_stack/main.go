package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	topMin := this.minStack[len(this.minStack)-1]
	if x < topMin {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, topMin)
	}
}

func (this *MinStack) Pop() {

	if this.stack != nil {
		// elem := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
		return
	}

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func main() {

}
