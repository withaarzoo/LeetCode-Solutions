class Solution:
    def findMin(self, nums: List[int]) -> int:
        
        # Initialize pointers
        left = 0
        right = len(nums) - 1

        # Binary search loop
        while left < right:

            # Middle index
            mid = left + (right - left) // 2

            # Minimum lies on left side including mid
            if nums[mid] < nums[right]:
                right = mid

            # Minimum lies on right side
            elif nums[mid] > nums[right]:
                left = mid + 1

            # Duplicate case
            # Remove one element safely
            else:
                right -= 1

        # Left points to minimum element
        return nums[left]