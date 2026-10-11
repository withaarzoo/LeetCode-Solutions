function sumOfSquares(nums: number[]): number {
    const n = nums.length;
    let ans = 0;
    for (let i = 1; i <= n; i++) {
        if (n % i === 0) {
            ans += nums[i - 1] * nums[i - 1];
        }
    }
    return ans;
};