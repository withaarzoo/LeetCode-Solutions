# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def averageOfSubtree(self, root: TreeNode) -> int:
        answer = 0  # Stores the number of nodes equal to their subtree average.

        def dfs(node):
            nonlocal answer  # Allows the recursive function to update the outer answer variable.

            if node is None:
                return 0, 0  # An empty subtree has sum 0 and contains 0 nodes.

            left_sum, left_count = dfs(node.left)  # Get the sum and count of the left subtree.
            right_sum, right_count = dfs(node.right)  # Get the sum and count of the right subtree.

            total_sum = left_sum + right_sum + node.val  # Calculate the sum of the current subtree.
            total_count = left_count + right_count + 1  # Calculate the number of nodes in the current subtree.

            if node.val == total_sum // total_count:
                answer += 1  # Count the node when its value equals the floored average.

            return total_sum, total_count  # Return this subtree's sum and count to its parent.

        dfs(root)  # Traverse the tree and process every node.
        return answer  # Return the total number of matching nodes.