/**
 * @param {number[]} nums
 * @return {number}
 */
var sumOfSquares = function(nums) {
    const n = nums.length;
    let ans = 0;
    for (let i = 1; i <= n; i++) {
        if (n % i === 0) {
            ans += nums[i - 1] * nums[i - 1];
        }
    }
    return ans;
};