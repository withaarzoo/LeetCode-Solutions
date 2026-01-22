class Solution:
    def minimumPairRemoval(self, nums):
        operations = 0

        def is_sorted():
            for i in range(1, len(nums)):
                if nums[i] < nums[i - 1]:
                    return False
            return True

        while not is_sorted():
            min_sum = float('inf')
            index = 0

            for i in range(len(nums) - 1):
                s = nums[i] + nums[i + 1]
                if s < min_sum:
                    min_sum = s
                    index = i

            nums[index] = min_sum
            nums.pop(index + 1)
            operations += 1

        return operations
