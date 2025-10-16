/**
 * @param {number[]} nums
 * @param {number} value
 * @return {number}
 */
var findSmallestInteger = function (nums, value) {
  const freq = new Array(value).fill(0);
  for (const a of nums) {
    // normalize modulo for negatives
    let r = a % value;
    if (r < 0) r += value;
    freq[r]++;
  }
  let x = 0;
  while (true) {
    const r = x % value;
    if (freq[r] === 0) return x;
    freq[r]--;
    x++;
  }
};
