/**
 * @param {number[]} nums
 * @param {number} original
 * @return {number}
 */
var findFinalValue = function (nums, original) {
  // Create a Set for constant time membership checks
  const s = new Set(nums);
  // While original is in the set, multiply by 2
  while (s.has(original)) {
    original *= 2;
  }
  return original;
};
