class Solution
{
public:
    int findMaxPathScore(vector<vector<int>> &edges, vector<bool> &online, long long k)
    {
        int n = online.size();

        // Build graph and indegree for topological sorting
        vector<vector<pair<int, int>>> graph(n);
        vector<int> indegree(n, 0);

        for (auto &e : edges)
        {
            graph[e[0]].push_back({e[1], e[2]});
            indegree[e[1]]++;
        }

        // Compute topological order once because the graph never changes
        queue<int> q;
        for (int i = 0; i < n; i++)
            if (indegree[i] == 0)
                q.push(i);

        vector<int> topo;
        while (!q.empty())
        {
            int u = q.front();
            q.pop();
            topo.push_back(u);

            for (auto &[v, w] : graph[u])
            {
                if (--indegree[v] == 0)
                    q.push(v);
            }
        }

        // Check whether a minimum edge value "limit" is possible
        auto check = [&](int limit)
        {
            const long long INF = (1LL << 60);

            // dp[i] = minimum total cost to reach node i
            vector<long long> dp(n, INF);
            dp[0] = 0;

            for (int u : topo)
            {

                // Skip unreachable nodes
                if (dp[u] == INF)
                    continue;

                // Offline intermediate nodes cannot be used
                if (u != 0 && u != n - 1 && !online[u])
                    continue;

                for (auto &[v, w] : graph[u])
                {

                    // Edge is too small
                    if (w < limit)
                        continue;

                    // Cannot move into an offline intermediate node
                    if (v != n - 1 && !online[v])
                        continue;

                    if (dp[u] + w < dp[v])
                        dp[v] = dp[u] + w;
                }
            }

            return dp[n - 1] <= k;
        };

        int left = 0, right = 1000000000;
        int ans = -1;

        // Binary search on the answer
        while (left <= right)
        {
            int mid = left + (right - left) / 2;

            if (check(mid))
            {
                ans = mid;
                left = mid + 1;
            }
            else
            {
                right = mid - 1;
            }
        }

        return ans;
    }
};