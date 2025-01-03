class Solution:
    def waysToSplitArray(self, nums: List[int]) -> int:
        total_sum = sum(nums)  # Total sum of the array
        prefix_sum = 0  # Prefix sum
        count = 0  # Count of valid splits

        for i in range(len(nums) - 1):
            prefix_sum += nums[i]
            right_sum = total_sum - prefix_sum
            if prefix_sum >= right_sum:
                count += 1
        
        return count
