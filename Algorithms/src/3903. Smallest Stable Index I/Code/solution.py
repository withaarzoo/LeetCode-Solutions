class Solution:

    def firstStableIndex(self, nums: list[int], k: int) -> int:
        n = len(nums)

        # suffix_min[i] stores the minimum value from i to the end.
        suffix_min = [0] * n

        # The last suffix contains only the last element.
        suffix_min[n - 1] = nums[n - 1]

        # Build suffix minimums from right to left.
        for i in range(n - 2, -1, -1):
            # Keep the smaller value between nums[i]
            # and the minimum value already found to its right.
            suffix_min[i] = min(nums[i], suffix_min[i + 1])

        # Store the largest value seen from index 0 to the current index.
        prefix_max = nums[0]

        # Check every index from smallest to largest.
        for i in range(n):
            # Update the maximum value in nums[0..i].
            prefix_max = max(prefix_max, nums[i])

            # Calculate the instability score at index i.
            instability = prefix_max - suffix_min[i]

            # This is the first stable index because we scan from left to right.
            if instability <= k:
                return i

        # No stable index exists.
        return -1