package main

func main(){

}
//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func kthToLast(head *ListNode, k int) int {

	if head == nil{
		return -1 
	}
	quick,slow:= &ListNode{},&ListNode{}

	quick.Next = head
	slow.Next = head

	for i:=0;i<k;i++{
		quick = quick.Next
	}

	for quick.Next != nil{
		quick = quick.Next
		slow = slow.Next
	}

	return slow.Next.Val
}