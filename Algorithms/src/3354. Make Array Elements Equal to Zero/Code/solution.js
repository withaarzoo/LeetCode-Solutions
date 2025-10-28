/**
 * @param {number[]} nums
 * @return {number}
 */
var countValidSelections = function (nums) {
  const n = nums.length;

  const simulate = (arr, start, dir) => {
    const a = arr.slice(); // copy
    let curr = start;
    while (curr >= 0 && curr < n) {
      if (a[curr] === 0) {
        curr += dir; // move same direction
      } else {
        a[curr] -= 1; // decrement
        dir = -dir; // reverse
        curr += dir; // step in new direction
      }
    }
    // check all zeros
    for (let v of a) if (v !== 0) return false;
    return true;
  };

  let ans = 0;
  for (let i = 0; i < n; ++i) {
    if (nums[i] !== 0) continue;
    if (simulate(nums, i, -1)) ans++; // left
    if (simulate(nums, i, +1)) ans++; // right
  }
  return ans;
};
