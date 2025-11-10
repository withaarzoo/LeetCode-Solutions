/**
 * @param {number[]} nums
 * @return {number}
 */
var minOperations = function (nums) {
  const stk = []; // non-decreasing stack
  let ans = 0;
  for (const x of nums) {
    while (stk.length && stk[stk.length - 1] > x) stk.pop();
    if (x === 0) continue;
    if (!stk.length || stk[stk.length - 1] < x) {
      ans++;
      stk.push(x);
    }
  }
  return ans;
};
