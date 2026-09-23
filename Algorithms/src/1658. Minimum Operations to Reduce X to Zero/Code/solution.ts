function minOperations(nums: number[], x: number): number {
    let total = 0;
    for (let num of nums) total += num;
    if (total < x) return -1;
    let target = total - x;
    if (target === 0) return nums.length;
    let n = nums.length;
    let left = 0;
    let sum = 0;
    let maxLen = -1;
    for (let right = 0; right < n; ++right) {
        sum += nums[right];
        while (sum > target && left <= right) {
            sum -= nums[left];
            ++left;
        }
        if (sum === target) {
            maxLen = Math.max(maxLen, right - left + 1);
        }
    }
    return maxLen === -1 ? -1 : n - maxLen;
};