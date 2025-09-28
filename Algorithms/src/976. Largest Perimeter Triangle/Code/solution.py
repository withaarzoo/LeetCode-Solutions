from typing import List

class Solution:
    def largestPerimeter(self, nums: List[int]) -> int:
        # Sort the list in ascending order
        nums.sort()
        # Iterate from the end checking each triple
        for i in range(len(nums) - 1, 1, -1):
            a = nums[i]       # largest of the three
            b = nums[i - 1]
            c = nums[i - 2]
            # If the sum of the two smaller sides is greater than the largest, it's a valid triangle
            if b + c > a:
                return a + b + c
        # No valid triangle found
        return 0
