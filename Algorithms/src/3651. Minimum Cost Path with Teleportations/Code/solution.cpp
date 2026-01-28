#include <bits/stdc++.h>
using namespace std;
using ll = long long;
const ll INF = (ll)4e18;

class Solution
{
public:
    int minCost(vector<vector<int>> &grid, int k)
    {
        int m = grid.size(), n = grid[0].size();
        // dp: best cost to reach each cell with up to current teleports
        vector<vector<ll>> dp(m, vector<ll>(n, INF));
        // base dp (no teleport): simple right/down DP starting from (0,0) cost 0
        dp[0][0] = 0;
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (i > 0)
                    dp[i][j] = min(dp[i][j], dp[i - 1][j] + grid[i][j]);
                if (j > 0)
                    dp[i][j] = min(dp[i][j], dp[i][j - 1] + grid[i][j]);
            }
        }

        // Prepare list of cells (value, i, j)
        vector<tuple<int, int, int>> cells;
        cells.reserve(m * n);
        for (int i = 0; i < m; ++i)
            for (int j = 0; j < n; ++j)
                cells.emplace_back(grid[i][j], i, j);
        // sort descending by value for grouping
        sort(cells.begin(), cells.end(), [](auto &a, auto &b)
             { return get<0>(a) > get<0>(b); });

        for (int step = 0; step < k; ++step)
        {
            // compute start_cost after using one teleport (or not) from dp
            vector<vector<ll>> start(m, vector<ll>(n, INF));
            ll running_min = INF;
            int idx = 0;
            while (idx < (int)cells.size())
            {
                int val = get<0>(cells[idx]);
                // find group of same value
                int j = idx;
                ll min_group = INF;
                while (j < (int)cells.size() && get<0>(cells[j]) == val)
                {
                    int ii = get<1>(cells[j]);
                    int jj = get<2>(cells[j]);
                    min_group = min(min_group, dp[ii][jj]);
                    ++j;
                }
                running_min = min(running_min, min_group);
                // assign start cost for this group
                for (int p = idx; p < j; ++p)
                {
                    int ii = get<1>(cells[p]);
                    int jj = get<2>(cells[p]);
                    start[ii][jj] = min(dp[ii][jj], running_min);
                }
                idx = j;
            }

            // propagate normal moves (right/down) starting from 'start'
            vector<vector<ll>> dp2(m, vector<ll>(n, INF));
            for (int i = 0; i < m; ++i)
            {
                for (int j = 0; j < n; ++j)
                {
                    // starting cost at (i,j)
                    dp2[i][j] = min(dp2[i][j], start[i][j]);
                    if (i + 1 < m && dp2[i][j] < INF)
                        dp2[i + 1][j] = min(dp2[i + 1][j], dp2[i][j] + grid[i + 1][j]);
                    if (j + 1 < n && dp2[i][j] < INF)
                        dp2[i][j + 1] = min(dp2[i][j + 1], dp2[i][j] + grid[i][j + 1]);
                }
            }
            dp.swap(dp2);
        }

        return (int)dp[m - 1][n - 1];
    }
};
