/**
 * @param {number[]} nums
 * @return {number}
 */
var sumCounts = function(nums) {
    const n = nums.length;
    let res = 0;
    for (let i = 0; i < n; i++) {
        const s = new Set();
        for (let j = i; j < n; j++) {
            s.add(nums[j]);
            const d = s.size;
            res += d * d;
        }
    }
    return res;
};