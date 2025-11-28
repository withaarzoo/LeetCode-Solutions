class Solution
{
public:
    int maxKDivisibleComponents(int n, vector<vector<int>> &edges, vector<int> &values, int k)
    {
        // Build adjacency list for the tree
        vector<vector<int>> adj(n);
        for (auto &e : edges)
        {
            int u = e[0], v = e[1];
            adj[u].push_back(v);
            adj[v].push_back(u);
        }

        long long ans = 0; // will store the number of components

        // DFS function: returns subtree sum % k for node 'u'
        function<long long(int, int)> dfs = [&](int u, int parent) -> long long
        {
            long long sum = values[u] % k; // start with current node's value modulo k

            for (int v : adj[u])
            {
                if (v == parent)
                    continue;                   // don't go back to parent
                long long childRem = dfs(v, u); // remainder from child subtree
                sum = (sum + childRem) % k;     // accumulate modulo k
            }

            // If subtree sum is divisible by k, it can form its own component
            if (sum % k == 0)
            {
                ans++;
                return 0; // nothing is passed up because this subtree is separated
            }
            return sum; // pass remainder up to parent
        };

        dfs(0, -1); // root the tree at node 0

        return (int)ans;
    }
};
