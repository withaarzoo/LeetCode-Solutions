function sumCounts(nums: number[]): number {
    const MOD = 1000000007;
    const n = nums.length;
    const sum = new Array(4 * n).fill(0);
    const sumsq = new Array(4 * n).fill(0);
    const lazy = new Array(4 * n).fill(0);

    const push = (v: number, tl: number, tr: number) => {
        if (lazy[v] !== 0 && tl !== tr) {
            const mid = (tl + tr) >> 1;
            const llen = mid - tl + 1;
            const rlen = tr - mid;
            const add = lazy[v];

            sumsq[v << 1] = (sumsq[v << 1] + 2 * add * sum[v << 1] + add * add % MOD * llen) % MOD;
            sum[v << 1] = (sum[v << 1] + add * llen) % MOD;
            lazy[v << 1] = (lazy[v << 1] + add) % MOD;

            sumsq[v << 1 | 1] = (sumsq[v << 1 | 1] + 2 * add * sum[v << 1 | 1] + add * add % MOD * rlen) % MOD;
            sum[v << 1 | 1] = (sum[v << 1 | 1] + add * rlen) % MOD;
            lazy[v << 1 | 1] = (lazy[v << 1 | 1] + add) % MOD;

            lazy[v] = 0;
        }
    };

    const update = (v: number, tl: number, tr: number, l: number, r: number, add: number) => {
        if (l > r) return;
        if (l === tl && r === tr) {
            const len = tr - tl + 1;
            sumsq[v] = (sumsq[v] + 2 * add * sum[v] + add * add % MOD * len) % MOD;
            sum[v] = (sum[v] + add * len) % MOD;
            lazy[v] = (lazy[v] + add) % MOD;
            return;
        }
        push(v, tl, tr);
        const mid = (tl + tr) >> 1;
        update(v << 1, tl, mid, l, Math.min(r, mid), add);
        update(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r, add);
        sum[v] = (sum[v << 1] + sum[v << 1 | 1]) % MOD;
        sumsq[v] = (sumsq[v << 1] + sumsq[v << 1 | 1]) % MOD;
    };

    const query = (v: number, tl: number, tr: number, l: number, r: number): number => {
        if (l > r) return 0;
        if (l === tl && r === tr) return sumsq[v];
        push(v, tl, tr);
        const mid = (tl + tr) >> 1;
        return (query(v << 1, tl, mid, l, Math.min(r, mid)) +
                query(v << 1 | 1, mid + 1, tr, Math.max(l, mid + 1), r)) % MOD;
    };

    const last = new Array(100001).fill(-1);
    let ans = 0;
    for (let r = 0; r < n; r++) {
        const prv = last[nums[r]];
        if (prv + 1 <= r - 1) update(1, 0, n - 1, prv + 1, r - 1, 1);
        update(1, 0, n - 1, r, r, 1);
        ans = (ans + query(1, 0, n - 1, 0, r)) % MOD;
        last[nums[r]] = r;
    }
    return ans;
}