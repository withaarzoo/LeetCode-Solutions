/**
 * @param {number[]} nums
 * @param {number} lower
 * @param {number} upper
 * @return {number}
 */
var countFairPairs = function (nums, lower, upper) {
  nums.sort((a, b) => a - b);
  let count = 0;
  const n = nums.length;

  for (let i = 0; i < n - 1; i++) {
    const minVal = lower - nums[i];
    const maxVal = upper - nums[i];

    // Using binary search
    const start = lowerBound(nums, minVal, i + 1);
    const end = upperBound(nums, maxVal, i + 1);

    count += end - start;
  }

  return count;
};

// Helper functions for binary search
function lowerBound(arr, target, start) {
  let low = start,
    high = arr.length;
  while (low < high) {
    const mid = low + Math.floor((high - low) / 2);
    if (arr[mid] < target) low = mid + 1;
    else high = mid;
  }
  return low;
}

function upperBound(arr, target, start) {
  let low = start,
    high = arr.length;
  while (low < high) {
    const mid = low + Math.floor((high - low) / 2);
    if (arr[mid] <= target) low = mid + 1;
    else high = mid;
  }
  return low;
}
