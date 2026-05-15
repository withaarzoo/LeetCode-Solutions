class Solution:
    def findMin(self, nums: List[int]) -> int:
        
        # Left pointer
        left = 0
        
        # Right pointer
        right = len(nums) - 1

        # Binary Search loop
        while left < right:
            
            # Middle index
            mid = left + (right - left) // 2

            # Minimum lies on right side
            if nums[mid] > nums[right]:
                
                # Move left pointer
                left = mid + 1
            else:
                
                # Minimum may be at mid
                right = mid

        # Return minimum element
        return nums[left]