/**
 * @param {number[]} nums
 * @return {number}
 */
var countPartitions = function (nums) {
  let total = 0;
  // Compute the total sum of the array
  for (const x of nums) {
    total += x;
  }

  // If total sum is odd, no valid partition
  if (total % 2 !== 0) return 0;

  // If total is even, every position between elements is a valid partition
  const n = nums.length;
  return n - 1;
};
