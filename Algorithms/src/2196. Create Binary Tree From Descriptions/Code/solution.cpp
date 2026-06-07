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
public:
    TreeNode *createBinaryTree(vector<vector<int>> &descriptions)
    {

        // Stores value -> TreeNode mapping
        unordered_map<int, TreeNode *> nodes;

        // Stores all nodes that appear as children
        unordered_set<int> children;

        for (auto &d : descriptions)
        {
            int parent = d[0];
            int child = d[1];
            int isLeft = d[2];

            // Create parent node if not present
            if (!nodes.count(parent))
                nodes[parent] = new TreeNode(parent);

            // Create child node if not present
            if (!nodes.count(child))
                nodes[child] = new TreeNode(child);

            // Connect child to correct side
            if (isLeft)
                nodes[parent]->left = nodes[child];
            else
                nodes[parent]->right = nodes[child];

            // Mark child node
            children.insert(child);
        }

        // Root is the node that never appeared as a child
        for (auto &[value, node] : nodes)
        {
            if (!children.count(value))
                return node;
        }

        return nullptr;
    }
};