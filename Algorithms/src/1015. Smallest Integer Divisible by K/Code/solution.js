/**
 * @param {number} k
 * @return {number}
 */
var smallestRepunitDivByK = function (k) {
  // If k is multiple of 2 or 5, no all-ones number is divisible by k.
  if (k % 2 === 0 || k % 5 === 0) return -1;

  let rem = 0; // current remainder

  for (let length = 1; length <= k; length++) {
    // Compute remainder of next repunit
    rem = (rem * 10 + 1) % k;
    if (rem === 0) {
      return length;
    }
  }

  return -1;
};
