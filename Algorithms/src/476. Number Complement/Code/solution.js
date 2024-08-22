/**
 * This function finds the complement of a given number `num`.
 * The complement is achieved by flipping all the bits in the binary representation of the number.
 *
 * @param {number} num - The input number whose complement is to be found.
 * @return {number} - The complement of the input number.
 */
var findComplement = function (num) {
  // Initialize a mask variable to 0. This mask will eventually have all bits set to 1
  // that are within the range of `num`.
  let mask = 0;

  // Create a temporary variable to hold the value of `num` for manipulation.
  let temp = num;

  // Loop to create the mask with all bits set to 1 for the length of `num`.
  // The loop runs until `temp` becomes 0.
  while (temp !== 0) {
    // Shift the mask to the left by 1 position and then set the least significant bit (LSB) to 1.
    // This gradually builds a mask of the same length as the binary representation of `num`.
    mask = (mask << 1) | 1;

    // Right shift `temp` by 1 to move to the next bit.
    temp >>= 1;
  }

  // The XOR operation between `num` and `mask` flips all the bits of `num`
  // where the mask has bits set to 1, effectively calculating the complement.
  return num ^ mask;
};
