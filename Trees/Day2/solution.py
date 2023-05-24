# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    maxx=float(-inf)
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        def count(root, c):
            if root ==None:
                self.maxx=max(self.maxx, c)
                return
            count(root.left, c+1)
            count(root.right, c+1)
        count(root, 0)
        return self.maxx
