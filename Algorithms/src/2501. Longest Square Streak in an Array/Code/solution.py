from typing import List

class Solution:
    def longestSquareStreak(self, nums: List[int]) -> int:
        nums.sort()
        num_set = set(nums)
        max_length = -1

        for num in nums:
            length = 0
            current = num
            while current in num_set:
                length += 1
                current = current * current
                if current > 1e9: break
            if length >= 2:
                max_length = max(max_length, length)
                
        return max_length