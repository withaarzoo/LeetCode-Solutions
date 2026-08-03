class Solution
{
public:
    string stoneGameIII(vector<int> &stoneValue)
    {
        int n = stoneValue.size();

        // dp[i] = maximum score difference current player can achieve
        // starting from index i
        vector<int> dp(n + 1, 0);

        // Process from back because current state depends on future states
        for (int i = n - 1; i >= 0; i--)
        {

            // Initialize with a very small value since we are maximizing
            dp[i] = INT_MIN;

            int sum = 0;

            // Try taking 1, 2 and 3 stones
            for (int j = i; j < min(n, i + 3); j++)
            {

                // Current player's collected score
                sum += stoneValue[j];

                // Current score difference =
                // current collected score - opponent's best difference
                dp[i] = max(dp[i], sum - dp[j + 1]);
            }
        }

        // Decide the winner using the final score difference
        if (dp[0] > 0)
            return "Alice";
        if (dp[0] < 0)
            return "Bob";
        return "Tie";
    }
};