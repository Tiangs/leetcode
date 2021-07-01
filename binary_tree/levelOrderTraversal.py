# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def levelOrderBottom(self, root: TreeNode): 
        if not root:
            return []
        q = [root]
        res = []

        while q:
            res_lvl = []
            for  _ in range(len(q)):

                # pop the first element
                node = q.pop(0)

                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)

                res_lvl.append(node.val)
            
            res.append(res_lvl)

        # reverse the final result
        res_reversed = []
        while res:
            res_reversed.append(res.pop())
        return res_reversed




        