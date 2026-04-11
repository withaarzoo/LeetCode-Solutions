/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDistance = function (nums) {
  const positions = new Map();

  // Store all indices for each value
  for (let i = 0; i < nums.length; i++) {
    if (!positions.has(nums[i])) {
      positions.set(nums[i], []);
    }
    positions.get(nums[i]).push(i);
  }

  let ans = Infinity;

  // Check every value's index list
  for (const idx of positions.values()) {
    if (idx.length < 3) continue;

    // Check every consecutive group of 3 indices
    for (let i = 0; i + 2 < idx.length; i++) {
      const distance = 2 * (idx[i + 2] - idx[i]);
      ans = Math.min(ans, distance);
    }
  }

  return ans === Infinity ? -1 : ans;
};
