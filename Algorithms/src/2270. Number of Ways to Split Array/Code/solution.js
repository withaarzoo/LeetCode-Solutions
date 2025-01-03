/**
 * @param {number[]} nums
 * @return {number}
 */
var waysToSplitArray = function (nums) {
  let totalSum = nums.reduce((a, b) => a + b, 0); // Total sum of the array
  let prefixSum = 0; // Prefix sum
  let count = 0; // Count of valid splits

  for (let i = 0; i < nums.length - 1; i++) {
    prefixSum += nums[i];
    let rightSum = totalSum - prefixSum;
    if (prefixSum >= rightSum) {
      count++;
    }
  }

  return count;
};
