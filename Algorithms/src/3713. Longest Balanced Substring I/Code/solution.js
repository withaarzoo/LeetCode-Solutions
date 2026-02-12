/**
 * @param {string} s
 * @return {number}
 */
var longestBalanced = function (s) {
  const n = s.length;
  let ans = 0;

  for (let i = 0; i < n; i++) {
    let freq = new Array(26).fill(0);
    let distinct = 0;
    let maxFreq = 0;

    for (let j = i; j < n; j++) {
      let idx = s.charCodeAt(j) - 97;

      if (freq[idx] === 0) distinct++;

      freq[idx]++;
      maxFreq = Math.max(maxFreq, freq[idx]);

      let length = j - i + 1;

      if (length === distinct * maxFreq) ans = Math.max(ans, length);
    }
  }

  return ans;
};
