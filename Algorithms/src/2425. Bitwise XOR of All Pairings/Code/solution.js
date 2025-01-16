/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var xorAllNums = function (nums1, nums2) {
  let xor1 = 0,
    xor2 = 0;

  // XOR all elements in nums1
  for (const num of nums1) {
    xor1 ^= num;
  }

  // XOR all elements in nums2
  for (const num of nums2) {
    xor2 ^= num;
  }

  // If nums1 has odd length, include xor2
  // If nums2 has odd length, include xor1
  return (nums1.length % 2 ? xor2 : 0) ^ (nums2.length % 2 ? xor1 : 0);
};
