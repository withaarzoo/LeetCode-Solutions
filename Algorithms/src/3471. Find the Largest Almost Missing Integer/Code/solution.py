from typing import List

class Solution:
    def largestInteger(self, nums: List[int], k: int) -> int:
        n = len(nums)

        # nums[i] is between 0 and 50, so a fixed frequency array is enough.
        # Its size never depends on n, so the extra space stays constant.
        freq = [0] * 51

        # Count how many times every value occurs in nums.
        for x in nums:
            freq[x] += 1

        # When k = 1, every subarray contains exactly one element.
        # So a value is valid only when it occurs exactly once in nums.
        if k == 1:
            # Check from the largest possible value to the smallest.
            for x in range(50, -1, -1):
                if freq[x] == 1:
                    return x

            # No value occurs exactly once.
            return -1

        # When k = n, the whole array is the only subarray.
        # Therefore, the largest value in nums is the answer.
        if k == n:
            # max(nums) gives the largest value in the only subarray.
            return max(nums)

        # For 1 < k < n, only the first and last elements
        # can appear in exactly one subarray of size k.
        answer = -1

        # The first value is valid only if it occurs once in the whole array.
        if freq[nums[0]] == 1:
            answer = max(answer, nums[0])

        # The last value is valid only if it occurs once in the whole array.
        if freq[nums[-1]] == 1:
            answer = max(answer, nums[-1])

        # Return the largest valid candidate, or -1 if none exists.
        return answer