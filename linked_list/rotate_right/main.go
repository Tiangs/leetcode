package main

type ListNode struct {
    Val int
    Next *ListNode
}

func main(){

}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil{
		return nil
	}

	//make the single linked list a cycle
	count := 1
	cur := head

	for cur.Next != nil{
		cur = cur.Next
		count++
	}

	if k%count==0{
		return head
	}

	cur.Next = head

	//still have steps to go
	steps := count - k%count

	for i:=0;i<steps;i++{
		cur = cur.Next
	}

	next:=cur.Next
	cur.Next = nil
	return next
}