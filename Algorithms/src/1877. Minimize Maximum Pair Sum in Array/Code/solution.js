/**
 * @param {number[]} nums
 * @return {number}
 */
var minPairSum = function (nums) {
  // Step 1: Sort the array
  nums.sort((a, b) => a - b);

  let left = 0;
  let right = nums.length - 1;
  let maxPairSum = 0;

  // Step 2: Pair smallest with largest
  while (left < right) {
    const pairSum = nums[left] + nums[right];
    maxPairSum = Math.max(maxPairSum, pairSum);
    left++;
    right--;
  }

  return maxPairSum;
};
