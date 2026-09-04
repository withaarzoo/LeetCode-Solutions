class Solution
{
public:
    int firstStableIndex(vector<int> &nums, int k)
    {
        int n = nums.size();

        // suffixMin[i] stores the minimum value from i to n - 1.
        vector<int> suffixMin(n);

        // For the last index, the suffix contains only nums[n - 1].
        suffixMin[n - 1] = nums[n - 1];

        // Build suffix minimums from right to left.
        for (int i = n - 2; i >= 0; --i)
        {
            // Keep the smaller value between the current element
            // and the minimum of the suffix to its right.
            suffixMin[i] = min(nums[i], suffixMin[i + 1]);
        }

        // Store the largest value seen from index 0 to the current index.
        int prefixMax = nums[0];

        // Check every index from smallest to largest.
        for (int i = 0; i < n; ++i)
        {
            // Update the maximum value in nums[0..i].
            prefixMax = max(prefixMax, nums[i]);

            // Calculate the instability score for this index.
            int instability = prefixMax - suffixMin[i];

            // Since indices are checked from left to right,
            // this is automatically the smallest stable index.
            if (instability <= k)
            {
                return i;
            }
        }

        // No stable index exists.
        return -1;
    }
};