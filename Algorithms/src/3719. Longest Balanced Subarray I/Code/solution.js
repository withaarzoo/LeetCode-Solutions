/**
 * @param {number[]} nums
 * @return {number}
 */
var longestBalanced = function (nums) {
  const n = nums.length;
  let ans = 0;

  for (let i = 0; i < n; i++) {
    let evenSet = new Set();
    let oddSet = new Set();

    for (let j = i; j < n; j++) {
      if (nums[j] % 2 === 0) evenSet.add(nums[j]);
      else oddSet.add(nums[j]);

      if (evenSet.size === oddSet.size) ans = Math.max(ans, j - i + 1);
    }
  }
  return ans;
};
