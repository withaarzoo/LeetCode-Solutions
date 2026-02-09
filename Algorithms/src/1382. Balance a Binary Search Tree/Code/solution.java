class Solution {
    List<Integer> arr = new ArrayList<>();

    // Inorder traversal
    void inorder(TreeNode root) {
        if (root == null)
            return;
        inorder(root.left);
        arr.add(root.val);
        inorder(root.right);
    }

    // Build balanced BST
    TreeNode build(int left, int right) {
        if (left > right)
            return null;

        int mid = left + (right - left) / 2;
        TreeNode node = new TreeNode(arr.get(mid));

        node.left = build(left, mid - 1);
        node.right = build(mid + 1, right);

        return node;
    }

    public TreeNode balanceBST(TreeNode root) {
        inorder(root);
        return build(0, arr.size() - 1);
    }
}
