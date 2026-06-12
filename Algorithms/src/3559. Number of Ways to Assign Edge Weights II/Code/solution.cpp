class Solution
{
public:
    static const int MOD = 1000000007;

    vector<int> depth;
    vector<vector<int>> up;
    vector<vector<int>> graph;
    int LOG;

    // DFS to fill depth and binary lifting table
    void dfs(int node, int parent)
    {
        up[node][0] = parent;

        for (int j = 1; j < LOG; j++)
        {
            up[node][j] = up[up[node][j - 1]][j - 1];
        }

        for (int next : graph[node])
        {
            if (next == parent)
                continue;

            depth[next] = depth[node] + 1;
            dfs(next, node);
        }
    }

    // Find LCA using binary lifting
    int lca(int a, int b)
    {
        if (depth[a] < depth[b])
            swap(a, b);

        int diff = depth[a] - depth[b];

        for (int j = LOG - 1; j >= 0; j--)
        {
            if ((diff >> j) & 1)
            {
                a = up[a][j];
            }
        }

        if (a == b)
            return a;

        for (int j = LOG - 1; j >= 0; j--)
        {
            if (up[a][j] != up[b][j])
            {
                a = up[a][j];
                b = up[b][j];
            }
        }

        return up[a][0];
    }

    vector<int> assignEdgeWeights(vector<vector<int>> &edges, vector<vector<int>> &queries)
    {
        int n = edges.size() + 1;

        LOG = 17;
        while ((1 << LOG) <= n)
            LOG++;

        graph.assign(n + 1, {});
        for (auto &e : edges)
        {
            int u = e[0];
            int v = e[1];

            graph[u].push_back(v);
            graph[v].push_back(u);
        }

        depth.assign(n + 1, 0);
        up.assign(n + 1, vector<int>(LOG, 1));

        dfs(1, 1);

        // Precompute powers of 2 modulo MOD
        vector<int> pow2(n + 1, 1);
        for (int i = 1; i <= n; i++)
        {
            pow2[i] = (long long)pow2[i - 1] * 2 % MOD;
        }

        vector<int> ans;

        for (auto &q : queries)
        {
            int u = q[0];
            int v = q[1];

            int ancestor = lca(u, v);

            int dist = depth[u] + depth[v] - 2 * depth[ancestor];

            if (dist == 0)
            {
                ans.push_back(0);
            }
            else
            {
                ans.push_back(pow2[dist - 1]);
            }
        }

        return ans;
    }
};