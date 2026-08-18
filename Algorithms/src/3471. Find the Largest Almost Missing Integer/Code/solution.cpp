class Solution
{
public:
    int largestInteger(vector<int> &nums, int k)
    {
        int n = nums.size();

        // nums[i] is between 0 and 50, so a fixed frequency array is enough.
        // This uses constant extra space instead of a hash table.
        int freq[51] = {};

        // Count how many times every value appears in nums.
        for (int x : nums)
        {
            freq[x]++;
        }

        // When k = 1, every subarray contains exactly one element.
        // Therefore, a value is almost missing exactly when it appears once.
        if (k == 1)
        {
            // Search from 50 down to 0 so the first valid value is the largest.
            for (int x = 50; x >= 0; x--)
            {
                if (freq[x] == 1)
                {
                    return x;
                }
            }

            // No value occurs exactly once.
            return -1;
        }

        // When k = n, there is only one subarray: the entire array.
        // Every distinct value appears in exactly one subarray.
        if (k == n)
        {
            int answer = 0;

            // Find the largest value in the whole array.
            for (int x : nums)
            {
                answer = max(answer, x);
            }

            return answer;
        }

        // For 1 < k < n, only the first and last elements
        // can belong to exactly one subarray of size k.
        int answer = -1;

        // The first element is valid only if its value occurs once in nums.
        if (freq[nums[0]] == 1)
        {
            answer = max(answer, nums[0]);
        }

        // The last element is valid only if its value occurs once in nums.
        if (freq[nums[n - 1]] == 1)
        {
            answer = max(answer, nums[n - 1]);
        }

        // Return the largest valid endpoint, or -1 if neither is valid.
        return answer;
    }
};