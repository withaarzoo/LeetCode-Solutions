function smallestUniqueSubarray(nums: number[]): number {
    const n = nums.length;
    const MOD1 = 1000000007n, MOD2 = 1000000009n;
    const BASE1 = 1000003n, BASE2 = 1000033n;
    const h1: bigint[] = new Array(n + 1).fill(0n);
    const h2: bigint[] = new Array(n + 1).fill(0n);
    const p1: bigint[] = new Array(n + 1).fill(1n);
    const p2: bigint[] = new Array(n + 1).fill(1n);
    for (let i = 0; i < n; i++) {
        h1[i + 1] = (h1[i] * BASE1 + BigInt(nums[i])) % MOD1;
        h2[i + 1] = (h2[i] * BASE2 + BigInt(nums[i])) % MOD2;
        p1[i + 1] = p1[i] * BASE1 % MOD1;
        p2[i + 1] = p2[i] * BASE2 % MOD2;
    }
    const hasUnique = (len: number): boolean => {
        const cnt = new Map<string, number>();
        for (let i = 0; i + len <= n; i++) {
            let a = (h1[i + len] - h1[i] * p1[len] % MOD1 + MOD1) % MOD1;
            let b = (h2[i + len] - h2[i] * p2[len] % MOD2 + MOD2) % MOD2;
            const key = a.toString() + "," + b.toString();
            cnt.set(key, (cnt.get(key) || 0) + 1);
        }
        for (const v of cnt.values()) if (v === 1) return true;
        return false;
    };
    let lo = 1, hi = n, ans = n;
    while (lo <= hi) {
        const mid = lo + ((hi - lo) >> 1);
        if (hasUnique(mid)) {
            ans = mid;
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }
    return ans;
};