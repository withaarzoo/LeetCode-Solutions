/**
 * @param {string} s
 * @return {string}
 */
var makeFancyString = function (s) {
  let result = "";
  for (let i = 0; i < s.length; i++) {
    let n = result.length;
    if (n < 2 || !(result[n - 1] === s[i] && result[n - 2] === s[i])) {
      result += s[i];
    }
  }
  return result;
};
