/**
 * @param {number[]} nums
 * @return {number}
 */
var minimumDistance = function (nums) {
  const map = new Map();

  // Store all indices for each value
  for (let i = 0; i < nums.length; i++) {
    if (!map.has(nums[i])) {
      map.set(nums[i], []);
    }
    map.get(nums[i]).push(i);
  }

  let ans = Infinity;

  // Process each value's indices
  for (const indices of map.values()) {
    if (indices.length < 3) continue;

    // Check every consecutive group of 3 indices
    for (let i = 0; i + 2 < indices.length; i++) {
      const distance = 2 * (indices[i + 2] - indices[i]);
      ans = Math.min(ans, distance);
    }
  }

  return ans === Infinity ? -1 : ans;
};
