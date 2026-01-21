/**
 * @param {number[]} nums
 * @return {number[]}
 */
var minBitwiseArray = function (nums) {
  for (let i = 0; i < nums.length; i++) {
    let p = nums[i];

    let removable = ((p + 1) & ~p) >> 1;

    if (removable === 0) {
      nums[i] = -1;
    } else {
      nums[i] = p ^ removable;
    }
  }
  return nums;
};
