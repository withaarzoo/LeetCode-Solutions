class Solution:
    def check(self, nums: List[int]) -> bool:
        
        n = len(nums)

        # Counts how many times order decreases
        count = 0

        # Traverse all indices
        for i in range(n):

            # Compare current element with next element
            # % n helps compare last with first
            if nums[i] > nums[(i + 1) % n]:
                count += 1

            # More than one decrease means invalid
            if count > 1:
                return False

        # Valid sorted rotated array
        return True