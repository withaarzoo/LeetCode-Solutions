/**
 * @param {number[]} arr
 * @return {number[][]}
 */
var minimumAbsDifference = function (arr) {
  // Step 1: Sort the array
  arr.sort((a, b) => a - b);

  let minDiff = Infinity;
  let result = [];

  // Step 2: Find minimum difference
  for (let i = 1; i < arr.length; i++) {
    minDiff = Math.min(minDiff, arr[i] - arr[i - 1]);
  }

  // Step 3: Collect valid pairs
  for (let i = 1; i < arr.length; i++) {
    if (arr[i] - arr[i - 1] === minDiff) {
      result.push([arr[i - 1], arr[i]]);
    }
  }

  return result;
};
