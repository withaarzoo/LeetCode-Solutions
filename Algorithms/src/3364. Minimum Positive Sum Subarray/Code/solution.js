/**
 * @param {number[]} nums
 * @param {number} l
 * @param {number} r
 * @return {number}
 */
var minimumSumSubarray = function(nums, l, r) {
    const n = nums.length;
    let ans = Infinity;
    for (let i = 0; i < n; i++) {
        let sum = 0;
        for (let j = i; j < n; j++) {
            sum += nums[j];
            const len = j - i + 1;
            if (len >= l && len <= r && sum > 0) {
                ans = Math.min(ans, sum);
            }
        }
    }
    return ans === Infinity ? -1 : ans;
};