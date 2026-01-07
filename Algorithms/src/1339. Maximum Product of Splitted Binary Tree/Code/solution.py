class Solution:
    def maxProduct(self, root: Optional[TreeNode]) -> int:
        MOD = 10**9 + 7
        self.max_prod = 0

        def get_total_sum(node):
            if not node:
                return 0
            return node.val + get_total_sum(node.left) + get_total_sum(node.right)

        def dfs(node):
            if not node:
                return 0

            left = dfs(node.left)
            right = dfs(node.right)

            subtree_sum = node.val + left + right
            self.max_prod = max(self.max_prod, subtree_sum * (total_sum - subtree_sum))

            return subtree_sum

        total_sum = get_total_sum(root)
        dfs(root)
        return self.max_prod % MOD
