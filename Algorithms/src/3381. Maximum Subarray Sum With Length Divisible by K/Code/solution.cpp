class Solution
{
public:
    long long maxSubarraySum(vector<int> &nums, int k)
    {
        int n = nums.size();

        const long long INF = (long long)4e18; // large number
        vector<long long> minPref(k, INF);

        long long prefix = 0;
        long long ans = -INF;

        // prefix index 0 has sum = 0 and remainder 0
        minPref[0] = 0;

        for (int i = 0; i < n; ++i)
        {
            prefix += (long long)nums[i];
            int rem = (i + 1) % k; // prefix index is i+1

            // if we've seen this remainder before, try forming a subarray
            if (minPref[rem] != INF)
            {
                ans = max(ans, prefix - minPref[rem]);
            }

            // update minimum prefix for this remainder
            if (prefix < minPref[rem])
            {
                minPref[rem] = prefix;
            }
        }

        return ans;
    }
};
