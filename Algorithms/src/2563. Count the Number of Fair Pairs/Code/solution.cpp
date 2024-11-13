#include <vector>
#include <algorithm>

class Solution
{
public:
    long long countFairPairs(std::vector<int> &nums, int lower, int upper)
    {
        // Sort the array for easier range search
        std::sort(nums.begin(), nums.end());
        long long count = 0;
        int n = nums.size();

        for (int i = 0; i < n - 1; ++i)
        {
            int min_val = lower - nums[i];
            int max_val = upper - nums[i];

            // Binary search to find valid range for nums[j] where i < j
            auto start = std::lower_bound(nums.begin() + i + 1, nums.end(), min_val);
            auto end = std::upper_bound(nums.begin() + i + 1, nums.end(), max_val);

            count += (end - start);
        }

        return count;
    }
};