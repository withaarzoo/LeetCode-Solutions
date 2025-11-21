/**
 * @param {string} s
 * @return {number}
 */
var countPalindromicSubsequence = function (s) {
  const n = s.length;
  const A = 26;
  const first = new Array(A).fill(Number.MAX_SAFE_INTEGER);
  const last = new Array(A).fill(-1);
  // record first and last occurrence
  for (let i = 0; i < n; ++i) {
    const c = s.charCodeAt(i) - 97;
    first[c] = Math.min(first[c], i);
    last[c] = Math.max(last[c], i);
  }

  let ans = 0;
  for (let c = 0; c < A; ++c) {
    if (first[c] < last[c]) {
      const seen = new Array(A).fill(false);
      for (let i = first[c] + 1; i < last[c]; ++i) {
        seen[s.charCodeAt(i) - 97] = true;
      }
      for (let j = 0; j < A; ++j) if (seen[j]) ans++;
    }
  }
  return ans;
};
