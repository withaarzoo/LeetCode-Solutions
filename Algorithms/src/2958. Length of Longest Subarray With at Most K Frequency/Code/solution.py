class Solution:
    def maxSubarrayLength(self, nums: List[int], k: int) -> int:
        freq = {}  # Stores the frequency of each value inside the current window.
        left = 0   # Left boundary of the sliding window.
        ans = 0    # Stores the longest valid window length.

        for right in range(len(nums)):
            # Add nums[right] to the window and increase its frequency.
            freq[nums[right]] = freq.get(nums[right], 0) + 1

            # Shrink the window while the newly added value appears more than k times.
            while freq[nums[right]] > k:
                # Decrease the frequency of the element leaving from the left.
                freq[nums[left]] -= 1

                # Move the left boundary forward.
                left += 1

            # The current window is valid, so update the maximum length.
            ans = max(ans, right - left + 1)

        # Return the length of the longest valid subarray.
        return ans