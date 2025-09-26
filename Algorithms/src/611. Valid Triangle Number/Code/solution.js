/**
 * @param {number[]} nums
 * @return {number}
 */
var triangleNumber = function (nums) {
  nums.sort((a, b) => a - b); // sort ascending
  const n = nums.length;
  let count = 0;
  for (let k = n - 1; k >= 2; k--) {
    let i = 0,
      j = k - 1; // two pointers
    while (i < j) {
      if (nums[i] + nums[j] > nums[k]) {
        count += j - i; // all i..j-1 with j are valid
        j--; // try smaller b
      } else {
        i++; // need larger a
      }
    }
  }
  return count;
};
