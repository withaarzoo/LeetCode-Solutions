/**
 * @param {number} n
 * @return {number}
 */
var sumAndMultiply = function (n) {
  // This stores the number formed by all non-zero digits.
  let x = 0;

  // This stores the sum of all non-zero digits.
  let sum = 0;

  // Find the highest place value to read digits left to right.
  let divisor = 1;
  while (Math.floor(n / divisor) >= 10) {
    divisor *= 10;
  }

  // Process every digit from left to right.
  while (divisor > 0) {
    // Extract the digit at the current place value.
    const digit = Math.floor(n / divisor);

    // Remove the current digit from n.
    n %= divisor;

    // Ignore zero digits completely.
    if (digit !== 0) {
      // Append the current digit to x.
      x = x * 10 + digit;

      // Add the current digit to the sum.
      sum += digit;
    }

    // Move to the next smaller place value.
    divisor = Math.floor(divisor / 10);
  }

  // Return the concatenated number multiplied by its digit sum.
  return x * sum;
};
