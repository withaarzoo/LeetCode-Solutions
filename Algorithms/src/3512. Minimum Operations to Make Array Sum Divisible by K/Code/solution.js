/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var minOperations = function (nums, k) {
  let sum = 0;

  // Calculate total sum of the array
  for (const x of nums) {
    sum += x;
  }

  // Minimum operations is sum % k
  const remainder = sum % k;

  return remainder;
};
