class Solution:
    def countPairs(self, nums, mid):
        """
        This function counts how many pairs (i, j) exist in the sorted array `nums` 
        such that the difference between nums[j] and nums[i] is less than or equal to `mid`.
        """
        count = 0  # Initialize the count of pairs
        j = 0      # Initialize the pointer j which will be used to find the upper bound of the pair

        # Iterate through each element in the array
        for i in range(len(nums)):
            # Move the pointer j to find the maximum index where the difference is <= mid
            while j < len(nums) and nums[j] - nums[i] <= mid:
                j += 1
            # The number of valid pairs with i as the first element is j - i - 1
            count += j - i - 1

        return count  # Return the total count of valid pairs

    def smallestDistancePair(self, nums: List[int], k: int) -> int:
        """
        This function finds the k-th smallest distance pair (difference between any two elements)
        in the array `nums`.
        """
        nums.sort()  # Sort the array to enable binary search on pair differences
        
        # Initialize the binary search range for the distance between pairs
        low, high = 0, nums[-1] - nums[0]

        # Perform binary search
        while low < high:
            mid = (low + high) // 2  # Calculate the mid-point of the current range

            # Count the number of pairs with difference <= mid
            if self.countPairs(nums, mid) >= k:
                # If there are at least k pairs with difference <= mid, 
                # the answer lies in the lower half including mid
                high = mid
            else:
                # Otherwise, the answer lies in the upper half excluding mid
                low = mid + 1

        return low  # The k-th smallest distance is found
