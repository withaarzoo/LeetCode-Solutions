/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findMissingElements = function (nums) {
  // Store all numbers for constant-time lookup
  const seen = new Set(nums);

  // Find the minimum and maximum values
  let mn = Math.min(...nums);
  let mx = Math.max(...nums);

  // Store missing numbers
  const ans = [];

  // Check every value in the range
  for (let x = mn; x <= mx; x++) {
    // If the number is missing, save it
    if (!seen.has(x)) {
      ans.push(x);
    }
  }

  return ans;
};
