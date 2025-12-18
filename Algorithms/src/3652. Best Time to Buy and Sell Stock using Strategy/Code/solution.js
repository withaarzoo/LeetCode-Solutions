var maxProfit = function (prices, strategy, k) {
  const n = prices.length;

  let base = 0;
  for (let i = 0; i < n; i++) {
    base += strategy[i] * prices[i];
  }

  const prefPrice = Array(n + 1).fill(0);
  const prefProfit = Array(n + 1).fill(0);

  for (let i = 0; i < n; i++) {
    prefPrice[i + 1] = prefPrice[i] + prices[i];
    prefProfit[i + 1] = prefProfit[i] + strategy[i] * prices[i];
  }

  let bestDelta = 0;
  const half = k / 2;

  for (let l = 0; l + k <= n; l++) {
    const m = l + half;
    const r = l + k;

    const oldProfit = prefProfit[r] - prefProfit[l];
    const newProfit = prefPrice[r] - prefPrice[m];

    bestDelta = Math.max(bestDelta, newProfit - oldProfit);
  }

  return base + bestDelta;
};
