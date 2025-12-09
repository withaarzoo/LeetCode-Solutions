/**
 * @param {number[]} nums
 * @return {number}
 */
var specialTriplets = function (nums) {
  const MOD = 1_000_000_007;
  const right = new Map();
  const left = new Map();

  // Fill 'right' with frequencies
  for (const x of nums) {
    right.set(x, (right.get(x) || 0) + 1);
  }

  let ans = 0;

  for (const x of nums) {
    // x is now the middle element, so remove it from right
    right.set(x, right.get(x) - 1);

    const target = x * 2; // value 2x

    const cntLeft = left.get(target) || 0;
    const cntRight = right.get(target) || 0;

    const add = (cntLeft * cntRight) % MOD;
    ans = (ans + add) % MOD;

    // Move x to left
    left.set(x, (left.get(x) || 0) + 1);
  }

  return ans;
};
