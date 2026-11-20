/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSelectedElements = function(nums) {
    nums.sort((a, b) => a - b);
    const dp = new Map();
    let ans = 0;
    for (const num of nums) {
        dp.set(num + 1, (dp.get(num) || 0) + 1);
        dp.set(num, (dp.get(num - 1) || 0) + 1);
        ans = Math.max(ans, dp.get(num), dp.get(num + 1));
    }
    return ans;
};