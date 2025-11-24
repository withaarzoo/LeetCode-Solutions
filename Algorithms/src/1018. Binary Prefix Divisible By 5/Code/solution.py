from typing import List

class Solution:
    def prefixesDivBy5(self, nums: List[int]) -> List[bool]:
        ans: List[bool] = []
        rem = 0                          # remainder of current prefix modulo 5
        
        for bit in nums:
            # Shift left in binary and add bit
            rem = (rem * 2 + bit) % 5
            # True if divisible by 5
            ans.append(rem == 0)
        
        return ans
