/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var getCommon = function (nums1, nums2) {
  // Pointer for nums1
  let i = 0;

  // Pointer for nums2
  let j = 0;

  // Traverse both arrays
  while (i < nums1.length && j < nums2.length) {
    // If values are equal,
    // return the common value
    if (nums1[i] === nums2[j]) {
      return nums1[i];
    }

    // Move pointer having smaller value
    if (nums1[i] < nums2[j]) {
      i++;
    } else {
      j++;
    }
  }

  // No common value found
  return -1;
};
