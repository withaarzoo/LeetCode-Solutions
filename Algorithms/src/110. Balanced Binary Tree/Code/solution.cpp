class Solution
{
public:
    // Helper function that returns height of the tree
    // If tree is unbalanced, return -1
    int height(TreeNode *root)
    {
        if (root == nullptr)
            return 0;

        int leftHeight = height(root->left);
        if (leftHeight == -1)
            return -1;

        int rightHeight = height(root->right);
        if (rightHeight == -1)
            return -1;

        if (abs(leftHeight - rightHeight) > 1)
            return -1;

        return 1 + max(leftHeight, rightHeight);
    }

    bool isBalanced(TreeNode *root)
    {
        return height(root) != -1;
    }
};
