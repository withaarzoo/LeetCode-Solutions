/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var countPartitions = function (nums, k) {
  const MOD = 1_000_000_007;
  const n = nums.length;

  const dp = new Array(n + 1).fill(0);
  const pref = new Array(n + 1).fill(0);

  dp[0] = 1;
  pref[0] = 1;

  const maxdq = []; // store indices
  const mindq = [];
  let l = 0;

  for (let r = 0; r < n; r++) {
    const x = nums[r];

    // maintain decreasing deque for max
    while (maxdq.length > 0 && nums[maxdq[maxdq.length - 1]] <= x) {
      maxdq.pop();
    }
    maxdq.push(r);

    // maintain increasing deque for min
    while (mindq.length > 0 && nums[mindq[mindq.length - 1]] >= x) {
      mindq.pop();
    }
    mindq.push(r);

    // shrink left until valid
    while (
      maxdq.length > 0 &&
      mindq.length > 0 &&
      nums[maxdq[0]] - nums[mindq[0]] > k
    ) {
      if (maxdq[0] === l) maxdq.shift();
      if (mindq[0] === l) mindq.shift();
      l++;
    }

    const L = l;
    const i = r + 1;

    let ways = pref[i - 1];
    if (L > 0) ways -= pref[L - 1];
    ways %= MOD;
    if (ways < 0) ways += MOD;

    dp[i] = ways;
    pref[i] = (pref[i - 1] + dp[i]) % MOD;
  }

  return dp[n];
};
