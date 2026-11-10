function climbStairs(n: number, costs: number[]): number {
    const dp: number[] = new Array(n + 1).fill(Number.MAX_SAFE_INTEGER);
    dp[0] = 0;
    for (let j = 1; j <= n; j++) {
        for (let k = 1; k <= 3; k++) {
            const i = j - k;
            if (i >= 0) {
                const jump = (j - i) * (j - i);
                dp[j] = Math.min(dp[j], dp[i] + costs[j - 1] + jump);
            }
        }
    }
    return dp[n];
};