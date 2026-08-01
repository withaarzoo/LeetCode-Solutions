class Solution
{
public:
    bool predictTheWinner(vector<int> &nums)
    {
        int n = nums.size();

        // dp[i][j] stores the maximum score difference the current player
        // can achieve over the opponent for subarray [i...j]
        vector<vector<int>> dp(n, vector<int>(n, 0));

        // Base case:
        // If only one number exists, the current player takes it.
        for (int i = 0; i < n; i++)
        {
            dp[i][i] = nums[i];
        }

        // Build the DP table for increasing subarray lengths.
        for (int len = 2; len <= n; len++)
        {
            for (int i = 0; i + len - 1 < n; i++)
            {
                int j = i + len - 1;

                // Choose the left element.
                int takeLeft = nums[i] - dp[i + 1][j];

                // Choose the right element.
                int takeRight = nums[j] - dp[i][j - 1];

                // Keep the better option.
                dp[i][j] = max(takeLeft, takeRight);
            }
        }

        // Non-negative means Player 1 can win or tie.
        return dp[0][n - 1] >= 0;
    }
};