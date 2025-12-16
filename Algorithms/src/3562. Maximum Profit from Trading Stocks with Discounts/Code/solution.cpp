class Solution
{
public:
    int B;
    vector<vector<int>> tree;
    vector<int> present, future;
    vector<vector<vector<int>>> dp;

    vector<int> merge(const vector<int> &A, const vector<int> &B2)
    {
        vector<int> C(B + 1, -1e9);
        for (int i = 0; i <= B; i++)
        {
            if (A[i] < 0)
                continue;
            for (int j = 0; i + j <= B; j++)
            {
                C[i + j] = max(C[i + j], A[i] + B2[j]);
            }
        }
        return C;
    }

    void dfs(int u)
    {
        dp[u][0] = vector<int>(B + 1, 0);
        dp[u][1] = vector<int>(B + 1, 0);

        for (int v : tree[u])
        {
            dfs(v);
        }

        for (int parentBought = 0; parentBought <= 1; parentBought++)
        {
            int price = parentBought ? present[u] / 2 : present[u];
            int profit = future[u] - price;

            vector<int> skip(B + 1, 0);
            for (int v : tree[u])
            {
                skip = merge(skip, dp[v][0]);
            }

            vector<int> take(B + 1, -1e9);
            if (price <= B)
            {
                vector<int> base(B + 1, 0);
                for (int v : tree[u])
                {
                    base = merge(base, dp[v][1]);
                }
                for (int b = price; b <= B; b++)
                {
                    take[b] = base[b - price] + profit;
                }
            }

            for (int b = 0; b <= B; b++)
            {
                dp[u][parentBought][b] = max(skip[b], take[b]);
            }
        }
    }

    int maxProfit(int n, vector<int> &p, vector<int> &f,
                  vector<vector<int>> &hierarchy, int budget)
    {
        B = budget;
        present = p;
        future = f;

        tree.assign(n, {});
        for (auto &e : hierarchy)
            tree[e[0] - 1].push_back(e[1] - 1);

        dp.assign(n, vector<vector<int>>(2));
        dfs(0);

        int ans = 0;
        for (int b = 0; b <= B; b++)
            ans = max(ans, dp[0][0][b]);
        return ans;
    }
};
