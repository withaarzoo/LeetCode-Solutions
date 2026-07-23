/**
 * @param {number[]} nums
 * @return {number}
 */
var uniqueXorTriplets = function (nums) {
  // Length of the permutation
  const n = nums.length;

  // Handle small cases
  if (n <= 2) return n;

  // Count the number of bits in n
  let bits = 0;
  let x = n;
  while (x > 0) {
    bits++;
    x >>= 1;
  }

  // Number of possible XOR values
  return 1 << bits;
};
