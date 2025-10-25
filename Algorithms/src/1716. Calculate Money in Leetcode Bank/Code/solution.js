/**
 * @param {number} n
 * @return {number}
 */
var totalMoney = function (n) {
  const w = Math.floor(n / 7); // number of full weeks
  const r = n % 7; // remaining days
  // sum of full weeks: w*28 + 7 * (0 + 1 + ... + (w-1))
  const fullWeeksSum = w * 28 + 7 * ((w * (w - 1)) / 2);
  // sum of remaining days: r*(1 + w) + r*(r-1)/2
  const remSum = r * (1 + w) + (r * (r - 1)) / 2;
  return fullWeeksSum + remSum;
};
