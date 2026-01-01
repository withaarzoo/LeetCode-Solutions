/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
  // Start from the last digit
  for (let i = digits.length - 1; i >= 0; i--) {
    digits[i]++;

    if (digits[i] < 10) {
      // No carry needed
      return digits;
    }

    digits[i] = 0; // Carry forward
  }

  // All digits were 9
  digits.unshift(1);
  return digits;
};
