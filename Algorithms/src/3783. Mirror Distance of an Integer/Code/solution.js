/**
 * @param {number} n
 * @return {number}
 */
var mirrorDistance = function (n) {
  let rev = 0;
  let temp = n;

  // Reverse the digits of n
  while (temp > 0) {
    let digit = temp % 10; // Get last digit
    rev = rev * 10 + digit; // Add digit to reversed number
    temp = Math.floor(temp / 10); // Remove last digit
  }

  // Return absolute difference
  return Math.abs(n - rev);
};
