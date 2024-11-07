/**
 * @param {number[]} candidates
 * @return {number}
 */
var largestCombination = function (candidates) {
  let bitCount = Array(31).fill(0); // Array to count '1's at each bit position

  // Count '1's in each bit position across all numbers
  for (let num of candidates) {
    for (let i = 0; i < 31; ++i) {
      if ((num & (1 << i)) !== 0) {
        bitCount[i]++;
      }
    }
  }

  // Find the maximum count in any bit position
  return Math.max(...bitCount);
};
