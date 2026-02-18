/**
 * @param {number} n
 * @return {boolean}
 */
var hasAlternatingBits = function (n) {
  let prev = n & 1;
  n = n >> 1;

  while (n > 0) {
    let curr = n & 1;

    if (curr === prev) {
      return false;
    }

    prev = curr;
    n = n >> 1;
  }

  return true;
};
