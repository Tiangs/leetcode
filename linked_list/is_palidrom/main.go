package main
 
type ListNode struct {
    Val int
    Next *ListNode
}
 
 
func main(){

}

func isPalindrome(head *ListNode) bool {
	vals:= []int{}
	for head != nil{
		vals = append(vals,head.Val)
		head = head.Next
	}

	i,j := 0,len(vals)-1
	for i<=j{
		if vals[i]!=vals[j]{
			return false
		}
		i++
		j--
	}

	return true

}