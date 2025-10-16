from typing import List

class Solution:
    def findSmallestInteger(self, nums: List[int], value: int) -> int:
        # freq[r] counts numbers with remainder r modulo value
        freq = [0] * value
        for a in nums:
            r = a % value
            if r < 0:
                r += value
            freq[r] += 1

        x = 0
        while True:
            r = x % value
            if freq[r] == 0:
                return x
            freq[r] -= 1
            x += 1
