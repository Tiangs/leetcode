# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:

    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        if t1 is None:
            return t2

        if t2 is None:
            return t1

        node = TreeNode()
        node.val = t1.val + t2.val
        node.left = mergeTrees(self,t1.left,t2.left)
        node.right= mergeTrees(self,t1.right,t2.right)

        return node
