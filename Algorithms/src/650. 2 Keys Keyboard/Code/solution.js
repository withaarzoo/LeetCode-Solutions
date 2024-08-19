/**
 * @param {number} n - The target number for which we want to determine the minimum number of operations.
 * @return {number} - The minimum number of operations needed to achieve the target number.
 */
var minSteps = function (n) {
  // Initialize a variable to keep track of the total number of operations required.
  let operations = 0;

  // Start a loop from 2 up to and including the target number `n`.
  // We start from 2 because 1 is not a valid factor for division in this context.
  for (let i = 2; i <= n; i++) {
    // As long as `n` is divisible by `i`, we can perform operations.
    // This is equivalent to "pasting" the previously copied sequence.
    while (n % i === 0) {
      // Add the current divisor `i` to the operation count.
      // This represents the series of "Copy All" and "Paste" operations required.
      operations += i;

      // Update `n` by dividing it by `i`.
      // This reduces `n` by removing the effect of the last "Paste" operation.
      n /= i;
    }
  }

  // Return the total number of operations required to construct the string of length `n`.
  return operations;
};
