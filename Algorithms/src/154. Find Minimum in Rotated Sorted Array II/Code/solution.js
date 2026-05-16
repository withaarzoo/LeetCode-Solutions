/**
 * @param {number[]} nums
 * @return {number}
 */
var findMin = function (nums) {
  // Initialize pointers
  let left = 0;
  let right = nums.length - 1;

  // Binary search
  while (left < right) {
    // Middle index
    let mid = Math.floor((left + right) / 2);

    // Minimum is on left side including mid
    if (nums[mid] < nums[right]) {
      right = mid;
    }

    // Minimum is on right side
    else if (nums[mid] > nums[right]) {
      left = mid + 1;
    }

    // Duplicate case
    else {
      right--;
    }
  }

  // Return minimum element
  return nums[left];
};
