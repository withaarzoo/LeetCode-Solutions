/**
 * @param {number[]} nums
 * @return {boolean}
 */
var check = function (nums) {
  const n = nums.length;

  // Counts how many decreasing points exist
  let count = 0;

  // Traverse the array
  for (let i = 0; i < n; i++) {
    // Compare current with next element
    // % n connects last element to first
    if (nums[i] > nums[(i + 1) % n]) {
      count++;
    }

    // Invalid if order breaks more than once
    if (count > 1) {
      return false;
    }
  }

  // Valid sorted rotated array
  return true;
};
