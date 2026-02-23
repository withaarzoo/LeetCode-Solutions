/**
 * @param {string} s
 * @param {number} k
 * @return {boolean}
 */
var hasAllCodes = function (s, k) {
  const n = s.length;

  if (n < k) return false;

  const total = 1 << k;
  if (n - k + 1 < total) return false;

  const seen = new Array(total).fill(false);
  const mask = total - 1;

  let curr = 0;
  let count = 0;

  // First window
  for (let i = 0; i < k; i++) {
    curr = (curr << 1) | (s[i] - "0");
  }

  if (!seen[curr]) {
    seen[curr] = true;
    count++;
  }

  // Sliding window
  for (let i = k; i < n; i++) {
    curr = ((curr << 1) & mask) | (s[i] - "0");

    if (!seen[curr]) {
      seen[curr] = true;
      count++;
      if (count === total) return true;
    }
  }

  return count === total;
};
