class Solution
{
public:
    vector<int> arr;

    // Step 1: Inorder traversal to get sorted values
    void inorder(TreeNode *root)
    {
        if (!root)
            return;
        inorder(root->left);
        arr.push_back(root->val);
        inorder(root->right);
    }

    // Step 2: Build balanced BST from sorted array
    TreeNode *build(int left, int right)
    {
        if (left > right)
            return nullptr;

        int mid = left + (right - left) / 2;
        TreeNode *node = new TreeNode(arr[mid]);

        node->left = build(left, mid - 1);
        node->right = build(mid + 1, right);

        return node;
    }

    TreeNode *balanceBST(TreeNode *root)
    {
        inorder(root);
        return build(0, arr.size() - 1);
    }
};
