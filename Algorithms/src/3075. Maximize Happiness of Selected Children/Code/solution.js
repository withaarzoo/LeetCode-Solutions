/**
 * @param {number[]} happiness
 * @param {number} k
 * @return {number}
 */
var maximumHappinessSum = function (happiness, k) {
  // Sort descending
  happiness.sort((a, b) => b - a);

  let ans = 0;

  for (let i = 0; i < k; i++) {
    let curr = happiness[i] - i;
    if (curr > 0) {
      ans += curr;
    }
  }

  return ans;
};
