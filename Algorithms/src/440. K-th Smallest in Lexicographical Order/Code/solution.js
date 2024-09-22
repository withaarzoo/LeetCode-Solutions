/**
 * Helper function to count how many steps we can take starting from the current prefix
 * and going through all its children (within the range of numbers from 1 to n).
 * This helps in determining how many numbers exist between the current prefix
 * and the next prefix in a lexicographical order.
 *
 * @param {number} curr - The current prefix number we're examining.
 * @param {number} n - The upper limit of numbers we are considering.
 * @return {number} - The total number of valid numbers starting with the current prefix.
 */
function countSteps(curr, n) {
  let steps = 0; // Initialize the step count
  let first = curr; // The smallest number starting with the current prefix
  let last = curr; // The largest number starting with the current prefix

  // Expand the range [first, last] to cover all possible numbers starting with the current prefix.
  // This loop calculates how many numbers are between 'first' and 'last' at each depth level (ones, tens, hundreds, etc.).
  while (first <= n) {
    // Add the number of valid numbers between 'first' and 'last'.
    // We cap the range at 'n' since we can't go beyond it.
    steps += Math.min(n + 1, last + 1) - first;

    // Move to the next level (increase the digit length by 1).
    // E.g., if first is 1, it becomes 10, then 100, etc.
    first *= 10;

    // Similarly, increase 'last' to cover the next level.
    // E.g., if last is 1, it becomes 19, then 199, etc.
    last = last * 10 + 9;
  }

  return steps; // Return the total number of steps for this prefix.
}

/**
 * Main function to find the k-th lexicographically smallest number from 1 to n.
 * We use a greedy approach with a prefix tree-like traversal, moving down levels of prefixes
 * or skipping entire prefixes when we know there aren't enough valid numbers left.
 *
 * @param {number} n - The upper limit of numbers we are considering (i.e., 1 to n).
 * @param {number} k - The position (k-th smallest) we want to find in lexicographical order.
 * @return {number} - The k-th lexicographically smallest number.
 */
var findKthNumber = function (n, k) {
  let curr = 1; // Start at the first number (the prefix "1")
  k--; // Decrease k by 1 to make it zero-indexed for easier calculations

  // Keep searching until we've found the k-th number.
  while (k > 0) {
    // Calculate how many numbers exist under the current prefix.
    let steps = countSteps(curr, n);

    // If the number of steps is less than or equal to k, we can skip this prefix.
    if (steps <= k) {
      // Move to the next lexicographical prefix (e.g., from 1 to 2).
      curr++;

      // Decrease k by the number of skipped numbers, since we move past these.
      k -= steps;
    } else {
      // If there are more steps than k, we must go deeper into the current prefix's subtree.
      // This means we move to the first child of the current prefix (e.g., from 1 to 10).
      curr *= 10;

      // We've now considered one number, so decrement k by 1.
      k--;
    }
  }

  // When the loop finishes, 'curr' will be the k-th lexicographically smallest number.
  return curr;
};
