/**
 * @param {number[]} nums
 * @return {number}
 */
var findMin = function (nums) {
  // Start pointer
  let left = 0;

  // End pointer
  let right = nums.length - 1;

  // Run Binary Search
  while (left < right) {
    // Middle index
    let mid = Math.floor((left + right) / 2);

    // Minimum exists on right side
    if (nums[mid] > nums[right]) {
      // Remove left sorted half
      left = mid + 1;
    } else {
      // Minimum can still be at mid
      right = mid;
    }
  }

  // Final answer
  return nums[left];
};
