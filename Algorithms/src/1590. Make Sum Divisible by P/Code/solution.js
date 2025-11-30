/**
 * @param {number[]} nums
 * @param {number} p
 * @return {number}
 */
var minSubarray = function (nums, p) {
  let total = 0;
  for (const x of nums) {
    total = (total + x) % p; // keep modulo
  }

  const need = total;
  if (need === 0) return 0; // already divisible

  const n = nums.length;
  const lastIndex = new Map();
  lastIndex.set(0, -1); // prefix before index 0

  let ans = n;
  let prefix = 0;

  for (let i = 0; i < n; i++) {
    prefix = (prefix + nums[i]) % p;

    const prefMod = prefix;
    let target = prefMod - need;
    if (target < 0) target += p; // (prefMod - need + p) % p

    if (lastIndex.has(target)) {
      ans = Math.min(ans, i - lastIndex.get(target));
    }

    // store latest index for this remainder
    lastIndex.set(prefMod, i);
  }

  return ans === n ? -1 : ans;
};
