/**
 * @param {number[]} energy
 * @param {number} k
 * @return {number}
 */
var maximumEnergy = function (energy, k) {
  const n = energy.length;
  let ans = -Infinity;
  for (let r = 0; r < k; ++r) {
    let cur = 0;
    // last index in this residue class
    const last = r + Math.floor((n - 1 - r) / k) * k;
    for (let i = last; i >= r; i -= k) {
      cur += energy[i]; // accumulate suffix sum
      if (cur > ans) ans = cur;
    }
  }
  return ans;
};
