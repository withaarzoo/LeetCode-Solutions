class Solution
{
public:
    vector<vector<int>> rangeAddQueries(int n, vector<vector<int>> &queries)
    {
        // diff is (n+1) x (n+1) to safely handle r2+1 and c2+1 indices
        vector<vector<int>> diff(n + 1, vector<int>(n + 1, 0));

        // Apply each query as four corner updates (2D difference technique)
        for (const auto &q : queries)
        {
            int r1 = q[0], c1 = q[1], r2 = q[2], c2 = q[3];
            diff[r1][c1] += 1;
            diff[r1][c2 + 1] -= 1;
            diff[r2 + 1][c1] -= 1;
            diff[r2 + 1][c2 + 1] += 1;
        }

        // Convert diff into the final matrix using 2D prefix sum
        vector<vector<int>> res(n, vector<int>(n, 0));
        for (int i = 0; i < n; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                int up = (i > 0) ? diff[i - 1][j] : 0;
                int left = (j > 0) ? diff[i][j - 1] : 0;
                int diag = (i > 0 && j > 0) ? diff[i - 1][j - 1] : 0;
                diff[i][j] = diff[i][j] + up + left - diag;
                res[i][j] = diff[i][j];
            }
        }
        return res;
    }
};
