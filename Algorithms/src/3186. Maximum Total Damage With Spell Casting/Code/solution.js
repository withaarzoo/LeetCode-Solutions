/**
 * @param {number[]} power
 * @return {number}
 */
var maximumTotalDamage = function (power) {
  if (!power || power.length === 0) return 0;
  // Frequency map
  const freq = new Map();
  for (const v of power) freq.set(v, (freq.get(v) || 0) + 1);
  // Unique sorted values
  const vals = Array.from(freq.keys()).sort((a, b) => a - b);
  const m = vals.length;
  if (m === 0) return 0;
  const valueSum = new Array(m);
  for (let i = 0; i < m; ++i) valueSum[i] = vals[i] * freq.get(vals[i]);
  const dp = new Array(m).fill(0);
  dp[0] = valueSum[0];
  for (let i = 1; i < m; ++i) {
    const need = vals[i] - 3; // we can combine with indices having value <= need
    // binary search for last index <= need
    let lo = 0,
      hi = i - 1,
      j = -1;
    while (lo <= hi) {
      const mid = (lo + hi) >> 1;
      if (vals[mid] <= need) {
        j = mid;
        lo = mid + 1;
      } else hi = mid - 1;
    }
    const take = valueSum[i] + (j >= 0 ? dp[j] : 0);
    const skip = dp[i - 1];
    dp[i] = Math.max(skip, take);
  }
  return dp[m - 1];
};
