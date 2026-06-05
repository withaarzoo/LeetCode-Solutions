class Solution
{
public:
    struct Node
    {
        long long cnt;
        long long wav;
    };

    string s;
    Node dp[20][2][11][11];
    bool vis[20][2][11][11];

    Node dfs(int pos, int started, int last, int secondLast, bool tight)
    {
        if (pos == (int)s.size())
        {
            return {1, 0};
        }

        if (!tight && vis[pos][started][last][secondLast])
        {
            return dp[pos][started][last][secondLast];
        }

        int limit = tight ? s[pos] - '0' : 9;

        Node res = {0, 0};

        for (int d = 0; d <= limit; d++)
        {
            bool ntight = tight && (d == limit);

            if (!started && d == 0)
            {
                Node nxt = dfs(pos + 1, 0, 10, 10, ntight);

                res.cnt += nxt.cnt;
                res.wav += nxt.wav;
            }
            else
            {
                long long add = 0;

                if (started && secondLast != 10)
                {
                    if ((last > secondLast && last > d) ||
                        (last < secondLast && last < d))
                    {
                        add = 1;
                    }
                }

                int nSecondLast = started ? last : 10;
                int nLast = d;

                Node nxt = dfs(pos + 1, 1, nLast, nSecondLast, ntight);

                res.cnt += nxt.cnt;
                res.wav += nxt.wav + add * nxt.cnt;
            }
        }

        if (!tight)
        {
            vis[pos][started][last][secondLast] = true;
            dp[pos][started][last][secondLast] = res;
        }

        return res;
    }

    long long solve(long long n)
    {
        if (n < 0)
            return 0;

        s = to_string(n);
        memset(vis, 0, sizeof(vis));

        return dfs(0, 0, 10, 10, true).wav;
    }

    long long totalWaviness(long long num1, long long num2)
    {
        return solve(num2) - solve(num1 - 1);
    }
};