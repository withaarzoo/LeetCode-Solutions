class Solution
{
public:
    int stoneGameII(vector<int> &piles)
    {
        int n = piles.size(); // Get the number of piles
        // Create a DP table initialized with 0. dp[i][M] represents the maximum number of stones the current player can get starting at pile i with the given M.
        vector<vector<int>> dp(n + 1, vector<int>(n + 1, 0));
        // Create a suffix sum array. suffixSum[i] represents the total number of stones from pile i to the end.
        vector<int> suffixSum(n + 1, 0);

        // Calculate the suffix sums
        for (int i = n - 1; i >= 0; --i)
        {
            // suffixSum[i] = current pile stones + suffixSum from the next pile
            suffixSum[i] = suffixSum[i + 1] + piles[i];
        }

        // Fill the DP table
        for (int i = n - 1; i >= 0; --i)
        { // Iterate through piles in reverse order
            for (int M = 1; M <= n; ++M)
            { // Iterate through possible values of M
                // Consider taking X piles where 1 <= X <= 2 * M
                for (int X = 1; X <= 2 * M && i + X <= n; ++X)
                {
                    // Calculate the maximum stones current player can take
                    // It is the maximum between the current value and the stones from the suffix minus what the next player can get.
                    dp[i][M] = max(dp[i][M], suffixSum[i] - dp[i + X][max(M, X)]);
                }
            }
        }

        // The result is the maximum stones the first player can get starting from pile 0 with M = 1
        return dp[0][1];
    }
};
