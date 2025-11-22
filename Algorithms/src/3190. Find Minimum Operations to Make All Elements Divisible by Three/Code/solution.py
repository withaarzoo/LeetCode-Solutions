from typing import List

class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        operations = 0
        
        # Check each number in nums
        for x in nums:
            # If x is not divisible by 3, we need one operation
            if x % 3 != 0:
                operations += 1
        
        return operations
