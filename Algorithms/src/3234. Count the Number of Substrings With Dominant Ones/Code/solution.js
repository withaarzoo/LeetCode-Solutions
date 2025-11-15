/**
 * @param {string} s
 * @return {number}
 */
var numberOfSubstrings = function (s) {
  const n = s.length;
  let ans = 0;
  // all-ones substrings
  let run = 0;
  for (let i = 0; i < n; ++i) {
    if (s[i] === "1") run++;
    else {
      ans += (run * (run + 1)) / 2;
      run = 0;
    }
  }
  ans += (run * (run + 1)) / 2;

  // positions of zeros
  const zeroPos = [];
  for (let i = 0; i < n; ++i) if (s[i] === "0") zeroPos.push(i);
  const m = zeroPos.length;
  if (m === 0) return ans;

  const K = Math.floor(Math.sqrt(n));
  for (let k = 1; k <= K && k <= m; ++k) {
    for (let i = 0; i + k - 1 < m; ++i) {
      const leftPrev = i === 0 ? -1 : zeroPos[i - 1];
      const rightNext = i + k - 1 === m - 1 ? n : zeroPos[i + k];
      const leftOnes = zeroPos[i] - leftPrev - 1;
      const rightOnes = rightNext - zeroPos[i + k - 1] - 1;
      const baseLen = zeroPos[i + k - 1] - zeroPos[i] + 1;
      const needLen = k * k + k;
      const t = needLen - baseLen;
      const totalPairs = (leftOnes + 1) * (rightOnes + 1);
      if (t <= 0) {
        ans += totalPairs;
        continue;
      }

      let pairs_lt = 0;
      let s0 = t - 1;
      if (s0 >= 0) {
        let L = leftOnes,
          R = rightOnes;
        let x_max = Math.min(L, s0);
        if (x_max >= 0) {
          let x0 = Math.max(0, s0 - R);
          if (x0 > x_max) pairs_lt = (x_max + 1) * (R + 1);
          else {
            let part1 = x0 * (R + 1);
            let n2 = x_max - x0 + 1;
            let sum_x = ((x0 + x_max) * n2) / 2;
            let part2 = n2 * (s0 + 1) - sum_x;
            pairs_lt = part1 + part2;
          }
        } else pairs_lt = 0;
      } else pairs_lt = 0;

      const valid = totalPairs - pairs_lt;
      if (valid > 0) ans += valid;
    }
  }

  return ans;
};
