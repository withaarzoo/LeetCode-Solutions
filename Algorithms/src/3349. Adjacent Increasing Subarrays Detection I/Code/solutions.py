from typing import List

class Solution:
    def hasIncreasingSubarrays(self, nums: List[int], k: int) -> bool:
        n = len(nums)
        if 2 * k > n:
            return False
        
        # next_inc[i] = number of consecutive increasing adjacent pairs starting at i
        next_inc = [0] * n
        for i in range(n - 2, -1, -1):
            if nums[i] < nums[i + 1]:
                next_inc[i] = next_inc[i + 1] + 1
            else:
                next_inc[i] = 0
        
        need = k - 1
        for i in range(0, n - 2 * k + 1):
            if next_inc[i] >= need and next_inc[i + k] >= need:
                return True
        return False
