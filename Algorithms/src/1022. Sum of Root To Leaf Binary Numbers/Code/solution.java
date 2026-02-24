class Solution {
    
    private int dfs(TreeNode node, int current) {
        if (node == null) return 0;
        
        // Build binary number
        current = current * 2 + node.val;
        
        // If leaf node
        if (node.left == null && node.right == null) {
            return current;
        }
        
        // Recurse on left and right
        return dfs(node.left, current) + dfs(node.right, current);
    }
    
    public int sumRootToLeaf(TreeNode root) {
        return dfs(root, 0);
    }
}