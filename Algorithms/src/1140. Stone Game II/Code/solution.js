/**
 * @param {number[]} piles
 * @return {number}
 */
var stoneGameII = function (piles) {
  const n = piles.length;
  const suffix = new Array(n + 1).fill(0);

  for (let i = n - 1; i >= 0; i--) {
    suffix[i] = suffix[i + 1] + piles[i];
  }

  const dp = Array.from({ length: n }, () => new Array(n + 1).fill(-1));

  const solve = (i, M) => {
    if (i === n) {
      return 0;
    }

    if (dp[i][M] !== -1) {
      return dp[i][M];
    }

    let best = 0;

    for (let X = 1; X <= 2 * M && i + X <= n; X++) {
      const nextM = Math.max(M, X);
      const current = suffix[i] - solve(i + X, nextM);
      best = Math.max(best, current);
    }

    dp[i][M] = best;
    return best;
  };

  return solve(0, 1);
};
