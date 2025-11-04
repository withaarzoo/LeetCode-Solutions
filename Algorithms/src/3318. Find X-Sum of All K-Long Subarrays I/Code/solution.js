/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} x
 * @return {number[]}
 */
var findXSum = function (nums, k, x) {
  const n = nums.length;
  const ans = [];
  const freq = new Map();

  // Build initial window
  for (let i = 0; i < k; i++) {
    freq.set(nums[i], (freq.get(nums[i]) || 0) + 1);
  }

  ans.push(computeXSum(freq, x));

  // Slide window
  for (let i = k; i < n; i++) {
    const add = nums[i];
    const rem = nums[i - k];

    freq.set(add, (freq.get(add) || 0) + 1);
    const fr = (freq.get(rem) || 0) - 1;
    if (fr === 0) freq.delete(rem);
    else freq.set(rem, fr);

    ans.push(computeXSum(freq, x));
  }

  return ans;

  function computeXSum(map, x) {
    const items = [];
    for (const [v, f] of map.entries()) items.push([v, f]);
    // sort by freq desc, value desc
    items.sort((a, b) => {
      if (a[1] !== b[1]) return b[1] - a[1];
      return b[0] - a[0];
    });
    let sum = 0;
    const take = Math.min(x, items.length);
    for (let i = 0; i < take; i++) {
      sum += items[i][0] * items[i][1];
    }
    return sum;
  }
};
