var maxDotProduct = function (nums1, nums2) {
  const n = nums1.length;
  const m = nums2.length;

  const dp = Array.from({ length: n + 1 }, () => Array(m + 1).fill(-Infinity));

  for (let i = n - 1; i >= 0; i--) {
    for (let j = m - 1; j >= 0; j--) {
      const product = nums1[i] * nums2[j];

      let takeBoth = product;
      if (dp[i + 1][j + 1] !== -Infinity) {
        takeBoth = Math.max(takeBoth, product + dp[i + 1][j + 1]);
      }

      dp[i][j] = Math.max(takeBoth, dp[i + 1][j], dp[i][j + 1]);
    }
  }

  return dp[0][0];
};
