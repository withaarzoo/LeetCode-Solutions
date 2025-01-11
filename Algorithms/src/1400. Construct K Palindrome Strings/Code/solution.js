/**
 * @param {string} s
 * @param {number} k
 * @return {boolean}
 */
var canConstruct = function (s, k) {
  if (k > s.length) return false; // More palindromes than characters
  let freq = Array(26).fill(0); // Frequency array for lowercase letters
  for (let char of s) {
    freq[char.charCodeAt(0) - 97]++;
  }
  let oddCount = 0;
  for (let count of freq) {
    if (count % 2 !== 0) {
      oddCount++;
    }
  }
  return oddCount <= k;
};
