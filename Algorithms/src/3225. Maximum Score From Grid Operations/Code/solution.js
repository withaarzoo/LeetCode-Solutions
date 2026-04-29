/**
 * @param {number[][]} grid
 * @return {number}
 */
var maximumScore = function (grid) {
  const n = grid.length;
  if (n === 1) return 0;

  // pref[c][k] = sum of first k cells in column c
  const pref = Array.from({ length: n }, () => Array(n + 1).fill(0));
  for (let c = 0; c < n; c++) {
    for (let r = 0; r < n; r++) {
      pref[c][r + 1] = pref[c][r] + grid[r][c];
    }
  }

  const NEG = -1e30;

  // dp[a][b] = best score after processing up to current column,
  // with previous height = a and current height = b.
  let dp = Array.from({ length: n + 1 }, () => Array(n + 1).fill(NEG));

  // Initialize using the first column.
  for (let a = 0; a <= n; a++) {
    for (let b = 0; b <= n; b++) {
      dp[a][b] = Math.max(0, pref[0][b] - pref[0][a]);
    }
  }

  for (let col = 1; col < n; col++) {
    const ndp = Array.from({ length: n + 1 }, () => Array(n + 1).fill(NEG));

    for (let mid = 0; mid <= n; mid++) {
      const q = Array(n + 1).fill(0);
      for (let x = 0; x <= n; x++) {
        q[x] = Math.max(0, pref[col][x] - pref[col][mid]);
      }

      const prefixBest = Array(n + 1).fill(NEG);
      prefixBest[0] = dp[0][mid];
      for (let a = 1; a <= n; a++) {
        prefixBest[a] = Math.max(prefixBest[a - 1], dp[a][mid]);
      }

      const suffixBest = Array(n + 2).fill(NEG);
      suffixBest[n] = dp[n][mid] + q[n];
      for (let a = n - 1; a >= 0; a--) {
        suffixBest[a] = Math.max(suffixBest[a + 1], dp[a][mid] + q[a]);
      }

      const limit = col === n - 1 ? 0 : n;
      for (let nxt = 0; nxt <= limit; nxt++) {
        let best = NEG;

        if (prefixBest[nxt] !== NEG) {
          best = Math.max(best, prefixBest[nxt] + q[nxt]);
        }
        if (suffixBest[nxt + 1] !== NEG) {
          best = Math.max(best, suffixBest[nxt + 1]);
        }

        ndp[mid][nxt] = Math.max(ndp[mid][nxt], best);
      }
    }

    dp = ndp;
  }

  let ans = 0;
  for (let a = 0; a <= n; a++) {
    for (let b = 0; b <= n; b++) {
      ans = Math.max(ans, dp[a][b]);
    }
  }
  return ans;
};
