/**
 * @param {number} n
 * @param {number} t
 * @return {number}
 */
var smallestNumber = function (n, t) {
  // Keep checking numbers until a valid one is found
  while (true) {
    let product = 1;
    let x = n;

    // Calculate the product of all digits
    while (x > 0) {
      product *= x % 10;
      x = Math.floor(x / 10);
    }

    // Return the first valid number
    if (product % t === 0) {
      return n;
    }

    // Check the next number
    n++;
  }
};
