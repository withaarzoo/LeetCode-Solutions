function maximumSaleItems(items: number[][], budget: number): number {
    const n = items.length;
    const cnt = new Array(n + 1).fill(0);
    for (const it of items) cnt[it[0]]++;
    const mult = new Array(n + 1).fill(0);
    for (let f = 1; f <= n; f++) {
        for (let m = f; m <= n; m += f) mult[f] += cnt[m];
    }
    let minP = Infinity;
    const cand: number[][] = [];
    for (const it of items) {
        const f = it[0], p = it[1];
        minP = Math.min(minP, p);
        const g = mult[f] - 1;
        if (g > 0) cand.push([p, g]);
    }
    cand.sort((a, b) => a[0] - b[0]);
    let spent = 0, total = 0, B = budget;
    for (const [p, g] of cand) {
        if (p > 2 * minP) break;
        const can = Math.min(g, Math.floor((B - spent) / p));
        total += can * 2;
        spent += can * p;
    }
    total += Math.floor((B - spent) / minP);
    return total;
};