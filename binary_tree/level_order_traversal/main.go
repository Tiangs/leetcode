package main
import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	queue := []*TreeNode{root}
	res := [][]int{}
	
	for queue!=nil{
		n := len(queue)
		t := []int{}
		for i:=0;i<n;i++{
			//pop out the last elem
			elem := queue[0]

			t = append(t,elem.Val)

			queue = queue[1:]
			
			queue = append(queue,elem.Left)
			queue = append(queue,elem.Right)

		}

		res = append(res,t)

	}
	return res 
}

func main(){

	// fmt.Println(res)
	// res  = append(res,[]int{1,2,3})
	// res  = append(res,[]int{1,2,3})

	m,n :=4,5
	res := make([][]int,m)
	for i:=range res{
		res[i] = make([]int,n)
	}
	fmt.Println(res)
}