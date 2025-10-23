/**
 * @param {string} s
 * @return {boolean}
 */
var hasSameDigits = function (s) {
  // convert to array of numbers
  let digits = new Array(s.length);
  for (let i = 0; i < s.length; ++i) digits[i] = s.charCodeAt(i) - 48;

  // reduce until length 2
  while (digits.length > 2) {
    const next = new Array(digits.length - 1);
    for (let i = 0; i + 1 < digits.length; ++i) {
      next[i] = (digits[i] + digits[i + 1]) % 10;
    }
    digits = next;
  }

  return digits.length === 2 && digits[0] === digits[1];
};
