class Solution
{
public:
    static const int MOD = 1000000007;

    // Fast modular exponentiation
    long long modPow(long long base, long long exp)
    {
        long long result = 1;

        while (exp > 0)
        {
            // If current bit is set, multiply answer
            if (exp & 1)
            {
                result = (result * base) % MOD;
            }

            // Square the base for next bit
            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    int assignEdgeWeights(vector<vector<int>> &edges)
    {
        int n = edges.size() + 1;

        // Build adjacency list
        vector<vector<int>> graph(n + 1);

        for (auto &e : edges)
        {
            int u = e[0];
            int v = e[1];

            graph[u].push_back(v);
            graph[v].push_back(u);
        }

        int maxDepth = 0;

        // Iterative DFS: {node, depth}
        stack<pair<int, int>> st;
        st.push({1, 0});

        vector<int> visited(n + 1, 0);
        visited[1] = 1;

        while (!st.empty())
        {
            auto [node, depth] = st.top();
            st.pop();

            maxDepth = max(maxDepth, depth);

            for (int nei : graph[node])
            {
                if (!visited[nei])
                {
                    visited[nei] = 1;
                    st.push({nei, depth + 1});
                }
            }
        }

        // Number of odd-parity assignments = 2^(maxDepth - 1)
        return (int)modPow(2, maxDepth - 1);
    }
};