class Solution
{
public:
    int firstStableIndex(vector<int> &nums, int k)
    {
        int n = nums.size(); // Store the number of elements in the array.

        vector<int> prefixMax(n); // prefixMax[i] stores max(nums[0..i]).
        vector<int> suffixMin(n); // suffixMin[i] stores min(nums[i..n-1]).

        prefixMax[0] = nums[0]; // For index 0, the prefix contains only nums[0].

        for (int i = 1; i < n; i++)
        {                                                  // Build prefix maximums from left to right.
            prefixMax[i] = max(prefixMax[i - 1], nums[i]); // Keep the largest value seen so far.
        }

        suffixMin[n - 1] = nums[n - 1]; // For the last index, the suffix contains only nums[n-1].

        for (int i = n - 2; i >= 0; i--)
        {                                                  // Build suffix minimums from right to left.
            suffixMin[i] = min(suffixMin[i + 1], nums[i]); // Keep the smallest value in the suffix.
        }

        for (int i = 0; i < n; i++)
        {                                                  // Check every index from smallest to largest.
            int instability = prefixMax[i] - suffixMin[i]; // Calculate the instability score at i.

            if (instability <= k)
            {             // A score at most k means this index is stable.
                return i; // Since we scan left to right, this is the smallest stable index.
            }
        }

        return -1; // No index satisfies the stability condition.
    }
};