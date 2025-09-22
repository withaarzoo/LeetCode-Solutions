from typing import List
from collections import Counter

class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        # Count frequencies of all numbers
        cnt = Counter(nums)

        # Find the maximum frequency value
        max_freq = max(cnt.values())

        # Sum frequencies of numbers that have frequency == max_freq
        return sum(v for v in cnt.values() if v == max_freq)
