/**
 * @param {number[]} nums
 * @return {number}
 */
var minimizeSum = function(nums) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    return Math.min(nums[n-1] - nums[2], nums[n-3] - nums[0], nums[n-2] - nums[1]);
};