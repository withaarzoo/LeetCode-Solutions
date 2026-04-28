class Solution:
    def minOperations(self, grid: List[List[int]], x: int) -> int:
        nums = []

        # Flatten grid
        for row in grid:
            nums.extend(row)

        # Check feasibility
        rem = nums[0] % x
        for num in nums:
            if num % x != rem:
                return -1

        # Sort
        nums.sort()

        # Median
        median = nums[len(nums) // 2]

        # Count operations
        ops = 0
        for num in nums:
            ops += abs(num - median) // x

        return ops