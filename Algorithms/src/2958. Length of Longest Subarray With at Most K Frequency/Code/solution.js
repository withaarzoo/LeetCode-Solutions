/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxSubarrayLength = function (nums, k) {
  const freq = new Map(); // Stores the frequency of each value inside the current window.
  let left = 0; // Left boundary of the sliding window.
  let ans = 0; // Stores the longest valid window length.

  for (let right = 0; right < nums.length; right++) {
    // Add nums[right] to the window and increase its frequency.
    freq.set(nums[right], (freq.get(nums[right]) || 0) + 1);

    // Shrink the window while the newly added value appears too many times.
    while (freq.get(nums[right]) > k) {
      // Decrease the frequency of the element leaving from the left.
      freq.set(nums[left], freq.get(nums[left]) - 1);

      // Move the left boundary forward.
      left++;
    }

    // The current window is valid, so update the maximum length.
    ans = Math.max(ans, right - left + 1);
  }

  // Return the length of the longest valid subarray.
  return ans;
};
