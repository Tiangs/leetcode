package main


func main(){

}

 type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode

 }
 func postorderTraversal(root *TreeNode) []int {

    var stack []*TreeNode 
    var res []int 
    var prev *TreeNode
    for root != nil || len(stack) > 0{

        //if current node has Left, push it into stack
        for root != nil{
            stack = append(stack,root)
            root = root.Left
        }
        
        //pop the last elem from stack
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if root.Right == nil || root.Right == prev{
            res = append(res,root.Val)
            prev = root
            root = nil
        }else{
            stack = append(stack,root)
            root = root.Right
        }

    }
    return res

}