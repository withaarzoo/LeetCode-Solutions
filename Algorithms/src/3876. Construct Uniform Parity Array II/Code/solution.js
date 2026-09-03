/**
 * @param {number[]} nums1
 * @return {boolean}
 */
var uniformArray = function (nums1) {
  // Use Infinity to represent that no odd or even value has been found yet.
  let minOdd = Infinity;
  let minEven = Infinity;

  // Scan every value once to find the minimum value of each parity.
  for (const x of nums1) {
    if (x % 2 === 0) {
      // The smallest even value is the hardest even value to convert.
      minEven = Math.min(minEven, x);
    } else {
      // The smallest odd value is the best value available for subtraction.
      minOdd = Math.min(minOdd, x);
    }
  }

  // If there are no odd values, all numbers are already even.
  if (minOdd === Infinity) {
    return true;
  }

  // Every even value must be larger than minOdd so it can subtract minOdd.
  return minOdd < minEven;
};
