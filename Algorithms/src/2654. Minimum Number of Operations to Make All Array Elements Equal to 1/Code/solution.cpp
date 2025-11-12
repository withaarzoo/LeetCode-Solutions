class Solution
{
public:
    int minOperations(vector<int> &nums)
    {
        int n = (int)nums.size();

        // 1) If there are already ones, we only need to convert others.
        int ones = 0;
        for (int x : nums)
            if (x == 1)
                ones++;
        if (ones > 0)
            return n - ones;

        // 2) If gcd of all numbers > 1, impossible to ever reach 1.
        int g = 0;
        for (int x : nums)
            g = std::gcd(g, x);
        if (g > 1)
            return -1;

        // 3) Find shortest subarray with gcd == 1.
        int best = INT_MAX;
        for (int i = 0; i < n; ++i)
        {
            int cur = 0;
            for (int j = i; j < n; ++j)
            {
                cur = std::gcd(cur, nums[j]);
                if (cur == 1)
                {
                    best = min(best, j - i + 1);
                    break; // extending won't be shorter
                }
            }
        }
        // 4) Cost to create first 1 from that subarray, then spread to all.
        return (best - 1) + (n - 1);
    }
};
