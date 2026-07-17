class Solution
{
public:
    vector<int> gcdValues(vector<int> &nums, vector<long long> &queries)
    {
        // Find the largest value because every helper array
        // only needs to be built up to this value.
        int mx = *max_element(nums.begin(), nums.end());

        // freq[x] = how many times value x appears.
        vector<int> freq(mx + 1, 0);
        for (int x : nums)
            freq[x]++;

        // exact[g] = number of pairs whose GCD is exactly g.
        vector<long long> exact(mx + 1, 0);

        // Process from largest divisor to smallest.
        for (int g = mx; g >= 1; g--)
        {

            // Count numbers divisible by g.
            long long cnt = 0;
            for (int m = g; m <= mx; m += g)
                cnt += freq[m];

            // Total pairs whose GCD is a multiple of g.
            long long pairs = cnt * (cnt - 1) / 2;

            // Remove pairs already assigned to larger GCDs.
            for (int m = g * 2; m <= mx; m += g)
                pairs -= exact[m];

            exact[g] = pairs;
        }

        // prefix[g] = number of pairs with GCD <= g.
        vector<long long> prefix(mx + 1, 0);
        for (int g = 1; g <= mx; g++)
            prefix[g] = prefix[g - 1] + exact[g];

        vector<int> ans;

        for (long long q : queries)
        {
            // First GCD whose prefix count is greater than q.
            int g = lower_bound(prefix.begin() + 1, prefix.end(), q + 1) - prefix.begin();
            ans.push_back(g);
        }

        return ans;
    }
};