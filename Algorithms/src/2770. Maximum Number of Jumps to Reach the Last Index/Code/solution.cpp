class Solution
{
public:
    int maximumJumps(vector<int> &nums, int target)
    {

        int n = nums.size();

        // dp[i] = maximum jumps needed to reach index i
        vector<int> dp(n, -1);

        // Starting index needs 0 jumps
        dp[0] = 0;

        // Try every starting index
        for (int i = 0; i < n; i++)
        {

            // If current index is unreachable, skip it
            if (dp[i] == -1)
                continue;

            // Try jumping to every next index
            for (int j = i + 1; j < n; j++)
            {

                // Difference between values
                long long diff = 1LL * nums[j] - nums[i];

                // Check whether jump is valid
                if (diff >= -target && diff <= target)
                {

                    // Update maximum jumps for index j
                    dp[j] = max(dp[j], dp[i] + 1);
                }
            }
        }

        // Final answer
        return dp[n - 1];
    }
};