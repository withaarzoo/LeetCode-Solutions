/**
 * @param {number[]} arr
 * @return {number}
 */
var maximumElementAfterDecrementingAndRearranging = function (arr) {
  // Sort the array in increasing order.
  arr.sort((a, b) => a - b);

  // The first element must become 1.
  arr[0] = 1;

  // Build the largest valid sequence.
  for (let i = 1; i < arr.length; i++) {
    // The current value cannot exceed previous + 1.
    arr[i] = Math.min(arr[i], arr[i - 1] + 1);
  }

  // Return the largest value.
  return arr[arr.length - 1];
};
