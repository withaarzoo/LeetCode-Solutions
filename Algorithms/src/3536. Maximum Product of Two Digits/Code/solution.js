/**
 * @param {number} n
 * @return {number}
 */
var maxProduct = function (n) {
  // Store the largest digit
  let first = 0;

  // Store the second largest digit
  let second = 0;

  // Process every digit
  while (n > 0) {
    // Extract the last digit
    const digit = n % 10;

    // Update the two largest digits
    if (digit >= first) {
      second = first;
      first = digit;
    } else if (digit > second) {
      second = digit;
    }

    // Remove the last digit
    n = Math.floor(n / 10);
  }

  // Return the maximum product
  return first * second;
};
