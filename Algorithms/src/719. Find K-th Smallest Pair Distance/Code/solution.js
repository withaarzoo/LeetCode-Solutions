/**
 * @param {number[]} nums - An array of integers.
 * @param {number} k - The k-th smallest distance pair to find.
 * @return {number} - The k-th smallest distance pair.
 */
var smallestDistancePair = function (nums, k) {
  /**
   * Helper function to count the number of pairs with a distance
   * less than or equal to the given 'mid' value.
   *
   * @param {number[]} nums - Sorted array of integers.
   * @param {number} mid - The maximum allowed distance for pairs.
   * @return {number} - The count of pairs with a distance <= 'mid'.
   */
  const countPairs = (nums, mid) => {
    let count = 0; // Initialize the pair count to zero
    let j = 0; // Initialize the right pointer

    // Iterate through the array with the left pointer 'i'
    for (let i = 0; i < nums.length; i++) {
      // Move the right pointer 'j' to find pairs where the difference is <= 'mid'
      while (j < nums.length && nums[j] - nums[i] <= mid) {
        j++;
      }
      // The number of valid pairs with 'i' as the left element is (j - i - 1)
      count += j - i - 1;
    }
    return count; // Return the total count of valid pairs
  };

  // Sort the array to facilitate the binary search approach
  nums.sort((a, b) => a - b);

  // Define the search space for binary search
  let low = 0; // The minimum possible distance
  let high = nums[nums.length - 1] - nums[0]; // The maximum possible distance

  // Perform binary search to find the k-th smallest distance pair
  while (low < high) {
    // Calculate the midpoint of the current search space
    let mid = Math.floor((low + high) / 2);

    // Use the helper function to count pairs with distance <= 'mid'
    if (countPairs(nums, mid) >= k) {
      // If there are at least 'k' pairs with distance <= 'mid', search the lower half
      high = mid;
    } else {
      // Otherwise, search the upper half
      low = mid + 1;
    }
  }

  // The 'low' value will be the k-th smallest distance pair after the loop
  return low;
};
