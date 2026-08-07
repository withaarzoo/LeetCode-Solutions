/**
 * @param {string} num
 * @param {number} t
 * @return {string}
 */
var smallestNumber = function (num, t) {
  let req2 = 0,
    req3 = 0,
    req5 = 0,
    req7 = 0;
  let temp = t;
  // Map initial factorization needed from the number
  while (temp % 2 === 0) {
    temp /= 2;
    req2++;
  }
  while (temp % 3 === 0) {
    temp /= 3;
    req3++;
  }
  while (temp % 5 === 0) {
    temp /= 5;
    req5++;
  }
  while (temp % 7 === 0) {
    temp /= 7;
    req7++;
  }
  // Abort if target cannot be formed using digits 1-9
  if (temp > 1) return "-1";

  const dp = Array.from({ length: 60 }, () => Array(40).fill(1e9));
  dp[0][0] = 0;

  // Test base digit additions to build forward map of dependencies
  const trans = [
    [1, 0],
    [0, 1],
    [2, 0],
    [1, 1],
    [3, 0],
    [0, 2],
  ];
  for (let i = 0; i < 60; ++i) {
    for (let j = 0; j < 40; ++j) {
      if (dp[i][j] === 1e9) continue;
      for (const tr of trans) {
        const ni = Math.min(59, i + tr[0]);
        const nj = Math.min(39, j + tr[1]);
        dp[ni][nj] = Math.min(dp[ni][nj], dp[i][j] + 1);
      }
    }
  }

  // Resolve states iteratively to handle 'at least' thresholding properly
  for (let i = 59; i >= 0; --i) {
    for (let j = 39; j >= 0; --j) {
      if (i < 59) dp[i][j] = Math.min(dp[i][j], dp[i + 1][j]);
      if (j < 39) dp[i][j] = Math.min(dp[i][j], dp[i][j + 1]);
    }
  }

  const F2 = [0, 0, 1, 0, 2, 0, 1, 0, 3, 0];
  const F3 = [0, 0, 0, 1, 0, 0, 1, 0, 0, 2];
  const F5 = [0, 0, 0, 0, 0, 1, 0, 0, 0, 0];
  const F7 = [0, 0, 0, 0, 0, 0, 0, 1, 0, 0];

  const n = num.length;
  let hasZero = false;
  let firstZero = n;
  for (let i = 0; i < n; ++i) {
    if (num[i] === "0") {
      hasZero = true;
      firstZero = i;
      break;
    }
  }

  // Verify whether zero-free original num evaluates as sufficient
  if (!hasZero) {
    let r2 = req2,
      r3 = req3,
      r5 = req5,
      r7 = req7;
    for (let char of num) {
      let d = Number(char);
      r2 = Math.max(0, r2 - F2[d]);
      r3 = Math.max(0, r3 - F3[d]);
      r5 = Math.max(0, r5 - F5[d]);
      r7 = Math.max(0, r7 - F7[d]);
    }
    if (r2 === 0 && r3 === 0 && r5 === 0 && r7 === 0) return num;
  }

  // Capture intact factor counts strictly ahead of zeros
  const limit = Math.min(n - 1, firstZero);
  let p2 = 0,
    p3 = 0,
    p5 = 0,
    p7 = 0;
  for (let i = 0; i < limit; ++i) {
    let d = Number(num[i]);
    p2 += F2[d];
    p3 += F3[d];
    p5 += F5[d];
    p7 += F7[d];
  }

  // Rewind loop finding longest matching prefix segment safely incrementable
  for (let i = limit; i >= 0; --i) {
    let startD = Number(num[i]) + 1;
    for (let d = startD; d <= 9; ++d) {
      let n2 = Math.max(0, req2 - p2 - F2[d]);
      let n3 = Math.max(0, req3 - p3 - F3[d]);
      let n5 = Math.max(0, req5 - p5 - F5[d]);
      let n7 = Math.max(0, req7 - p7 - F7[d]);
      let L = n - 1 - i;

      // Assemble output string cleanly checking remaining capacity DP maps
      if (n7 + n5 + dp[n2][n3] <= L) {
        let ans = num.substring(0, i) + d.toString();
        let rem2 = n2,
          rem3 = n3,
          rem5 = n5,
          rem7 = n7;
        for (let pos = 0; pos < L; ++pos) {
          for (let x = 1; x <= 9; ++x) {
            let nn2 = Math.max(0, rem2 - F2[x]);
            let nn3 = Math.max(0, rem3 - F3[x]);
            let nn5 = Math.max(0, rem5 - F5[x]);
            let nn7 = Math.max(0, rem7 - F7[x]);
            if (nn7 + nn5 + dp[nn2][nn3] <= L - 1 - pos) {
              ans += x.toString();
              rem2 = nn2;
              rem3 = nn3;
              rem5 = nn5;
              rem7 = nn7;
              break;
            }
          }
        }
        return ans;
      }
    }
    if (i > 0) {
      let d = Number(num[i - 1]);
      p2 -= F2[d];
      p3 -= F3[d];
      p5 -= F5[d];
      p7 -= F7[d];
    }
  }

  // Force build longer string assuming original length bounds are exhausted
  let minLenNeeded = req7 + req5 + dp[req2][req3];
  let M = Math.max(n + 1, minLenNeeded);
  let ans = "";
  let rem2 = req2,
    rem3 = req3,
    rem5 = req5,
    rem7 = req7;

  for (let pos = 0; pos < M; ++pos) {
    for (let x = 1; x <= 9; ++x) {
      let nn2 = Math.max(0, rem2 - F2[x]);
      let nn3 = Math.max(0, rem3 - F3[x]);
      let nn5 = Math.max(0, rem5 - F5[x]);
      let nn7 = Math.max(0, rem7 - F7[x]);
      if (nn7 + nn5 + dp[nn2][nn3] <= M - 1 - pos) {
        ans += x.toString();
        rem2 = nn2;
        rem3 = nn3;
        rem5 = nn5;
        rem7 = nn7;
        break;
      }
    }
  }
  return ans;
};
