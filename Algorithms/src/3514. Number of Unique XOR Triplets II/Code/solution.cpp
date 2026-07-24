class Solution
{
public:
    int uniqueXorTriplets(vector<int> &nums)
    {
        const int MAX_XOR = 2048;

        // Store whether a value exists in the array.
        vector<bool> present(MAX_XOR, false);
        for (int x : nums)
        {
            present[x] = true;
        }

        // Initially only XOR = 0 is possible before picking any value.
        vector<bool> dp(MAX_XOR, false);
        dp[0] = true;

        // Pick exactly 3 values.
        for (int step = 0; step < 3; step++)
        {
            vector<bool> next(MAX_XOR, false);

            // Try extending every reachable XOR.
            for (int cur = 0; cur < MAX_XOR; cur++)
            {
                if (!dp[cur])
                    continue;

                // Add every value that exists in the array.
                for (int v = 0; v < MAX_XOR; v++)
                {
                    if (present[v])
                    {
                        next[cur ^ v] = true;
                    }
                }
            }

            dp = move(next);
        }

        // Count all unique XOR values.
        int ans = 0;
        for (bool ok : dp)
        {
            if (ok)
                ans++;
        }

        return ans;
    }
};