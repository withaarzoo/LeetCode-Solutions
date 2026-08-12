class Solution
{
public:
    int maxSubarrayLength(vector<int> &nums, int k)
    {
        unordered_map<int, int> freq; // Stores how many times each value appears in the current window.
        int left = 0;                 // Left boundary of the sliding window.
        int ans = 0;                  // Stores the maximum valid window length found so far.

        for (int right = 0; right < nums.size(); right++)
        {
            freq[nums[right]]++; // Add the current element to the window and increase its frequency.

            while (freq[nums[right]] > k)
            {
                freq[nums[left]]--; // Remove the leftmost element because the current window is invalid.
                left++;             // Move the left boundary forward to shrink the window.
            }

            ans = max(ans, right - left + 1); // The window is valid, so update the maximum length.
        }

        return ans; // Return the length of the longest valid subarray.
    }
};