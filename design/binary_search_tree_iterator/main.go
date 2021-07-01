package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type BSTIterator struct {
	vals []int
	current int
}

//helper function for constructor
//inorder traversal
func dfs(root *TreeNode,iterator *BSTIterator){
	if root == nil{
		return
	}

	dfs(root.Left,iterator)
	iterator.vals = append(iterator.vals,root.Val)
	dfs(root.Right,iterator)
}

func Constructor(root *TreeNode) BSTIterator {

	bst_iter := BSTIterator{current:0,}

	dfs(root,&bst_iter)
	return bst_iter
}

func (this *BSTIterator) Next() int {
	if len(this.vals) == 0{
		return 0
	}
	ans := this.vals[this.current]
	this.current++
	return ans
}


func (this *BSTIterator) HasNext() bool {
	
	if this.current<len(this.vals){
		return true
	}
	return false
}


func main(){

}