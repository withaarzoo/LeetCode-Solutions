/**
 * @param {number[]} nums
 * @return {number[]}
 */
var leftRightDifference = function (nums) {
  const n = nums.length;

  // Calculate total array sum
  let rightSum = 0;
  for (const num of nums) {
    rightSum += num;
  }

  // Sum of elements on the left side
  let leftSum = 0;

  // Result array
  const ans = new Array(n);

  for (let i = 0; i < n; i++) {
    // Remove current element from right side sum
    rightSum -= nums[i];

    // Store absolute difference
    ans[i] = Math.abs(leftSum - rightSum);

    // Add current element to left side sum
    leftSum += nums[i];
  }

  return ans;
};
