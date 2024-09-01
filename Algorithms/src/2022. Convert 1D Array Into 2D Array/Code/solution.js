/**
 * This function converts a 1D array into a 2D array with specified dimensions.
 *
 * @param {number[]} original - The original 1D array that needs to be converted.
 * @param {number} m - The number of rows desired in the 2D array.
 * @param {number} n - The number of columns desired in the 2D array.
 * @return {number[][]} - The resulting 2D array, or an empty array if the conversion is not possible.
 */
var construct2DArray = function (original, m, n) {
  // Step 1: Check if the total number of elements in the original array
  // matches the required number of elements in the 2D array (m * n).
  // If not, it's impossible to form the desired 2D array, so return an empty array.
  if (original.length !== m * n) return [];

  // Step 2: Initialize an empty array to hold the rows of the 2D array.
  let result = [];

  // Step 3: Loop over the number of rows (m) to create each row of the 2D array.
  for (let i = 0; i < m; i++) {
    // For each row, slice a portion of the original array that corresponds to the row.
    // The slice starts at index i * n and ends just before index (i + 1) * n.
    // This gives us a sub-array with 'n' elements, representing a single row.
    result.push(original.slice(i * n, (i + 1) * n));
  }

  // Step 4: Return the fully constructed 2D array.
  return result;
};
