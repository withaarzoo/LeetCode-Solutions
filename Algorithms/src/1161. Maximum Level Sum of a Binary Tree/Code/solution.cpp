class Solution {
public:
    int maxLevelSum(TreeNode* root) {
        queue<TreeNode*> q;
        q.push(root);

        int level = 1;
        int answerLevel = 1;
        long long maxSum = LLONG_MIN;

        // BFS traversal
        while (!q.empty()) {
            int size = q.size();
            long long levelSum = 0;

            // Process all nodes of current level
            for (int i = 0; i < size; i++) {
                TreeNode* node = q.front();
                q.pop();

                levelSum += node->val;

                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
            }

            // Update maximum sum and level
            if (levelSum > maxSum) {
                maxSum = levelSum;
                answerLevel = level;
            }

            level++;
        }

        return answerLevel;
    }
};
