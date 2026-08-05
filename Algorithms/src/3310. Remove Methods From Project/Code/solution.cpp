class Solution {
public:
    vector<int> remainingMethods(int n, int k, vector<vector<int>>& invocations) {
        // Build the directed graph
        vector<vector<int>> graph(n);
        for (auto &edge : invocations) {
            graph[edge[0]].push_back(edge[1]);
        }

        // Marks whether a method is suspicious
        vector<bool> vis(n, false);

        // DFS to mark every method reachable from k
        function<void(int)> dfs = [&](int u) {
            vis[u] = true;

            // Visit every invoked method
            for (int v : graph[u]) {
                if (!vis[v]) {
                    dfs(v);
                }
            }
        };

        dfs(k);

        // Check whether any non-suspicious method calls a suspicious one
        for (auto &edge : invocations) {
            int u = edge[0];
            int v = edge[1];

            if (!vis[u] && vis[v]) {
                // Removal is impossible, return all methods
                vector<int> ans;
                for (int i = 0; i < n; i++) {
                    ans.push_back(i);
                }
                return ans;
            }
        }

        // Keep only non-suspicious methods
        vector<int> ans;
        for (int i = 0; i < n; i++) {
            if (!vis[i]) {
                ans.push_back(i);
            }
        }

        return ans;
    }
};