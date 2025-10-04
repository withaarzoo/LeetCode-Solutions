/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let left = 0; // left pointer
  let right = height.length - 1; // right pointer
  let maxArea = 0; // best area found

  while (left < right) {
    const width = right - left;
    const h = Math.min(height[left], height[right]);
    const area = h * width;
    if (area > maxArea) maxArea = area;

    // move the pointer with smaller height inward
    if (height[left] < height[right]) {
      left++;
    } else {
      right--;
    }
  }
  return maxArea;
};
