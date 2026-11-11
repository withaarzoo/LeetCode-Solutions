/**
 * @param {number} cost1
 * @param {number} cost2
 * @param {number} costBoth
 * @param {number} need1
 * @param {number} need2
 * @return {number}
 */
var minimumCost = function(cost1, cost2, costBoth, need1, need2) {
    let c1 = cost1, c2 = cost2, cb = costBoth;
    let n1 = need1, n2 = need2;
    let pairs = Math.min(n1, n2);
    let costPairs = pairs * Math.min(cb, c1 + c2);
    let rem1 = n1 - pairs;
    let rem2 = n2 - pairs;
    let costRem = rem1 * Math.min(c1, cb) + rem2 * Math.min(c2, cb);
    return costPairs + costRem;
};