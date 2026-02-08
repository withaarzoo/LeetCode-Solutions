class Solution {

    // Returns height if balanced, else -1
    private int height(TreeNode root) {
        if (root == null)
            return 0;

        int leftHeight = height(root.left);
        if (leftHeight == -1)
            return -1;

        int rightHeight = height(root.right);
        if (rightHeight == -1)
            return -1;

        if (Math.abs(leftHeight - rightHeight) > 1)
            return -1;

        return 1 + Math.max(leftHeight, rightHeight);
    }

    public boolean isBalanced(TreeNode root) {
        return height(root) != -1;
    }
}
