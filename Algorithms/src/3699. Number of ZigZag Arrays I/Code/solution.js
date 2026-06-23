/**
 * @param {number} n
 * @param {number} l
 * @param {number} r
 * @return {number}
 */
var zigZagArrays = function (n, l, r) {
  const MOD = 1000000007;
  const m = r - l + 1;

  // Length 1 DP
  const dp = Array(m).fill(1);

  for (let len = 2; len <= n; len++) {
    // Reverse to reuse the same prefix-sum transition
    dp.reverse();

    let pref = 0;

    for (let i = 0; i < m; i++) {
      const old = dp[i];

      // New state = accumulated prefix
      dp[i] = pref;

      pref = (pref + old) % MOD;
    }
  }

  let ans = 0;

  for (const x of dp) {
    ans = (ans + x) % MOD;
  }

  // Both possible starting directions
  return (ans * 2) % MOD;
};
