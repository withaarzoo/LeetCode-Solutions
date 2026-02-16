/**
 * @param {number} n
 * @return {number}
 */
var reverseBits = function (n) {
  let result = 0;

  for (let i = 0; i < 32; i++) {
    // Shift result left
    result = result << 1;

    // Add last bit of n
    result = result | (n & 1);

    // Unsigned right shift
    n = n >>> 1;
  }

  // Ensure unsigned 32-bit result
  return result >>> 0;
};
