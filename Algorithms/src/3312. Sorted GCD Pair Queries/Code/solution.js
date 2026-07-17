/**
 * @param {number[]} nums
 * @param {number[]} queries
 * @return {number[]}
 */
var gcdValues = function (nums, queries) {
  // Find maximum value.
  let mx = 0;
  for (const x of nums) mx = Math.max(mx, x);

  // Frequency array.
  const freq = new Array(mx + 1).fill(0);
  for (const x of nums) freq[x]++;

  // exact[g] = pairs with GCD exactly g.
  const exact = new Array(mx + 1).fill(0);

  // Compute exact counts.
  for (let g = mx; g >= 1; g--) {
    let cnt = 0;

    // Count divisible numbers.
    for (let m = g; m <= mx; m += g) cnt += freq[m];

    let pairs = (cnt * (cnt - 1)) / 2;

    // Remove larger GCD contributions.
    for (let m = g * 2; m <= mx; m += g) pairs -= exact[m];

    exact[g] = pairs;
  }

  // Prefix sums.
  const prefix = new Array(mx + 1).fill(0);
  for (let g = 1; g <= mx; g++) prefix[g] = prefix[g - 1] + exact[g];

  const ans = [];

  for (const q of queries) {
    let l = 1,
      r = mx;

    // Binary search.
    while (l < r) {
      const mid = (l + r) >> 1;

      if (prefix[mid] >= q + 1) r = mid;
      else l = mid + 1;
    }

    ans.push(l);
  }

  return ans;
};
