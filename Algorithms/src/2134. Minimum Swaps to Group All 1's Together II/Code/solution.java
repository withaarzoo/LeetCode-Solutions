class Solution {
    public int minSwaps(int[] nums) {
        // Calculate the total number of 1s in the array
        int totalOnes = 0;
        for (int num : nums) {
            if (num == 1)
                totalOnes++;
        }
        // If there are no 1s, no swaps are needed
        if (totalOnes == 0)
            return 0;

        int n = nums.length;
        // Initialize variables to track the maximum number of 1s in any window and the
        // current number of 1s in the current window
        int maxOnesInWindow = 0, currentOnesInWindow = 0;

        // Calculate the number of 1s in the initial window of size 'totalOnes'
        for (int i = 0; i < totalOnes; i++) {
            currentOnesInWindow += nums[i];
        }

        // Set the maximum number of 1s in the window to the initial window's count
        maxOnesInWindow = currentOnesInWindow;

        // Slide the window across the array to find the maximum number of 1s in any
        // window of size 'totalOnes'
        for (int i = 1; i < n; i++) {
            // Remove the element that is sliding out of the window and add the new element
            // that is sliding into the window
            currentOnesInWindow -= nums[i - 1];
            currentOnesInWindow += nums[(i + totalOnes - 1) % n];
            // Update the maximum number of 1s found in any window so far
            maxOnesInWindow = Math.max(maxOnesInWindow, currentOnesInWindow);
        }

        // The minimum number of swaps needed is the difference between the total number
        // of 1s and the maximum number of 1s in any window
        return totalOnes - maxOnesInWindow;
    }
}
