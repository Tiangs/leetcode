package main

type ListNode struct {
    Val int
    Next *ListNode
}

func reverse(head *ListNode)*ListNode{
	var prev *ListNode
	cur:= head
	for cur != nil{
		next:=cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// define a dummy head to be returned
	dummy:= &ListNode{}
	dummy.Next = head

	//pre move to left-1
	pre:= dummy
	for i:= 0;i<left-1;i++{
		pre = pre.Next
	}

	//right node to right
	rightNode := pre
	for i:= 0;i<right-left+1;i++{
		rightNode = rightNode.Next
	}
	
	//left node to left
	leftNode := pre.Next

	//curr to node right+1
	curr := rightNode.Next

	//cut off the sub linkedlist
	pre.Next = nil
	rightNode.Next = nil

	//reverse from left node to right node
	reverse(leftNode)
	
	//link the  sublist
	pre.Next = rightNode
	leftNode.Next = curr

	return dummy.Next
}

func main(){

}