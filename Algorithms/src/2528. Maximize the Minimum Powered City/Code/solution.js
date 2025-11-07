/**
 * @param {number[]} stations
 * @param {number} r
 * @param {number} k
 * @return {number}
 */
var maxPower = function (stations, r, k) {
  const n = stations.length;

  // 1) Base power using a difference array.
  const diff = Array(n + 1).fill(0);
  for (let i = 0; i < n; i++) {
    const L = Math.max(0, i - r);
    const R = Math.min(n, i + r + 1);
    diff[L] += stations[i];
    diff[R] -= stations[i];
  }
  const base = Array(n).fill(0);
  let run = 0;
  for (let i = 0; i < n; i++) {
    run += diff[i];
    base[i] = run;
  }

  // 2) Binary search the best T.
  let lo = 0;
  let hi = stations.reduce((s, v) => s + v, 0) + k;
  let ans = 0;

  const feasible = (T) => {
    const add = Array(n + 1).fill(0);
    let extra = 0;
    let used = 0;
    for (let i = 0; i < n; i++) {
      extra += add[i];
      let curr = base[i] + extra;
      if (curr < T) {
        const need = T - curr;
        used += need;
        if (used > k) return false;
        extra += need;
        const end = Math.min(n, i + 2 * r + 1);
        add[end] -= need; // coverage from these new stations ends after 'end-1'
      }
    }
    return true;
  };

  while (lo <= hi) {
    const mid = Math.floor((lo + hi) / 2);
    if (feasible(mid)) {
      ans = mid;
      lo = mid + 1;
    } else {
      hi = mid - 1;
    }
  }
  return ans;
};
