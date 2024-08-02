/**
 * Function to find the minimum number of swaps required to group all 1's together
 * @param {number[]} nums - Array containing 0's and 1's
 * @return {number} - Minimum number of swaps required
 */
var minSwaps = function (nums) {
  // Calculate the total number of 1's in the array
  const totalOnes = nums.reduce((acc, num) => acc + num, 0);

  // If there are no 1's in the array, no swaps are needed
  if (totalOnes === 0) return 0;

  // Length of the input array
  const n = nums.length;

  // Variables to track the maximum number of 1's in any window and the current number of 1's in the current window
  let maxOnesInWindow = 0,
    currentOnesInWindow = 0;

  // Initialize the first window of size equal to the total number of 1's
  for (let i = 0; i < totalOnes; i++) {
    currentOnesInWindow += nums[i];
  }

  // Set the maximum number of 1's in the window to the current number of 1's in the initial window
  maxOnesInWindow = currentOnesInWindow;

  // Slide the window across the array
  for (let i = 1; i < n; i++) {
    // Subtract the element that is sliding out of the window
    currentOnesInWindow -= nums[i - 1];
    // Add the next element in the array, wrapping around using modulo operator
    currentOnesInWindow += nums[(i + totalOnes - 1) % n];
    // Update the maximum number of 1's in any window
    maxOnesInWindow = Math.max(maxOnesInWindow, currentOnesInWindow);
  }

  // The minimum swaps needed is the total number of 1's minus the maximum number of 1's in any window
  return totalOnes - maxOnesInWindow;
};
