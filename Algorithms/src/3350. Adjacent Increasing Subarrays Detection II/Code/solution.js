var maxIncreasingSubarrays = function (nums) {
  const n = nums.length;
  if (n < 2) return 0;

  const inc = new Array(n).fill(1);
  for (let i = n - 2; i >= 0; --i)
    inc[i] = nums[i] < nums[i + 1] ? inc[i + 1] + 1 : 1;

  const feasible = (k) => {
    if (k === 0) return true;
    for (let a = 0; a + 2 * k <= n; ++a)
      if (inc[a] >= k && inc[a + k] >= k) return true;
    return false;
  };

  let lo = 0,
    hi = Math.floor(n / 2),
    ans = 0;
  while (lo <= hi) {
    const mid = Math.floor((lo + hi) / 2);
    if (feasible(mid)) (ans = mid), (lo = mid + 1);
    else hi = mid - 1;
  }
  return ans;
};
