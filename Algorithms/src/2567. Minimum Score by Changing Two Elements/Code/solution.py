class Solution:
    def minimizeSum(self, nums: List[int]) -> int:
        nums.sort()
        n = len(nums)
        return min(nums[n-1] - nums[2], nums[n-3] - nums[0], nums[n-2] - nums[1])