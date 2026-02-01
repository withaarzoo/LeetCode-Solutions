class Solution:
    def minimumCost(self, nums: List[int]) -> int:
        first = nums[0]

        # Sort remaining elements
        nums[1:] = sorted(nums[1:])

        return first + nums[1] + nums[2]
