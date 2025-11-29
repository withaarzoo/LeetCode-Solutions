from typing import List

class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        # Calculate total sum of the array
        total_sum = sum(nums)
        
        # Minimum operations is the remainder of total_sum when divided by k
        remainder = total_sum % k
        
        return remainder
