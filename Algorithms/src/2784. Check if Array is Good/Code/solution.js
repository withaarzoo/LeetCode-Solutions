/**
 * @param {number[]} nums
 * @return {boolean}
 */
var isGood = function (nums) {
  // Sort numbers in ascending order
  nums.sort((a, b) => a - b);

  // Length of array
  let n = nums.length;

  // Maximum element
  let mx = nums[n - 1];

  // Size must be mx + 1
  if (n !== mx + 1) {
    return false;
  }

  // Check numbers from 1 to mx
  for (let i = 0; i < n - 1; i++) {
    // Expected value is i + 1
    if (nums[i] !== i + 1) {
      return false;
    }
  }

  // Last element must also be mx
  return nums[n - 1] === mx;
};
