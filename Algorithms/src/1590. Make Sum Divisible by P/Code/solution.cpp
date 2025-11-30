class Solution
{
public:
    int minSubarray(vector<int> &nums, int p)
    {
        long long total = 0;
        for (int x : nums)
        {
            total = (total + x) % p; // keep total always modulo p
        }

        int need = (int)total;
        if (need == 0)
            return 0; // already divisible

        int n = nums.size();
        unordered_map<int, int> lastIndex;
        lastIndex.reserve(n * 2); // small optimization
        lastIndex[0] = -1;        // prefix before start

        int ans = n;
        long long prefix = 0;

        for (int i = 0; i < n; ++i)
        {
            prefix = (prefix + nums[i]) % p;
            int prefMod = (int)prefix;

            // We want a previous prefix with remainder "target"
            int target = prefMod - need;
            if (target < 0)
                target += p; // (prefMod - need + p) % p

            if (lastIndex.count(target))
            {
                ans = min(ans, i - lastIndex[target]);
            }

            // Store / update the latest index for this remainder
            lastIndex[prefMod] = i;
        }

        return ans == n ? -1 : ans;
    }
};
