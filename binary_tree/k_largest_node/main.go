package main

type TreeNode struct {
   Val int
   Left *TreeNode
   Right *TreeNode
}

type Res struct{
	arr []int
}
func kthLargest(root *TreeNode, k int) int {
	r := &Res{arr:[]int{}}

	dfs(root,r)
	// fmt.Println(r.arr)
	return r.arr[len(r.arr)-k]
}

func dfs(root *TreeNode,r *Res){
	if root == nil{
		return 
	}
	dfs(root.Left,r)
	r.arr = append(r.arr,root.Val)
	dfs(root.Right,r)

}

func main(){

}