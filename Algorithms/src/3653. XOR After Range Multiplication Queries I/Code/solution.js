/**
 * @param {number[]} nums
 * @param {number[][]} queries
 * @return {number}
 */
var xorAfterQueries = function (nums, queries) {
  const MOD = 1000000007n;

  // Process each query
  for (const [l, r, k, v] of queries) {
    // Visit indices: l, l+k, l+2k, ... <= r
    for (let i = l; i <= r; i += k) {
      nums[i] = Number((BigInt(nums[i]) * BigInt(v)) % MOD);
    }
  }

  // Compute XOR of all final values
  let ans = 0;
  for (const num of nums) {
    ans ^= num;
  }

  return ans;
};
