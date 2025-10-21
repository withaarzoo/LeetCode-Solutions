/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} numOperations
 * @return {number}
 */
const maxFrequency = (nums, k, numOperations) => {
  if (nums.length === 0) return 0;
  let mx = Math.max(...nums);
  let size = mx + k + 2;
  const count = new Array(size).fill(0);

  for (const v of nums) count[v]++;

  // prefix sums
  for (let i = 1; i < size; ++i) count[i] += count[i - 1];

  let ans = 0;
  for (let t = 0; t < size; ++t) {
    const L = Math.max(0, t - k);
    const R = Math.min(size - 1, t + k);
    const total = count[R] - (L > 0 ? count[L - 1] : 0);
    const freq_t = t > 0 ? count[t] - count[t - 1] : count[t];
    const canConvert = total - freq_t;
    const take = Math.min(numOperations, canConvert);
    ans = Math.max(ans, freq_t + take);
  }
  return ans;
};
