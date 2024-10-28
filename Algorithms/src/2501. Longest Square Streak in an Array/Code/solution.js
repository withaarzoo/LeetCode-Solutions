/**
 * @param {number[]} nums
 * @return {number}
 */
var longestSquareStreak = function (nums) {
  nums.sort((a, b) => a - b);
  const numSet = new Set(nums);
  let maxLength = -1;

  for (let num of nums) {
    let length = 0;
    let current = num;

    while (numSet.has(current)) {
      length++;
      current *= current;
      if (current > 1e9) break;
    }

    if (length >= 2) {
      maxLength = Math.max(maxLength, length);
    }
  }
  return maxLength;
};
