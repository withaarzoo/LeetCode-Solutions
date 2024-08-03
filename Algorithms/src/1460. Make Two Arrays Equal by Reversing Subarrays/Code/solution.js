/**
 * Determines if one array can be made equal to another by sorting.
 *
 * @param {number[]} target - The target array to compare against.
 * @param {number[]} arr - The array that needs to be checked for equality with the target array.
 * @return {boolean} - Returns true if the sorted arrays are equal, otherwise false.
 */
var canBeEqual = function (target, arr) {
  // Step 1: Sort the target array in ascending order
  target.sort((a, b) => a - b);

  // Step 2: Sort the arr array in ascending order
  arr.sort((a, b) => a - b);

  // Step 3: Convert both sorted arrays to strings and compare them
  // If both strings are equal, it means both arrays have the same elements in the same order
  return target.toString() === arr.toString();
};
