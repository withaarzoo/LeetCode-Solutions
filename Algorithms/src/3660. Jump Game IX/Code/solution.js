/**
 * @param {number[]} nums
 * @return {number[]}
 */
var maxValue = function (nums) {
  const n = nums.length;

  // suffixMin[i] = smallest value in nums[i...n-1]
  // I add one extra slot so the last segment stops naturally.
  const suffixMin = new Array(n + 1);
  suffixMin[n] = Infinity;
  for (let i = n - 1; i >= 0; i--) {
    suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
  }

  const ans = new Array(n);
  let l = 0;

  // I split the array into connected components.
  while (l < n) {
    let r = l;
    let componentMax = nums[l];

    // I keep expanding while the current segment still has an inversion
    // crossing the next boundary.
    while (r + 1 < n && componentMax > suffixMin[r + 1]) {
      r++;
      componentMax = Math.max(componentMax, nums[r]);
    }

    // Every index inside this component gets the same best reachable value.
    for (let i = l; i <= r; i++) {
      ans[i] = componentMax;
    }

    // Move to the next component.
    l = r + 1;
  }

  return ans;
};
