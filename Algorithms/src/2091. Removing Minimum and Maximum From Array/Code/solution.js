/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDeletions = function (nums) {
  // Store the total number of elements in the array.
  const n = nums.length;

  // Start by assuming the first element is both minimum and maximum.
  let minIndex = 0;
  let maxIndex = 0;

  // Find the positions of the minimum and maximum elements.
  for (let i = 1; i < n; i++) {
    // Update the minimum index if a smaller value is found.
    if (nums[i] < nums[minIndex]) {
      minIndex = i;
    }

    // Update the maximum index if a larger value is found.
    if (nums[i] > nums[maxIndex]) {
      maxIndex = i;
    }
  }

  // Remove everything from the front up to the farther special element.
  const removeFromFront = Math.max(minIndex, maxIndex) + 1;

  // Remove everything from the back up to the farther special element.
  const removeFromBack = n - Math.min(minIndex, maxIndex);

  // Calculate both ways of removing one element from each side.
  const removeFromBothSides = Math.min(
    minIndex + 1 + (n - maxIndex),
    maxIndex + 1 + (n - minIndex),
  );

  // Return the minimum deletions among all possible strategies.
  return Math.min(removeFromFront, removeFromBack, removeFromBothSides);
};
