function maxSum(nums: number[], k: number): number {
  const MOD = 1000000007n;
  const n = nums.length;
  const cnt = new Array(30).fill(0);
  for (const x of nums) {
    for (let b = 0; b < 30; b++) {
      if (x & (1 << b)) cnt[b]++;
    }
  }
  const a = new Array(n).fill(0n);
  for (let b = 29; b >= 0; b--) {
    const c = cnt[b];
    for (let i = 0; i < c; i++) {
      a[i] |= 1n << BigInt(b);
    }
  }
  let ans = 0n;
  for (let i = 0; i < k; i++) {
    ans = (ans + a[i] * a[i]) % MOD;
  }
  return Number(ans);
}
