class Solution
{
public:
    static const unsigned long long INF = ULLONG_MAX;
    unsigned long long dist[201][201];
    unsigned long long dp[1001];

    unordered_map<string, int> id;
    unordered_set<int> lens;

    long long minimumCost(string source, string target,
                          vector<string> &original,
                          vector<string> &changed,
                          vector<int> &cost)
    {

        int m = original.size();
        int n = source.size();

        id.clear();
        lens.clear();

        int sz = 0;
        memset(dist, 255, sizeof(dist));

        // Build graph
        for (int i = 0; i < m; i++)
        {
            if (!id.count(original[i]))
            {
                id[original[i]] = sz++;
                lens.insert(original[i].size());
            }
            if (!id.count(changed[i]))
            {
                id[changed[i]] = sz++;
            }
            int u = id[original[i]];
            int v = id[changed[i]];
            dist[u][v] = min(dist[u][v], (unsigned long long)cost[i]);
        }

        for (int i = 0; i < sz; i++)
            dist[i][i] = 0;

        // Floyd Warshall
        for (int k = 0; k < sz; k++)
            for (int i = 0; i < sz; i++)
                if (dist[i][k] != INF)
                    for (int j = 0; j < sz; j++)
                        if (dist[k][j] != INF)
                            dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]);

        memset(dp, 255, sizeof(dp));
        dp[0] = 0;

        // DP
        for (int i = 0; i < n; i++)
        {
            if (dp[i] == INF)
                continue;

            if (source[i] == target[i])
                dp[i + 1] = min(dp[i + 1], dp[i]);

            for (int L : lens)
            {
                if (i + L > n)
                    continue;

                string s = source.substr(i, L);
                string t = target.substr(i, L);

                if (id.count(s) && id.count(t))
                {
                    unsigned long long d = dist[id[s]][id[t]];
                    if (d != INF)
                        dp[i + L] = min(dp[i + L], dp[i] + d);
                }
            }
        }

        return dp[n] == INF ? -1 : dp[n];
    }
};
