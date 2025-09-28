/**
 * @param {number[]} nums
 * @return {number}
 */
var largestPerimeter = function (nums) {
  // Sort ascending
  nums.sort((x, y) => x - y);
  // Start from the largest end and check triples
  for (let i = nums.length - 1; i >= 2; --i) {
    const a = nums[i]; // largest in triple
    const b = nums[i - 1];
    const c = nums[i - 2];
    if (b + c > a) return a + b + c;
  }
  return 0;
};
