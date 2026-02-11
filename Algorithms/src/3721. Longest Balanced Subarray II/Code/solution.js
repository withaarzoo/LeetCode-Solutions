/**
 * JavaScript (Node / leetcode style)
 * @param {number[]} nums
 * @return {number}
 */
var longestBalanced = function (nums) {
  const n = nums.length;
  const pos = new Map();
  for (let i = 0; i < n; ++i) {
    if (!pos.has(nums[i])) pos.set(nums[i], []);
    pos.get(nums[i]).push(i);
  }

  // Segment tree with mn, mx, lazy
  class SegTree {
    constructor(n) {
      this.n = n;
      this.mn = new Array(4 * n).fill(0);
      this.mx = new Array(4 * n).fill(0);
      this.lazy = new Array(4 * n).fill(0);
    }
    apply(idx, v) {
      this.mn[idx] += v;
      this.mx[idx] += v;
      this.lazy[idx] += v;
    }
    push(idx) {
      const z = this.lazy[idx];
      if (z !== 0) {
        this.apply(idx << 1, z);
        this.apply((idx << 1) | 1, z);
        this.lazy[idx] = 0;
      }
    }
    pull(idx) {
      this.mn[idx] = Math.min(this.mn[idx << 1], this.mn[(idx << 1) | 1]);
      this.mx[idx] = Math.max(this.mx[idx << 1], this.mx[(idx << 1) | 1]);
    }
    addRange(idx, l, r, ql, qr, val) {
      if (ql > qr) return;
      if (ql <= l && r <= qr) {
        this.apply(idx, val);
        return;
      }
      this.push(idx);
      const mid = (l + r) >> 1;
      if (ql <= mid)
        this.addRange(idx << 1, l, mid, ql, Math.min(qr, mid), val);
      if (qr > mid)
        this.addRange(
          (idx << 1) | 1,
          mid + 1,
          r,
          Math.max(ql, mid + 1),
          qr,
          val,
        );
      this.pull(idx);
    }
    add(l, r, v) {
      if (l > r) return;
      this.addRange(1, 0, this.n - 1, l, r, v);
    }
    findRightmostZero(idx, l, r, ql, qr) {
      if (ql > qr || qr < l || ql > r) return -1;
      if (this.mn[idx] > 0 || this.mx[idx] < 0) return -1;
      if (l === r) {
        return this.mn[idx] === 0 ? l : -1;
      }
      this.push(idx);
      const mid = (l + r) >> 1;
      if (qr > mid) {
        const res = this.findRightmostZero(
          (idx << 1) | 1,
          mid + 1,
          r,
          Math.max(ql, mid + 1),
          qr,
        );
        if (res !== -1) return res;
      }
      if (ql <= mid) {
        return this.findRightmostZero(idx << 1, l, mid, ql, Math.min(qr, mid));
      }
      return -1;
    }
    findRightmost(l, r) {
      if (l > r) return -1;
      return this.findRightmostZero(1, 0, this.n - 1, l, r);
    }
  }

  const st = new SegTree(n);
  for (let [val, arr] of pos) {
    const sign = val & 1 ? 1 : -1;
    st.add(arr[0], n - 1, sign);
  }
  const ptr = new Map();
  for (let k of pos.keys()) ptr.set(k, 0);

  let ans = 0;
  for (let l = 0; l < n; ++l) {
    const r = st.findRightmost(l, n - 1);
    if (r !== -1) ans = Math.max(ans, r - l + 1);

    const x = nums[l];
    const pi = ptr.get(x);
    ptr.set(x, pi + 1);
    const arr = pos.get(x);
    const nextPos = pi + 1 < arr.length ? arr[pi + 1] : n;
    const sign = x & 1 ? 1 : -1;
    const L = l,
      R = nextPos - 1;
    if (L <= R) st.add(L, R, -sign);
  }
  return ans;
};
