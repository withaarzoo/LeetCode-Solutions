/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var resultsArray = function (nums, k) {
  let n = nums.length;
  let result = [];

  for (let i = 0; i <= n - k; i++) {
    let subarray = nums.slice(i, i + k);
    let sortedSubarray = [...subarray].sort((a, b) => a - b);

    // Check if elements are consecutive
    let isConsecutive = true;
    for (let j = 1; j < k; j++) {
      if (sortedSubarray[j] - sortedSubarray[j - 1] !== 1) {
        isConsecutive = false;
        break;
      }
    }

    // Add the result based on conditions
    if (
      isConsecutive &&
      subarray.every((val, idx) => val === sortedSubarray[idx])
    ) {
      result.push(sortedSubarray[k - 1]); // Max element
    } else {
      result.push(-1);
    }
  }

  return result;
};
