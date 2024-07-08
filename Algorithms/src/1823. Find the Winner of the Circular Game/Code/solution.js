/**
 * This function determines the winner of the game using the Josephus problem solution.
 * @param {number} n - The total number of people in the circle.
 * @param {number} k - The step count for each elimination.
 * @return {number} - The 1-indexed position of the winner.
 */
var findTheWinner = function (n, k) {
  // Call the helper function josephus and convert the result to 1-indexed
  return josephus(n, k) + 1;
};

/**
 * This is a helper function that implements the Josephus problem.
 * @param {number} n - The total number of people left in the circle.
 * @param {number} k - The step count for each elimination.
 * @return {number} - The 0-indexed position of the last remaining person.
 */
function josephus(n, k) {
  // Base case: if there's only one person left, they are the winner (position 0 in 0-indexing)
  if (n === 1) {
    return 0;
  }
  // Recursive case: find the position of the winner in the smaller circle
  // and adjust it using the modulus operation to get the correct position in the current circle
  return (josephus(n - 1, k) + k) % n;
}
