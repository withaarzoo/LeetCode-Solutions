from typing import List

class Solution:
    def longestSubsequence(self, nums: List[int]) -> int:
        xor_value = 0  # Stores the XOR of all elements.
        has_non_zero = False  # Tracks whether at least one element is non-zero.

        for x in nums:
            xor_value ^= x  # Add the current element to the total XOR.

            if x != 0:
                has_non_zero = True  # A non-zero element can be removed if needed.

        if xor_value != 0:
            return len(nums)  # The entire array already has a non-zero XOR.

        if has_non_zero:
            return len(nums) - 1  # Remove one non-zero element to make XOR non-zero.

        return 0  # All elements are zero, so every subsequence has XOR zero.