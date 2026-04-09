/**
 * @param {number[]} nums
 * @param {number[][]} queries
 * @return {number}
 */
var xorAfterQueries = function (nums, queries) {
  const MOD = 1000000007;
  const n = nums.length;
  const limit = Math.floor(Math.sqrt(n)) + 1;

  function modPow(base, exp) {
    let result = 1n;
    let b = BigInt(base);
    let e = BigInt(exp);
    const mod = BigInt(MOD);

    while (e > 0n) {
      if (e & 1n) {
        result = (result * b) % mod;
      }

      b = (b * b) % mod;
      e >>= 1n;
    }

    return result;
  }

  function modInverse(x) {
    return modPow(x, MOD - 2);
  }

  const smallQueries = new Map();

  for (const [l, r, k, v] of queries) {
    if (k >= limit) {
      for (let i = l; i <= r; i += k) {
        nums[i] = Number((BigInt(nums[i]) * BigInt(v)) % BigInt(MOD));
      }
    } else {
      if (!smallQueries.has(k)) {
        smallQueries.set(k, []);
      }

      smallQueries.get(k).push([l, r, v]);
    }
  }

  for (const [k, group] of smallQueries.entries()) {
    const diff = Array(n).fill(1n);

    for (const [l, r, v] of group) {
      diff[l] = (diff[l] * BigInt(v)) % BigInt(MOD);

      const steps = Math.floor((r - l) / k);
      const nextPos = l + (steps + 1) * k;

      if (nextPos < n) {
        diff[nextPos] = (diff[nextPos] * modInverse(v)) % BigInt(MOD);
      }
    }

    for (let i = 0; i < n; i++) {
      if (i >= k) {
        diff[i] = (diff[i] * diff[i - k]) % BigInt(MOD);
      }

      nums[i] = Number((BigInt(nums[i]) * diff[i]) % BigInt(MOD));
    }
  }

  let answer = 0;

  for (const num of nums) {
    answer ^= num;
  }

  return answer;
};
