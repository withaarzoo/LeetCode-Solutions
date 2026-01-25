class Solution:
    def minimumDifference(self, nums: List[int], k: int) -> int:
        # If only one student is selected
        if k == 1:
            return 0

        # Step 1: Sort the array
        nums.sort()

        min_diff = float('inf')

        # Step 2: Sliding window
        for i in range(len(nums) - k + 1):
            diff = nums[i + k - 1] - nums[i]
            min_diff = min(min_diff, diff)

        return min_diff
