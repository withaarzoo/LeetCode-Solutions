/**
 * @param {number} n
 * @return {number}
 */
var smallestNumber = function (n) {
  // Increment k until (2^k - 1) >= n
  let k = 1;
  while (true) {
    // (1 << k) might overflow for large k in JS, but for n <= 1000 it's fine.
    let val = (1 << k) - 1;
    if (val >= n) return val;
    k++;
  }
};
