package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1 *list.List
	stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

//put the element at the tail
func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

//return head
func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() > 0 {
		// fmt.Println("stack2 is not null")

		elem := this.stack2.Back()
		this.stack2.Remove(elem)
		return elem.Value.(int)
	}

	for this.stack1.Len() > 0 {
		elem := this.stack1.Back()
		this.stack2.PushBack(elem.Value.(int))
		this.stack1.Remove(elem)

	}
	if this.stack2.Len() > 0 {
		outElem := this.stack2.Back()
		this.stack2.Remove(outElem)
		// out = outElem.Value.(int)
		return outElem.Value.(int)
	}
	return -1
}

func main() {
	obj := Constructor()
	obj.AppendTail(1)
	fmt.Printf("poping %d\n", obj.DeleteHead())
	fmt.Println(obj)
}
