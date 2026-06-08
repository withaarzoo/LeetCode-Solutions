/**
 * @param {number[]} nums
 * @param {number} pivot
 * @return {number[]}
 */
var pivotArray = function (nums, pivot) {
  // Store elements smaller than pivot
  const smaller = [];

  // Store elements equal to pivot
  const equal = [];

  // Store elements greater than pivot
  const greater = [];

  // Classify every element
  for (const num of nums) {
    if (num < pivot) {
      smaller.push(num);
    } else if (num === pivot) {
      equal.push(num);
    } else {
      greater.push(num);
    }
  }

  // Combine all three groups in required order
  return [...smaller, ...equal, ...greater];
};
