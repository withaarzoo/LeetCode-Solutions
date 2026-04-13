/**
 * @param {number[]} nums
 * @param {number} target
 * @param {number} start
 * @return {number}
 */
var getMinDistance = function (nums, target, start) {
  // Store the minimum distance found so far
  let answer = Infinity;

  // Traverse through the array
  for (let i = 0; i < nums.length; i++) {
    // Check if current element is the target
    if (nums[i] === target) {
      // Update the minimum distance
      answer = Math.min(answer, Math.abs(i - start));
    }
  }

  return answer;
};
