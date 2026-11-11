function minimumCost(cost1: number, cost2: number, costBoth: number, need1: number, need2: number): number {
    let c1 = cost1, c2 = cost2, cb = costBoth;
    let n1 = need1, n2 = need2;
    let pairs = Math.min(n1, n2);
    let costPairs = pairs * Math.min(cb, c1 + c2);
    let rem1 = n1 - pairs;
    let rem2 = n2 - pairs;
    let costRem = rem1 * Math.min(c1, cb) + rem2 * Math.min(c2, cb);
    return costPairs + costRem;
};