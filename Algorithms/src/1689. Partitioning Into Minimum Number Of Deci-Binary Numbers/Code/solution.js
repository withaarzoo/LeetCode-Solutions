/**
 * @param {string} n
 * @return {number}
 */
var minPartitions = function (n) {
  let maxDigit = 0; // Store maximum digit

  for (let i = 0; i < n.length; i++) {
    let digit = n[i] - "0"; // Convert to number

    if (digit > maxDigit) {
      maxDigit = digit;
    }

    // If 9 found, no need to continue
    if (maxDigit === 9) {
      break;
    }
  }

  return maxDigit;
};
