class Solution:
    def minBitwiseArray(self, nums):
        for i in range(len(nums)):
            p = nums[i]

            removable = ((p + 1) & ~p) >> 1

            if removable == 0:
                nums[i] = -1
            else:
                nums[i] = p ^ removable

        return nums
