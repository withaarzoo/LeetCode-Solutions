var maximumProfit = function (prices, k) {
  const n = prices.length;
  let prev = Array(n).fill(0);
  let curr = Array(n).fill(0);

  for (let t = 1; t <= k; t++) {
    let bestLong = -prices[0];
    let bestShort = prices[0];
    curr[0] = 0;

    for (let i = 1; i < n; i++) {
      curr[i] = Math.max(
        curr[i - 1],
        prices[i] + bestLong,
        -prices[i] + bestShort
      );

      bestLong = Math.max(bestLong, prev[i - 1] - prices[i]);
      bestShort = Math.max(bestShort, prev[i - 1] + prices[i]);
    }
    [prev, curr] = [curr, prev];
  }
  return prev[n - 1];
};
