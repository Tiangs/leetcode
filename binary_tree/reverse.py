# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def invertTree(self, root: TreeNode) -> TreeNode:
        def dfs(node):
            if node is None: 
                return

            dfs(node.left)
            dfs(node.right)

            node.left, node.right = node.right,node.left

            return node 

        return dfs(root)

        

        