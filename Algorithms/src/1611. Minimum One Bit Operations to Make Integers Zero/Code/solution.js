/**
 * @param {number} n
 * @return {number}
 */
var minimumOneBitOperations = function (n) {
  // In JS, numbers are 64-bit floats but bitwise ops use 32-bit signed ints.
  // Constraint fits in 32-bit, so this is safe.
  let ans = 0;
  while (n !== 0) {
    ans ^= n;
    n >>= 1;
  }
  return ans;
};
