/**
 * Function to return the numbers from 1 to n in lexicographical order.
 * Lexicographical order is dictionary order, where numbers are compared digit by digit.
 * @param {number} n - The maximum number to consider.
 * @return {number[]} - An array of numbers from 1 to n in lexicographical order.
 */
var lexicalOrder = function (n) {
  // Initialize an empty array to store the result.
  let result = [];

  // Start the Depth-First Search (DFS) traversal from numbers 1 to 9.
  // These are the first possible digits in lexicographical order.
  for (let i = 1; i <= 9; i++) {
    dfs(i, n, result);
  }

  // Return the final result array containing numbers in lexicographical order.
  return result;
};

/**
 * Helper function to perform DFS traversal.
 * This function generates the lexicographical sequence by recursively constructing numbers.
 * @param {number} curr - The current number in the sequence.
 * @param {number} n - The maximum number limit.
 * @param {number[]} result - The result array where numbers are collected.
 */
function dfs(curr, n, result) {
  // If the current number exceeds n, stop the recursion (base case).
  if (curr > n) return;

  // Add the current number to the result array.
  result.push(curr);

  // Explore the next possible numbers by appending digits (0 to 9) to the current number.
  for (let i = 0; i <= 9; i++) {
    // Calculate the next number by appending 'i' to the current number.
    let next = curr * 10 + i;

    // If the next number exceeds n, there's no need to explore further.
    if (next > n) return;

    // Recursively perform DFS with the new number.
    dfs(next, n, result);
  }
}
