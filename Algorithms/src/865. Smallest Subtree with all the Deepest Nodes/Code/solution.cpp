class Solution {
public:
    // Helper function returns:
    // {maximum depth from this node, subtree root containing all deepest nodes}
    pair<int, TreeNode*> dfs(TreeNode* root) {
        if (!root) return {0, nullptr};

        auto left = dfs(root->left);
        auto right = dfs(root->right);

        if (left.first == right.first) {
            // Both sides have deepest nodes at same depth
            return {left.first + 1, root};
        } else if (left.first > right.first) {
            return {left.first + 1, left.second};
        } else {
            return {right.first + 1, right.second};
        }
    }

    TreeNode* subtreeWithAllDeepest(TreeNode* root) {
        return dfs(root).second;
    }
};
