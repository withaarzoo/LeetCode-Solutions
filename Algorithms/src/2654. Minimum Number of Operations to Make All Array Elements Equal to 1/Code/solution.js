/**
 * @param {number[]} nums
 * @return {number}
 */
var minOperations = function (nums) {
  const n = nums.length;

  // 1) If there are ones, we only need to fix the rest.
  let ones = 0;
  for (const x of nums) if (x === 1) ones++;
  if (ones > 0) return n - ones;

  // gcd helper
  const gcd = (a, b) => {
    while (b !== 0) {
      const t = a % b;
      a = b;
      b = t;
    }
    return Math.abs(a);
  };

  // 2) If global gcd > 1, impossible.
  let g = 0;
  for (const x of nums) g = gcd(g, x);
  if (g > 1) return -1;

  // 3) Find shortest subarray with gcd == 1.
  let best = Number.POSITIVE_INFINITY;
  for (let i = 0; i < n; i++) {
    let cur = 0;
    for (let j = i; j < n; j++) {
      cur = gcd(cur, nums[j]);
      if (cur === 1) {
        best = Math.min(best, j - i + 1);
        break;
      }
    }
  }
  // 4) Create first 1 and spread it.
  return best - 1 + (n - 1);
};
