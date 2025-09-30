/**
 * @param {number[]} nums
 * @return {number}
 */
var triangularSum = function (nums) {
  for (let len = nums.length; len > 1; --len) {
    // update in-place; reading nums[i+1] is safe because we haven't changed it yet this round
    for (let i = 0; i < len - 1; ++i) {
      nums[i] = (nums[i] + nums[i + 1]) % 10;
    }
  }
  return nums[0];
};
