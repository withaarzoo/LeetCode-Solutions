/**
 * @param {number} n
 * @return {number}
 */
var countTriples = function (n) {
  let count = 0;

  // Try all possible pairs (a, b)
  for (let a = 1; a <= n; a++) {
    for (let b = 1; b <= n; b++) {
      const sumSquares = a * a + b * b; // this should be c^2

      const c = Math.floor(Math.sqrt(sumSquares)); // integer square root

      // Check if c is within range and forms a perfect square
      if (c <= n && c * c === sumSquares) {
        count++; // (a, b, c) is a valid square triple
      }
    }
  }

  return count;
};
