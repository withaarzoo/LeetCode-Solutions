/**
 * Heap-based solution with lazy deletion and a chosen set.
 * We avoid naming collisions by scoping helpers inside the function and using unique class names.
 */
/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number} x
 * @return {number[]}
 */
var findXSum = function (nums, k, x) {
  // ---- tiny heap implementation with custom comparator ----
  class _PQ {
    constructor(cmp) {
      this.a = [];
      this.cmp = cmp;
    }
    size() {
      return this.a.length;
    }
    peek() {
      return this.a[0];
    }
    push(v) {
      this.a.push(v);
      this._up(this.size() - 1);
    }
    pop() {
      const n = this.size();
      if (!n) return undefined;
      [this.a[0], this.a[n - 1]] = [this.a[n - 1], this.a[0]];
      const v = this.a.pop();
      this._down(0);
      return v;
    }
    _up(i) {
      while (i) {
        const p = (i - 1) >> 1;
        if (this.cmp(this.a[p], this.a[i]) <= 0) break;
        [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
        i = p;
      }
    }
    _down(i) {
      const n = this.size();
      for (;;) {
        let l = i * 2 + 1,
          r = l + 1,
          b = i;
        if (l < n && this.cmp(this.a[b], this.a[l]) > 0) b = l;
        if (r < n && this.cmp(this.a[b], this.a[r]) > 0) b = r;
        if (b === i) break;
        [this.a[b], this.a[i]] = [this.a[i], this.a[b]];
        i = b;
      }
    }
  }
  const minCmp = (a, b) => (a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1]); // (f asc, v asc)
  const maxCmp = (a, b) => (a[0] !== b[0] ? b[0] - a[0] : b[1] - a[1]); // (f desc, v desc)

  const n = nums.length;
  const ans = new Array(n - k + 1);

  const freq = new Map();
  const chosen = new Set(); // values in TOP
  const hot = new _PQ(minCmp); // TOP worst at top
  const pool = new _PQ(maxCmp); // REST best at top
  let sum = 0n;

  const clean = () => {
    while (hot.size()) {
      const [f, v] = hot.peek();
      if (chosen.has(v) && (freq.get(v) || 0) === f) break;
      hot.pop();
    }
    while (pool.size()) {
      const [f, v] = pool.peek();
      if (!chosen.has(v) && (freq.get(v) || 0) === f && f > 0) break;
      pool.pop();
    }
  };
  const demoteIfChosen = (v) => {
    if (chosen.has(v)) {
      chosen.delete(v);
      const f = freq.get(v) || 0;
      sum -= BigInt(v) * BigInt(f);
    }
  };
  const promoteWhileNeeded = () => {
    clean();
    while (chosen.size < x && pool.size()) {
      const [f, v] = pool.pop();
      if ((freq.get(v) || 0) !== f || chosen.has(v) || f === 0) continue;
      chosen.add(v);
      sum += BigInt(v) * BigInt(f);
      hot.push([f, v]);
      clean();
    }
  };

  const addOne = (v) => {
    demoteIfChosen(v);
    const f = (freq.get(v) || 0) + 1;
    freq.set(v, f);
    pool.push([f, v]);
    if (chosen.size < x) {
      promoteWhileNeeded();
    } else {
      clean();
      if (pool.size() && hot.size()) {
        const [bf, bv] = pool.peek();
        const [wf, wv] = hot.peek();
        if (bf > wf || (bf === wf && bv > wv)) {
          // promote best, demote worst
          pool.pop();
          chosen.add(bv);
          sum += BigInt(bv) * BigInt(bf);
          hot.push([bf, bv]);
          clean();
          const [df, dv] = hot.pop();
          if (chosen.has(dv) && (freq.get(dv) || 0) === df) {
            chosen.delete(dv);
            sum -= BigInt(dv) * BigInt(df);
            pool.push([df, dv]);
          }
          clean();
        }
      }
    }
  };
  const removeOne = (v) => {
    demoteIfChosen(v);
    const f = (freq.get(v) || 0) - 1;
    if (f <= 0) freq.delete(v);
    else {
      freq.set(v, f);
      pool.push([f, v]);
    }
    promoteWhileNeeded();
  };

  for (let i = 0; i < k; ++i) addOne(nums[i]);
  ans[0] = Number(sum);
  for (let i = k; i < n; ++i) {
    removeOne(nums[i - k]);
    addOne(nums[i]);
    ans[i - k + 1] = Number(sum);
  }
  return ans;
};
