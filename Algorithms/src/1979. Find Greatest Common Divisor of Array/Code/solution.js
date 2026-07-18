/**
 * @param {number[]} nums
 * @return {number}
 */

// Function to find GCD using the Euclidean Algorithm
function gcd(a, b) {
  // Repeat until b becomes 0
  while (b !== 0) {
    let temp = b; // Store current b
    b = a % b; // Update b with remainder
    a = temp; // Move previous b into a
  }

  // a is the GCD
  return a;
}

var findGCD = function (nums) {
  // Initialize minimum and maximum
  let min = nums[0];
  let max = nums[0];

  // Find smallest and largest values
  for (const num of nums) {
    min = Math.min(min, num);
    max = Math.max(max, num);
  }

  // Return their GCD
  return gcd(min, max);
};
