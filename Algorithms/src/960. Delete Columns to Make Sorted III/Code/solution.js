var minDeletionSize = function (strs) {
  const n = strs.length;
  const m = strs[0].length;

  const dp = Array(m).fill(1);

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < i; j++) {
      let valid = true;

      for (let r = 0; r < n; r++) {
        if (strs[r][j] > strs[r][i]) {
          valid = false;
          break;
        }
      }

      if (valid) {
        dp[i] = Math.max(dp[i], dp[j] + 1);
      }
    }
  }

  const keep = Math.max(...dp);
  return m - keep;
};
