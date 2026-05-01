/**
 * @param {number[]} nums
 * @return {number}
 */
var maxRotateFunction = function (nums) {
  let n = nums.length;

  let sum = 0; // total sum
  let F = 0; // F(0)

  // Step 1: compute sum and F(0)
  for (let i = 0; i < n; i++) {
    sum += nums[i];
    F += i * nums[i];
  }

  let result = F;

  // Step 2: compute next rotations
  for (let k = 1; k < n; k++) {
    F = F + sum - n * nums[n - k];
    result = Math.max(result, F);
  }

  return result;
};
