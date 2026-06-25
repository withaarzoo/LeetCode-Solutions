class Solution
{
public:
    int countMajoritySubarrays(vector<int> &nums, int target)
    {
        int n = nums.size();
        int ans = 0;

        // Try every possible starting index
        for (int left = 0; left < n; left++)
        {
            int countTarget = 0;

            // Extend the subarray one element at a time
            for (int right = left; right < n; right++)
            {

                // Update target frequency for the current subarray
                if (nums[right] == target)
                    countTarget++;

                // Current subarray length
                int len = right - left + 1;

                // Target is a majority if it appears more than half the time
                if (2 * countTarget > len)
                    ans++;
            }
        }

        return ans;
    }
};