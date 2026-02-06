class Solution:
    def minRemoval(self, nums, k):
        nums.sort()

        left = 0
        max_keep = 1

        for right in range(len(nums)):
            while nums[right] > nums[left] * k:
                left += 1
            max_keep = max(max_keep, right - left + 1)

        return len(nums) - max_keep
