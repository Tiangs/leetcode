package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	cur := &ListNode{0, nil}

	//dummy head , no ops, to be returned
	head := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	//l1 or l2 still contain values
	if l1 != nil {
		cur.Next = l1
	}

	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}
