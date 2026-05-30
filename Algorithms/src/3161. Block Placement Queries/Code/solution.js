/**
 * Fenwick tree because I need fast predecessor and successor lookup.
 */
class Fenwick {
  constructor(n) {
    this.n = n;
    this.bit = new Array(n + 1).fill(0);
  }

  // Point add because each obstacle insertion changes one count.
  add(idx, val) {
    for (; idx <= this.n; idx += idx & -idx) {
      this.bit[idx] += val;
    }
  }

  // Prefix sum because I need how many obstacles are on the left side.
  sum(idx) {
    let res = 0;
    for (; idx > 0; idx -= idx & -idx) {
      res += this.bit[idx];
    }
    return res;
  }

  // Find the smallest index with prefix sum at least k.
  kth(k) {
    let idx = 0;
    let step = 1;
    while (step << 1 <= this.n) step <<= 1;

    for (let d = step; d > 0; d >>= 1) {
      const next = idx + d;
      if (next <= this.n && this.bit[next] < k) {
        idx = next;
        k -= this.bit[next];
      }
    }
    return idx + 1;
  }
}

/**
 * Segment tree because I need a fast prefix maximum over gap lengths.
 */
class SegTree {
  constructor(n) {
    this.n = n;
    this.tree = new Array(Math.max(4, 4 * n)).fill(0);
  }

  // Update one coordinate because only one stored gap changes at a time.
  update(node, l, r, pos, val) {
    if (l === r) {
      this.tree[node] = val;
      return;
    }
    const mid = (l + r) >> 1;
    if (pos <= mid) this.update(node << 1, l, mid, pos, val);
    else this.update((node << 1) | 1, mid + 1, r, pos, val);
    this.tree[node] = Math.max(
      this.tree[node << 1],
      this.tree[(node << 1) | 1],
    );
  }

  // Query a prefix because every useful gap ends at or before x.
  query(node, l, r, ql, qr) {
    if (ql > r || qr < l) return 0;
    if (ql <= l && r <= qr) return this.tree[node];
    const mid = (l + r) >> 1;
    return Math.max(
      this.query(node << 1, l, mid, ql, qr),
      this.query((node << 1) | 1, mid + 1, r, ql, qr),
    );
  }
}

/**
 * @param {number[][]} queries
 * @return {boolean[]}
 */
var getResults = function (queries) {
  let mx = 0;
  for (const q of queries) {
    mx = Math.max(mx, q[1]);
  }

  const fw = new Fenwick(mx + 2);
  const st = new SegTree(mx + 1);

  // Sentinel obstacle at 0 keeps predecessor logic simple.
  fw.add(1, 1);

  const ans = [];

  for (const q of queries) {
    const type = q[0];
    const x = q[1];

    if (type === 1) {
      // Count occupied positions strictly smaller than x.
      const leftCount = fw.sum(x);
      const leftPos = fw.kth(leftCount) - 1;

      // Find the next occupied position after x if it exists.
      const occupiedUpToX = fw.sum(x + 1);
      const totalOccupied = fw.sum(mx + 2);
      let rightPos = -1;
      if (occupiedUpToX < totalOccupied) {
        rightPos = fw.kth(occupiedUpToX + 1) - 1;
      }

      // x becomes the end of a new gap.
      st.update(1, 0, mx, x, x - leftPos);

      // The next obstacle now sees a shorter gap.
      if (rightPos !== -1) {
        st.update(1, 0, mx, rightPos, rightPos - x);
      }

      // Mark x as occupied.
      fw.add(x + 1, 1);
    } else {
      const sz = q[2];

      // The last obstacle strictly before x gives the free suffix up to x.
      const leftCount = fw.sum(x);
      const leftPos = fw.kth(leftCount) - 1;

      // Prefix maximum of gaps ending at or before x.
      const bestPrefix = st.query(1, 0, mx, 0, x);

      ans.push(x - leftPos >= sz || bestPrefix >= sz);
    }
  }

  return ans;
};
