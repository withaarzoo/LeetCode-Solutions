/**
 * @param {number[]} prices
 * @return {number}
 */
var getDescentPeriods = function (prices) {
  let ans = 1; // first day
  let len = 1; // current smooth descent length

  for (let i = 1; i < prices.length; i++) {
    if (prices[i] === prices[i - 1] - 1) {
      len++;
    } else {
      len = 1;
    }
    ans += len;
  }
  return ans;
};
