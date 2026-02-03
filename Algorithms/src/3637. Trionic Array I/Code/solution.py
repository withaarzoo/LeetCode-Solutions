class Solution:
    def isTrionic(self, nums: List[int]) -> bool:
        n = len(nums)
        i = 0

        # 1) strictly increasing
        while i + 1 < n and nums[i] < nums[i + 1]:
            i += 1
        if i == 0 or i == n - 1:
            return False

        # 2) strictly decreasing
        mid = i
        while i + 1 < n and nums[i] > nums[i + 1]:
            i += 1
        if i == mid or i == n - 1:
            return False

        # 3) strictly increasing again
        while i + 1 < n and nums[i] < nums[i + 1]:
            i += 1

        return i == n - 1
