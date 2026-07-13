/**
 * @param {number} low
 * @param {number} high
 * @return {number[]}
 */
var sequentialDigits = function (low, high) {
  // String containing all consecutive digits
  const digits = "123456789";

  // Result array
  const ans = [];

  // Number of digits in low and high
  const minLen = low.toString().length;
  const maxLen = high.toString().length;

  // Try every possible length
  for (let len = minLen; len <= maxLen; len++) {
    // Generate every substring of current length
    for (let start = 0; start + len <= 9; start++) {
      // Convert substring into a number
      const num = Number(digits.substring(start, start + len));

      // Keep only numbers inside the range
      if (num >= low && num <= high) {
        ans.push(num);
      }
    }
  }

  return ans;
};
