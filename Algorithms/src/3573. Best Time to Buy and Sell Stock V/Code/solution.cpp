class Solution
{
public:
    int n, K;
    vector<int> p;
    long long dp[1005][505][3];

    long long solve(int i, int t, int s)
    {
        if (i == n)
        {
            return (s == 0 ? 0 : LLONG_MIN / 4);
        }

        long long &res = dp[i][t][s];
        if (res != -1)
            return res;

        res = solve(i + 1, t, s); // skip

        if (s == 0)
        {
            res = max(res, solve(i + 1, t, 1) - p[i]); // buy
            res = max(res, solve(i + 1, t, 2) + p[i]); // short sell
        }
        else if (s == 1 && t < K)
        {
            res = max(res, solve(i + 1, t + 1, 0) + p[i]); // sell
        }
        else if (s == 2 && t < K)
        {
            res = max(res, solve(i + 1, t + 1, 0) - p[i]); // buy back
        }

        return res;
    }

    long long maximumProfit(vector<int> &prices, int k)
    {
        p = prices;
        n = prices.size();
        K = k;
        memset(dp, -1, sizeof(dp));
        return solve(0, 0, 0);
    }
};
