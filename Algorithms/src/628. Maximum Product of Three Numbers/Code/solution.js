/**
 * @param {number[]} nums
 * @return {number}
 */
var maximumProduct = function (nums) {
  // Store the three largest numbers
  let max1 = -Infinity;
  let max2 = -Infinity;
  let max3 = -Infinity;

  // Store the two smallest numbers
  let min1 = Infinity;
  let min2 = Infinity;

  // Traverse the array once
  for (const num of nums) {
    // Update the three largest numbers
    if (num >= max1) {
      max3 = max2;
      max2 = max1;
      max1 = num;
    } else if (num >= max2) {
      max3 = max2;
      max2 = num;
    } else if (num >= max3) {
      max3 = num;
    }

    // Update the two smallest numbers
    if (num <= min1) {
      min2 = min1;
      min1 = num;
    } else if (num <= min2) {
      min2 = num;
    }
  }

  // Product of the three largest numbers
  const product1 = max1 * max2 * max3;

  // Product of the two smallest numbers and the largest number
  const product2 = min1 * min2 * max1;

  // Return the larger product
  return Math.max(product1, product2);
};
