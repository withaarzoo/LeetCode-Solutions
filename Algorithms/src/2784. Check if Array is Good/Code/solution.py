class Solution:
    def isGood(self, nums: List[int]) -> bool:

        # Sort the array
        nums.sort()

        # Length of array
        n = len(nums)

        # Maximum element
        mx = nums[-1]

        # Size must be mx + 1
        if n != mx + 1:
            return False

        # Check numbers from 1 to mx
        for i in range(n - 1):

            # Expected value is i + 1
            if nums[i] != i + 1:
                return False

        # Last element must also be mx
        return nums[-1] == mx