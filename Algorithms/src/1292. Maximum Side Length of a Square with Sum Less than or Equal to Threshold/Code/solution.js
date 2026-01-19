var maxSideLength = function (mat, threshold) {
  const m = mat.length,
    n = mat[0].length;
  const pre = Array.from({ length: m + 1 }, () => Array(n + 1).fill(0));

  // Build prefix sum
  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      pre[i][j] =
        mat[i - 1][j - 1] + pre[i - 1][j] + pre[i][j - 1] - pre[i - 1][j - 1];
    }
  }

  let left = 0,
    right = Math.min(m, n),
    ans = 0;

  while (left <= right) {
    const mid = Math.floor((left + right) / 2);
    let found = false;

    for (let i = mid; i <= m && !found; i++) {
      for (let j = mid; j <= n; j++) {
        const sum =
          pre[i][j] - pre[i - mid][j] - pre[i][j - mid] + pre[i - mid][j - mid];

        if (sum <= threshold) {
          found = true;
          break;
        }
      }
    }

    if (found) {
      ans = mid;
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return ans;
};
