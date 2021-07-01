# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# recursive solution
class Solution:

    def inorderTraversal(self, root: TreeNode):
        
        if not root:
            return []
        
        res  =[]
        def inOrder(root):
            if root is None:
                return 
            
            inOrder(root.left)
            res.append(root.val)
            inOrder(root.right)
            
        inOrder(root)
    
        return res

# iterative solution
class Solution_iterative:
    pass