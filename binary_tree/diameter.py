# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def __init__(self):
        self.max_diameter = 0

    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        self.dfs(root)
        return self.max_diameter

    def dfs(self,root):

        if root is None:
            return 0
            
        L = self.dfs(root.left)
        R = self.dfs(root.right)

        if L + R > self.max_diameter:
            self.max_diameter = L + R

        return max(L,R) + 1
        
