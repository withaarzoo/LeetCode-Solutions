/**
 * @param {number[][]} matrix
 * @return {number}
 */
var maxMatrixSum = function (matrix) {
  let sum = 0;
  let negativeCount = 0;
  let minAbs = Infinity;

  for (let row of matrix) {
    for (let val of row) {
      sum += Math.abs(val); // add absolute value
      if (val < 0) negativeCount++;
      minAbs = Math.min(minAbs, Math.abs(val));
    }
  }

  if (negativeCount % 2 === 1) {
    sum -= 2 * minAbs;
  }

  return sum;
};
