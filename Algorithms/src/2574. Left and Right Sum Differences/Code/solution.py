class Solution:
    def leftRightDifference(self, nums: List[int]) -> List[int]:
        
        # Total sum of all elements
        right_sum = sum(nums)

        # Sum of elements on the left side
        left_sum = 0

        # Result array
        ans = [0] * len(nums)

        for i in range(len(nums)):

            # Remove current element so right_sum contains
            # only elements to the right
            right_sum -= nums[i]

            # Store absolute difference
            ans[i] = abs(left_sum - right_sum)

            # Add current element to left_sum
            left_sum += nums[i]

        return ans