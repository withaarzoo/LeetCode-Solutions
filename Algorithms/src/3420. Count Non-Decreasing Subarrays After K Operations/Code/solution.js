/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var countNonDecreasingSubarrays = function(nums, k) {
    let ans = 0;
    let cost = 0;
    const dq = [];
    const n = nums.length;
    for (let i = n - 1, j = n - 1; i >= 0; --i) {
        let num = nums[i];
        let count = 1;
        while (dq.length && dq[dq.length - 1][0] < num) {
            const [nextNum, nextCount] = dq.pop();
            count += nextCount;
            cost += (num - nextNum) * nextCount;
        }
        dq.push([num, count]);
        while (cost > k) {
            const [rightmostNum, rightmostCount] = dq.shift();
            cost -= rightmostNum - nums[j];
            --j;
            if (rightmostCount > 1) {
                dq.unshift([rightmostNum, rightmostCount - 1]);
            }
        }
        ans += j - i + 1;
    }
    return ans;
};