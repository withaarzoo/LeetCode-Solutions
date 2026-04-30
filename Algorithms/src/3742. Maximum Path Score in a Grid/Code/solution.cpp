class Solution
{
public:
    int maxPathScore(vector<vector<int>> &grid, int k)
    {
        int m = grid.size();    // Number of rows in the grid.
        int n = grid[0].size(); // Number of columns in the grid.
        const int NEG = -1e9;   // Sentinel for impossible states.

        // prev[j][c] = best score at column j in the previous row using exact cost c.
        vector<vector<int>> prev(n, vector<int>(k + 1, NEG));

        for (int i = 0; i < m; ++i)
        {
            // Rebuild the current row from scratch so stale values never leak in.
            vector<vector<int>> curr(n, vector<int>(k + 1, NEG));

            for (int j = 0; j < n; ++j)
            {
                int gain = grid[i][j];         // Score gained by stepping on this cell.
                int need = (gain > 0 ? 1 : 0); // Cost spent by this cell: 0 for 0, 1 for 1/2.

                // A path to (i, j) cannot spend more than i + j budget points.
                int limit = min(k, i + j);

                // The start cell is fixed by the statement: it is always 0.
                if (i == 0 && j == 0)
                {
                    curr[0][0] = 0; // Start with zero score and zero cost.
                    continue;
                }

                for (int c = need; c <= limit; ++c)
                {
                    int best = NEG;

                    // Take the path from above, then pay the current cell cost and add its score.
                    if (i > 0 && prev[j][c - need] != NEG)
                    {
                        best = max(best, prev[j][c - need] + gain);
                    }

                    // Take the path from the left, then pay the current cell cost and add its score.
                    if (j > 0 && curr[j - 1][c - need] != NEG)
                    {
                        best = max(best, curr[j - 1][c - need] + gain);
                    }

                    curr[j][c] = best; // Store the best exact-cost result for this cell.
                }
            }

            prev.swap(curr); // Move the current row into prev for the next round.
        }

        int ans = NEG; // Best score among all valid costs at the end cell.
        for (int c = 0; c <= k; ++c)
        {
            ans = max(ans, prev[n - 1][c]);
        }

        return ans < 0 ? -1 : ans; // If nothing is reachable, the answer is -1.
    }
};