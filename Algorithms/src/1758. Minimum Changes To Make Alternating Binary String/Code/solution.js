/**
 * @param {string} s
 * @return {number}
 */
var minOperations = function (s) {
  let startWith0 = 0;
  let startWith1 = 0;

  for (let i = 0; i < s.length; i++) {
    let expected0 = i % 2 === 0 ? "0" : "1";
    let expected1 = i % 2 === 0 ? "1" : "0";

    if (s[i] !== expected0) startWith0++;
    if (s[i] !== expected1) startWith1++;
  }

  return Math.min(startWith0, startWith1);
};
