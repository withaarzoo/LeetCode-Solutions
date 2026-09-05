class Solution:

    def firstStableIndex(self, nums: list[int], k: int) -> int:
        n = len(nums)  # Store the number of elements in the array.

        prefixMax = [0] * n  # prefixMax[i] stores max(nums[0..i]).
        suffixMin = [0] * n  # suffixMin[i] stores min(nums[i..n-1]).

        prefixMax[0] = nums[0]  # For index 0, the prefix contains only nums[0].

        for i in range(1, n):  # Build prefix maximums from left to right.
            prefixMax[i] = max(prefixMax[i - 1], nums[i])  # Keep the largest value seen so far.

        suffixMin[n - 1] = nums[n - 1]  # For the last index, the suffix contains only nums[n-1].

        for i in range(n - 2, -1, -1):  # Build suffix minimums from right to left.
            suffixMin[i] = min(suffixMin[i + 1], nums[i])  # Keep the smallest value in the suffix.

        for i in range(n):  # Check indices from smallest to largest.
            instability = prefixMax[i] - suffixMin[i]  # Calculate the score for index i.

            if instability <= k:  # The index is stable when its score is at most k.
                return i  # This is the smallest stable index because we scan left to right.

        return -1  # No stable index was found.