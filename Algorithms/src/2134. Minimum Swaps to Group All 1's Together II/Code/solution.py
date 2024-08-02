from typing import List

class Solution:
    def minSwaps(self, nums: List[int]) -> int:
        # Calculate the total number of 1s in the array
        totalOnes = sum(nums)
        
        # If there are no 1s, no swaps are needed
        if totalOnes == 0:
            return 0

        n = len(nums)  # Get the length of the array
        maxOnesInWindow = 0  # Initialize the maximum number of 1s in a window
        currentOnesInWindow = 0  # Initialize the current number of 1s in the current window

        # Calculate the number of 1s in the initial window of size totalOnes
        for i in range(totalOnes):
            currentOnesInWindow += nums[i]

        # Set the initial maxOnesInWindow to the current number of 1s in the initial window
        maxOnesInWindow = currentOnesInWindow

        # Slide the window across the array to find the maximum number of 1s in any window of size totalOnes
        for i in range(1, n):
            # Slide the window to the right: subtract the element that goes out of the window
            currentOnesInWindow -= nums[i - 1]
            # Add the element that comes into the window (using modulo to wrap around the array)
            currentOnesInWindow += nums[(i + totalOnes - 1) % n]
            # Update maxOnesInWindow if the current window has more 1s
            maxOnesInWindow = max(maxOnesInWindow, currentOnesInWindow)

        # The minimum number of swaps is the total number of 1s minus the maximum number of 1s in any window of size totalOnes
        return totalOnes - maxOnesInWindow
