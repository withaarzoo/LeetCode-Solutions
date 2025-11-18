/**
 * @param {number[]} bits
 * @return {boolean}
 */
var isOneBitCharacter = function (bits) {
  const n = bits.length;
  let i = 0;
  // walk until the last bit (we stop when i >= n-1)
  while (i < n - 1) {
    if (bits[i] === 1) {
      // '1' is always start of two-bit char => skip two
      i += 2;
    } else {
      // '0' is one-bit => skip one
      i += 1;
    }
  }
  // if we landed on the last index, that last bit is one-bit
  return i === n - 1;
};
