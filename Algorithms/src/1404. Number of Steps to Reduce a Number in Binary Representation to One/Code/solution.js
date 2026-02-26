/**
 * @param {string} s
 * @return {number}
 */
var numSteps = function (s) {
  let steps = 0;
  let carry = 0;

  // Traverse from right to left (ignore first bit)
  for (let i = s.length - 1; i > 0; i--) {
    let bit = s[i] - "0" + carry;

    if (bit === 1) {
      // Odd case
      steps += 2;
      carry = 1;
    } else {
      // Even case
      steps += 1;
    }
  }

  return steps + carry;
};
