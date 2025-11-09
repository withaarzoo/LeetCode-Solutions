/**
 * @param {number} num1
 * @param {number} num2
 * @return {number}
 */
var countOperations = function (num1, num2) {
  let a = num1,
    b = num2;
  let ops = 0;
  while (a > 0 && b > 0) {
    if (a < b) {
      const t = a;
      a = b;
      b = t;
    } // ensure a >= b
    ops += Math.floor(a / b); // number of b's inside a
    a = a % b; // leftover after removing them
  }
  return ops;
};
