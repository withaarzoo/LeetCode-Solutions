/**
 * @param {number} n
 * @return {number}
 */
var bitwiseComplement = function (n) {
  // Edge case
  if (n === 0) return 1;

  let mask = 0;

  // Build mask with all bits = 1
  while (mask < n) {
    mask = (mask << 1) | 1;
  }

  // XOR flips bits
  return mask ^ n;
};
