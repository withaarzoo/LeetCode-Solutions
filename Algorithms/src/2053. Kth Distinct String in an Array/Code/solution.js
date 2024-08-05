/**
 * Function to find the k-th distinct string in an array.
 * @param {string[]} arr - The array of strings.
 * @param {number} k - The position of the distinct string to find.
 * @return {string} - The k-th distinct string or an empty string if it doesn't exist.
 */
var kthDistinct = function (arr, k) {
  let count = new Map(); // Create a map to count occurrences of each string
  let distinct = []; // Array to store distinct strings

  // Count occurrences of each string
  for (let str of arr) {
    count.set(str, (count.get(str) || 0) + 1); // Increment count for each string
  }

  // Collect distinct strings in order
  for (let str of arr) {
    if (count.get(str) === 1) {
      // Check if the string is distinct
      distinct.push(str); // Add distinct string to the array
    }
  }

  // Return the k-th distinct string or an empty string if it doesn't exist
  return k <= distinct.length ? distinct[k - 1] : ""; // Adjust index for 1-based k
};
