class Solution {
    long totalSum = 0;
    long maxProduct = 0;
    int MOD = 1000000007;

    private long getTotalSum(TreeNode root) {
        if (root == null) return 0;
        return root.val + getTotalSum(root.left) + getTotalSum(root.right);
    }

    private long dfs(TreeNode root) {
        if (root == null) return 0;

        long left = dfs(root.left);
        long right = dfs(root.right);

        long subtreeSum = root.val + left + right;
        maxProduct = Math.max(maxProduct, subtreeSum * (totalSum - subtreeSum));

        return subtreeSum;
    }

    public int maxProduct(TreeNode root) {
        totalSum = getTotalSum(root);
        dfs(root);
        return (int)(maxProduct % MOD);
    }
}
