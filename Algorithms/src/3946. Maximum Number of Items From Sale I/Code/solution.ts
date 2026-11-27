function maximumSaleItems(items: number[][], budget: number): number {
    const mult = new Array(1501).fill(0);
    for (const it of items) {
        const fac = it[0];
        for (let f = 1; f <= 1500; ++f) {
            if (fac % f === 0) ++mult[f];
        }
    }
    let dp = new Array(budget + 1).fill(0);
    for (const it of items) {
        const f = it[0], p = it[1];
        const ndp = dp.slice();
        for (let b = budget; b >= p; --b) {
            ndp[b] = Math.max(ndp[b], dp[b - p] + mult[f]);
        }
        for (let b = p; b <= budget; ++b) {
            ndp[b] = Math.max(ndp[b], ndp[b - p] + 1);
        }
        dp = ndp;
    }
    return Math.max(...dp);
};