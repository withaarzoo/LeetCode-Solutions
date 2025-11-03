/**
 * @param {string} colors
 * @param {number[]} neededTime
 * @return {number}
 */
var minCost = function (colors, neededTime) {
  let ans = 0;
  let blockSum = 0;
  let blockMax = 0;
  const n = colors.length;

  for (let i = 0; i < n; ++i) {
    if (i > 0 && colors[i] !== colors[i - 1]) {
      ans += blockSum - blockMax;
      blockSum = 0;
      blockMax = 0;
    }
    blockSum += neededTime[i];
    blockMax = Math.max(blockMax, neededTime[i]);
  }
  ans += blockSum - blockMax;
  return ans;
};
