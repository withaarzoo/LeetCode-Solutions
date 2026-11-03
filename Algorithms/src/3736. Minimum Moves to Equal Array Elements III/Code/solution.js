/**
 * @param {number[]} nums
 * @return {number}
 */
var minMoves = function (nums) {
  let mx = Math.max(...nums);
  let ans = 0;
  for (let x of nums) ans += mx - x;
  return ans;
};
