/**
 * @param {number[]} arr
 * @return {number}
 */
var findLengthOfShortestSubarray = function (arr) {
  const n = arr.length;

  // Step 1: Find the longest non-decreasing prefix
  let left = 0;
  while (left + 1 < n && arr[left] <= arr[left + 1]) {
    left++;
  }

  // If the entire array is already sorted
  if (left === n - 1) return 0;

  // Step 2: Find the longest non-decreasing suffix
  let right = n - 1;
  while (right > 0 && arr[right - 1] <= arr[right]) {
    right--;
  }

  // Step 3: Find the minimum length to remove by comparing prefix and suffix
  let result = Math.min(n - left - 1, right);

  // Step 4: Use two pointers to find the smallest middle part to remove
  let i = 0,
    j = right;
  while (i <= left && j < n) {
    if (arr[i] <= arr[j]) {
      result = Math.min(result, j - i - 1);
      i++;
    } else {
      j++;
    }
  }

  return result;
};
