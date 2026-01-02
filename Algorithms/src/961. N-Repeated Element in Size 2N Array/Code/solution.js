/**
 * @param {number[]} nums
 * @return {number}
 */
var repeatedNTimes = function (nums) {
  const seen = new Set();

  for (let x of nums) {
    // If already present, this is the answer
    if (seen.has(x)) {
      return x;
    }
    // Otherwise, add to set
    seen.add(x);
  }
};
