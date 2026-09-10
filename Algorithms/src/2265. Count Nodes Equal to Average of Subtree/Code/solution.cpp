/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution
{
    int answer = 0; // Stores how many nodes have a value equal to their subtree average.

    pair<int, int> dfs(TreeNode *node)
    {
        if (node == nullptr)
        {
            return {0, 0}; // An empty subtree has sum 0 and contains 0 nodes.
        }

        auto [leftSum, leftCount] = dfs(node->left);    // Get sum and count from the left subtree.
        auto [rightSum, rightCount] = dfs(node->right); // Get sum and count from the right subtree.

        int sum = leftSum + rightSum + node->val; // Add both child sums and the current node value.
        int count = leftCount + rightCount + 1;   // Add both child counts and the current node.

        if (node->val == sum / count)
        {
            answer++; // Count this node when its value equals the floored subtree average.
        }

        return {sum, count}; // Return this subtree's sum and node count to its parent.
    }

public:
    int averageOfSubtree(TreeNode *root)
    {
        dfs(root);     // Process every node using postorder traversal.
        return answer; // Return the total number of matching nodes.
    }
};