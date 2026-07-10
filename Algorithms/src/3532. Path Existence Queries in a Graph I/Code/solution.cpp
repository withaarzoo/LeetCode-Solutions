class Solution {
public:
    vector<bool> pathExistenceQueries(int n, vector<int>& nums, int maxDiff, vector<vector<int>>& queries) {
        // component[i] stores which connected component node i belongs to.
        vector<int> component(n, 0);

        // Start with component 0 for the first node.
        int componentId = 0;

        // Check every gap between two consecutive sorted values.
        for (int i = 1; i < n; i++) {
            // A gap larger than maxDiff separates the graph into two parts.
            if (nums[i] - nums[i - 1] > maxDiff) {
                componentId++;
            }

            // Store the component of the current node.
            component[i] = componentId;
        }

        // Store the result of every query.
        vector<bool> answer;
        answer.reserve(queries.size());

        // Two nodes have a path exactly when their component IDs are equal.
        for (const auto& query : queries) {
            int u = query[0];
            int v = query[1];

            answer.push_back(component[u] == component[v]);
        }

        return answer;
    }
}; 