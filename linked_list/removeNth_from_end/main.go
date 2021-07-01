package main

type ListNode struct {
    Val int
	Next *ListNode
}

func main(){

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

    if head.Next == nil{
        return nil
    }

	quick,slow := &ListNode{},&ListNode{}
	quick.Next = head
	slow.Next = head

	for i:=0;i<n;i++{
		quick = quick.Next
	}

	for quick.Next != nil{
		quick = quick.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	if slow.Next == head.Next && slow != head{
		return slow.Next
	}else{
        return head
    }

	 
}