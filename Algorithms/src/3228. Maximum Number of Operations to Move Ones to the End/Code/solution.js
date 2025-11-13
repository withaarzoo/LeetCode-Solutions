/**
 * @param {string} s
 * @return {number}
 */
var maxOperations = function (s) {
  let ans = 0n; // use BigInt to be safe for large intermediate sums
  let ones = 0n;
  for (let i = 0; i < s.length; ++i) {
    if (s[i] === "1") {
      ones += 1n;
    } else {
      // '0'
      if (i > 0 && s[i - 1] === "1") ans += ones;
    }
  }
  // convert BigInt back to Number (fits in JS Number for given constraints in practice)
  return Number(ans);
};
