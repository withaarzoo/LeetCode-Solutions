class Solution
{
public:
    long long maxPoints(vector<vector<int>> &points)
    {
        int m = points.size();    // Number of rows
        int n = points[0].size(); // Number of columns

        // DP array to store the maximum points that can be collected up to the current row
        vector<long long> dp(n, 0);

        // Initialize dp array with the values from the first row
        for (int j = 0; j < n; ++j)
        {
            dp[j] = points[0][j];
        }

        // Traverse through each row starting from the second row
        for (int i = 1; i < m; ++i)
        {
            // Arrays to store the maximum values moving from left to right and right to left
            vector<long long> leftMax(n), rightMax(n), newDp(n);

            // Calculate leftMax array
            // leftMax[j] will store the maximum value for positions [0, j] considering the left side
            leftMax[0] = dp[0];
            for (int j = 1; j < n; ++j)
            {
                leftMax[j] = max(leftMax[j - 1], dp[j] + j);
            }

            // Calculate rightMax array
            // rightMax[j] will store the maximum value for positions [j, n-1] considering the right side
            rightMax[n - 1] = dp[n - 1] - (n - 1);
            for (int j = n - 2; j >= 0; --j)
            {
                rightMax[j] = max(rightMax[j + 1], dp[j] - j);
            }

            // Calculate the new DP values for the current row
            for (int j = 0; j < n; ++j)
            {
                // The new DP value for the current cell is calculated by taking the maximum of
                // either the leftMax or rightMax adjusted by the current column index,
                // and then adding the current cell's points
                newDp[j] = max(leftMax[j] - j, rightMax[j] + j) + points[i][j];
            }

            // Update the dp array with the newly calculated values for the next row
            dp = newDp;
        }

        // The answer will be the maximum value in the dp array after processing all rows
        return *max_element(dp.begin(), dp.end());
    }
};
