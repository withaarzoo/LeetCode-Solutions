/**
 * @param {number[]} nums
 * @return {number}
 */
var maxFrequencyElements = function (nums) {
  const freq = new Map();
  // Build frequency map
  for (const x of nums) freq.set(x, (freq.get(x) || 0) + 1);

  // Find max frequency
  let maxFreq = 0;
  for (const v of freq.values()) if (v > maxFreq) maxFreq = v;

  // Sum counts of elements whose frequency == maxFreq
  let result = 0;
  for (const v of freq.values()) if (v === maxFreq) result += v;
  return result;
};
