package main

 type ListNode struct {
    Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
    dummy := head
 
	for head.Next != nil{

		if head.Next.Val == head.Val{
			head.Next = head.Next.Next
		}else{
			head = head.Next
		}

	}
	return dummy

}

func main(){

}