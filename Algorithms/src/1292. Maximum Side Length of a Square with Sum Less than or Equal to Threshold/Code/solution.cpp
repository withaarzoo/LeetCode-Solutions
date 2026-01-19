class Solution
{
public:
    int maxSideLength(vector<vector<int>> &mat, int threshold)
    {
        int m = mat.size(), n = mat[0].size();

        // Prefix sum matrix
        vector<vector<int>> pre(m + 1, vector<int>(n + 1, 0));

        // Build prefix sum
        for (int i = 1; i <= m; i++)
        {
            for (int j = 1; j <= n; j++)
            {
                pre[i][j] = mat[i - 1][j - 1] + pre[i - 1][j] + pre[i][j - 1] - pre[i - 1][j - 1];
            }
        }

        int left = 0, right = min(m, n), ans = 0;

        while (left <= right)
        {
            int mid = (left + right) / 2;
            bool found = false;

            // Check all possible squares of size mid
            for (int i = mid; i <= m && !found; i++)
            {
                for (int j = mid; j <= n; j++)
                {
                    int sum = pre[i][j] - pre[i - mid][j] - pre[i][j - mid] + pre[i - mid][j - mid];

                    if (sum <= threshold)
                    {
                        found = true;
                        break;
                    }
                }
            }

            if (found)
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
