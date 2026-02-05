class Solution:
    def constructTransformedArray(self, nums):
        n = len(nums)
        result = [0] * n

        for i in range(n):
            if nums[i] == 0:
                result[i] = nums[i]
            else:
                target = (i + nums[i]) % n
                result[i] = nums[target]

        return result
