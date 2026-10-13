/**
 * @param {number[]} nums
 * @return {number}
 */
var maxAlternatingSum = function (nums) {
  const n = nums.length;
  const sq = nums.map((x) => x * x);
  sq.sort((a, b) => b - a);
  let ans = 0;
  const pos = Math.floor((n + 1) / 2);
  for (let i = 0; i < n; i++) {
    if (i < pos) ans += sq[i];
    else ans -= sq[i];
  }
  return ans;
};
