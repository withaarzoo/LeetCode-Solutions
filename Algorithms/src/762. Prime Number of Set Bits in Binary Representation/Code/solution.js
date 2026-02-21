/**
 * @param {number} left
 * @param {number} right
 * @return {number}
 */
var countPrimeSetBits = function (left, right) {
  const primeSet = new Set([2, 3, 5, 7, 11, 13, 17, 19]);

  let ans = 0;

  for (let num = left; num <= right; num++) {
    // Convert to binary and count 1s
    let setBits = num.toString(2).split("1").length - 1;

    if (primeSet.has(setBits)) {
      ans++;
    }
  }

  return ans;
};
