class Solution:
    def search(self, nums: List[int], target: int) -> int:
        
        # Start pointer
        left = 0

        # End pointer
        right = len(nums) - 1

        # Continue until search space becomes empty
        while left <= right:

            # Find middle index safely
            mid = left + (right - left) // 2

            # If target is found
            if nums[mid] == target:
                return mid

            # Check if left half is sorted
            if nums[left] <= nums[mid]:

                # Check whether target lies inside left sorted half
                if nums[left] <= target < nums[mid]:

                    # Search left side
                    right = mid - 1
                else:

                    # Search right side
                    left = mid + 1

            # Otherwise right half is sorted
            else:

                # Check whether target lies inside right sorted half
                if nums[mid] < target <= nums[right]:

                    # Search right side
                    left = mid + 1
                else:

                    # Search left side
                    right = mid - 1

        # Target not found
        return -1