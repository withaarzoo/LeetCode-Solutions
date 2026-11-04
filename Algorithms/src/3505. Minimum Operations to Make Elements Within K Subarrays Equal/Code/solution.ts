function minOperations(nums: number[], x: number, k: number): number {
  const n: number = nums.length;
  const sorted: Int32Array = Int32Array.from(nums).sort();
  let m: number = 0;
  for (let i = 0; i < n; i++) {
    if (m === 0 || sorted[i] !== sorted[m - 1]) {
      sorted[m++] = sorted[i];
    }
  }
  const vals: Int32Array = sorted;
  const rk: Int32Array = new Int32Array(n);
  for (let i = 0; i < n; i++) {
    let lo = 0,
      hi = m - 1;
    while (lo < hi) {
      const mid = (lo + hi) >> 1;
      if (vals[mid] < nums[i]) lo = mid + 1;
      else hi = mid;
    }
    rk[i] = lo + 1;
  }
  const cnt: Int32Array = new Int32Array(m + 1);
  const sm: Float64Array = new Float64Array(m + 1);
  let top: number = 1;
  while (top * 2 <= m) top *= 2;
  const need: number = (x + 1) >> 1;
  const w: number = n - x + 1;
  const cost: Float64Array = new Float64Array(w);
  let total: number = 0;
  for (let i = 0; i < n; i++) {
    total += nums[i];
    for (let p = rk[i]; p <= m; p += p & -p) {
      cnt[p] += 1;
      sm[p] += nums[i];
    }
    if (i >= x) {
      total -= nums[i - x];
      for (let p = rk[i - x]; p <= m; p += p & -p) {
        cnt[p] -= 1;
        sm[p] -= nums[i - x];
      }
    }
    if (i >= x - 1) {
      let pos = 0,
        rem = need,
        s0 = 0;
      for (let pw = top; pw > 0; pw >>= 1) {
        const nxt = pos + pw;
        if (nxt <= m && cnt[nxt] < rem) {
          pos = nxt;
          rem -= cnt[nxt];
          s0 += sm[nxt];
        }
      }
      const med: number = vals[pos];
      const c0: number = need - rem;
      cost[i - x + 1] = med * c0 - s0 + (total - s0) - med * (x - c0);
    }
  }
  let prev: Float64Array = new Float64Array(n + 1);
  let cur: Float64Array = new Float64Array(n + 1);
  for (let j = 0; j < k; j++) {
    for (let i = 0; i < x; i++) cur[i] = Infinity;
    for (let i = x; i <= n; i++) {
      const a = cur[i - 1];
      const b = prev[i - x] + cost[i - x];
      cur[i] = a < b ? a : b;
    }
    const t = prev;
    prev = cur;
    cur = t;
  }
  return prev[n];
}
