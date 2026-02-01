/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumCost = function (nums) {
  let first = nums[0];

  // Sort remaining elements
  let rest = nums.slice(1).sort((a, b) => a - b);

  return first + rest[0] + rest[1];
};
