/**
 * @param {number[]} nums
 * @return {number}
 */
var sortPermutation = function(nums) {
    let ans = -1;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] !== i) {
            if (ans === -1) ans = nums[i];
            else ans &= nums[i];
        }
    }
    return ans === -1 ? 0 : ans;
};