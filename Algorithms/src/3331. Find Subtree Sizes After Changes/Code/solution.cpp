class Solution {
public:
    vector<int> findSubtreeSizes(vector<int>& parent, string s) {
        int n = parent.size();
        vector<vector<int>> g(n);
        for (int i = 1; i < n; i++) g[parent[i]].push_back(i);
        vector<int> np(parent.begin(), parent.end());
        vector<int> last(26, -1);
        function<void(int)> dfs = [&](int u) {
            int c = s[u] - 'a';
            int prev = last[c];
            if (prev != -1) np[u] = prev;
            last[c] = u;
            for (int v : g[u]) dfs(v);
            last[c] = prev;
        };
        dfs(0);
        vector<vector<int>> ng(n);
        for (int i = 1; i < n; i++) ng[np[i]].push_back(i);
        vector<int> ans(n);
        function<int(int)> size = [&](int u) {
            int sz = 1;
            for (int v : ng[u]) sz += size(v);
            return ans[u] = sz;
        };
        size(0);
        return ans;
    }
};