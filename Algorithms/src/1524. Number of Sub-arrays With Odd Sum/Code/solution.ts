function numOfSubarrays(arr: number[]): number {
    const MOD = 1e9 + 7;
    let even = 1, odd = 0;
    let ans = 0;
    let sum = 0;
    for (const x of arr) {
        sum = (sum + x) % 2;
        if (sum === 1) {
            ans = (ans + even) % MOD;
            odd++;
        } else {
            ans = (ans + odd) % MOD;
            even++;
        }
    }
    return ans;
};